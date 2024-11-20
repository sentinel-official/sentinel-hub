package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) HandleInactiveSession(ctx sdk.Context, id uint64) error {
	item, found := k.session.GetSession(ctx, id)
	if !found {
		return fmt.Errorf("session %d does not exist", id)
	}
	if !item.GetStatus().Equal(v1base.StatusInactivePending) {
		return fmt.Errorf("invalid status %s for session %d", item.GetStatus(), item.GetID())
	}

	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	accAddr, err := sdk.AccAddressFromBech32(session.AccAddress)
	if err != nil {
		return err
	}

	nodeAddr, err := base.NodeAddressFromBech32(session.NodeAddress)
	if err != nil {
		return err
	}

	amount := session.PaymentAmount()
	share := k.StakingShare(ctx)

	reward := baseutils.GetProportionOfCoin(amount, share)
	if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return err
	}

	payment := amount.Sub(reward)
	if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment); err != nil {
		return err
	}

	ctx.EventManager().EmitTypedEvent(
		&v3.EventPay{
			ID:            session.ID,
			AccAddress:    session.AccAddress,
			NodeAddress:   session.NodeAddress,
			Payment:       payment.String(),
			StakingReward: reward.String(),
		},
	)

	refund := session.RefundAmount()
	if err := k.SubtractDeposit(ctx, accAddr, refund); err != nil {
		return err
	}

	ctx.EventManager().EmitTypedEvent(
		&v3.EventRefund{
			ID:         session.ID,
			AccAddress: session.AccAddress,
			Amount:     refund.String(),
		},
	)

	k.session.DeleteSession(ctx, item.GetID())
	k.session.DeleteSessionForAccount(ctx, accAddr, item.GetID())
	k.session.DeleteSessionForNode(ctx, nodeAddr, item.GetID())

	return nil
}
