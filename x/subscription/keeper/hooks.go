package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

// SessionInactiveHook is a function that handles the end of a session.
// It updates the allocation's utilized bytes, calculates and sends payments, and staking rewards.
func (k *Keeper) SessionInactiveHook(ctx sdk.Context, id uint64, accAddr sdk.AccAddress, nodeAddr base.NodeAddress, utilisedBytes sdkmath.Int) error {
	// Retrieve the session associated with the provided session ID.
	session, found := k.GetSession(ctx, id)
	if !found {
		return fmt.Errorf("session %d does not exist", id)
	}

	// Check if the session has the correct status for processing.
	if !session.Status.Equal(base.StatusInactivePending) {
		return fmt.Errorf("invalid status %s for session %d", session.Status, session.ID)
	}

	// Retrieve the subscription associated with the session.
	subscription, found := k.GetSubscription(ctx, session.SubscriptionID)
	if !found {
		return fmt.Errorf("subscription %d does not exist", session.SubscriptionID)
	}

	// If the subscription is a NodeSubscription with non-zero duration (hours), no further action is needed.
	if s, ok := subscription.(*v2.NodeSubscription); ok && s.Hours != 0 {
		return nil
	}

	// Retrieve the allocation associated with the subscription and account address.
	alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", subscription.GetID(), accAddr)
	}

	var (
		gigabytePrice  sdk.Coin            // Gigabyte price based on the deposit amount and gigabytes (for NodeSubscription).
		previousAmount = sdkmath.ZeroInt() // Amount paid for previous utilization (for NodeSubscription).
	)

	// Calculate payment amounts based on the subscription type (NodeSubscription).
	if s, ok := subscription.(*v2.NodeSubscription); ok && s.Gigabytes != 0 {
		gigabytePrice = sdk.NewCoin(
			s.Deposit.Denom,
			s.Deposit.Amount.QuoRaw(s.Gigabytes),
		)
		previousAmount = baseutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
	}

	// Update the allocation's utilized bytes by adding the provided bytes.
	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(utilisedBytes)
	// Ensure that the utilized bytes don't exceed the granted bytes.
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	// Save the updated allocation to the store.
	k.SetAllocation(ctx, alloc)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventAllocate{
			Address:       alloc.Address,
			GrantedBytes:  alloc.GrantedBytes,
			UtilisedBytes: alloc.UtilisedBytes,
			ID:            alloc.ID,
		},
	)

	// Calculate the current payment amount based on the subscription type (NodeSubscription).
	if s, ok := subscription.(*v2.NodeSubscription); ok && s.Gigabytes != 0 {
		// Calculate the payment to be made for the current utilization.
		var (
			currentAmount = baseutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
			payment       = sdk.NewCoin(gigabytePrice.Denom, currentAmount.Sub(previousAmount))
			stakingShare  = k.node.StakingShare(ctx)
			stakingReward = baseutils.GetProportionOfCoin(payment, stakingShare)
		)

		// Move the staking reward from the deposit to the fee collector module account.
		if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
			return err
		}

		// Subtract the staking reward from the payment to get the final payment amount.
		payment = payment.Sub(stakingReward)

		// Send the payment amount from the deposit to the node address.
		if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment); err != nil {
			return err
		}

		// Emit an event for the session payment.
		ctx.EventManager().EmitTypedEvent(
			&v2.EventPayForSession{
				Address:        session.Address,
				NodeAddress:    session.NodeAddress,
				Payment:        payment.String(),
				StakingReward:  stakingReward.String(),
				SessionID:      session.ID,
				SubscriptionID: session.SubscriptionID,
			},
		)
	}

	return nil
}
