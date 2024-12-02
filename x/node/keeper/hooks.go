package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

// SessionInactivePreHook handles the necessary operations for when a session becomes inactive.
func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	// Retrieve the session by ID and check if it is found, return an error if not.
	item, found := k.GetSession(ctx, id)
	if !found {
		return types.NewErrorSessionNotFound(id)
	}

	// Validate that the session status is "InactivePending", otherwise return an error.
	if !item.GetStatus().Equal(v1base.StatusInactivePending) {
		return types.NewErrorInvalidSessionStatus(item.GetID(), item.GetStatus())
	}

	// Assert the item to a v3.Session type and proceed only if the assertion is successful.
	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	// Convert account address and node address from Bech32 format.
	accAddr, err := sdk.AccAddressFromBech32(session.AccAddress)
	if err != nil {
		return err
	}

	nodeAddr, err := base.NodeAddressFromBech32(session.NodeAddress)
	if err != nil {
		return err
	}

	// Retrieve the staking share and the total payment amount for the session.
	share := k.StakingShare(ctx)
	totalPayment := session.PaymentAmount()

	// Calculate the staking reward based on the total payment and staking share, then send it to the fee collector module.
	reward := baseutils.GetProportionOfCoin(totalPayment, share)
	if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return err
	}

	// Send remaining payment from the user's deposit to the node's account.
	payment := totalPayment.Sub(reward)
	if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment); err != nil {
		return err
	}

	// Emit an event indicating the payment and staking reward.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventPay{
			ID:            session.ID,
			AccAddress:    session.AccAddress,
			NodeAddress:   session.NodeAddress,
			Payment:       payment.String(),
			StakingReward: reward.String(),
		},
	)

	// Process refund by subtracting the refund amount from the user's deposit.
	refund := session.RefundAmount()
	if err := k.SubtractDeposit(ctx, accAddr, refund); err != nil {
		return err
	}

	// Emit an event indicating the refund has been processed.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventRefund{
			ID:         session.ID,
			AccAddress: session.AccAddress,
			Amount:     refund.String(),
		},
	)

	// Delete session records from the store, including records for account and node.
	k.DeleteSession(ctx, item.GetID())
	k.DeleteSessionForAccount(ctx, accAddr, item.GetID())
	k.DeleteSessionForNode(ctx, nodeAddr, item.GetID())

	return nil
}
