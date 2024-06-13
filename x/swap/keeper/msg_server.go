package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/swap/types"
	"github.com/sentinel-official/hub/v12/x/swap/types/v1"
)

var (
	_ v1.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) v1.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) MsgSwap(c context.Context, msg *v1.MsgSwapRequest) (*v1.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if !k.SwapEnabled(ctx) {
		return nil, types.ErrorSwapIsDisabled
	}
	if k.ApproveBy(ctx) != msg.From {
		return nil, types.ErrorUnauthorized
	}
	if k.HasSwap(ctx, types.BytesToHash(msg.TxHash)) {
		return nil, types.ErrorDuplicateSwap
	}

	var (
		coin = sdk.NewCoin(k.SwapDenom(ctx), msg.Amount.Quo(types.PrecisionLoss))
		swap = v1.Swap{
			TxHash:   msg.TxHash,
			Receiver: msg.Receiver,
			Amount:   coin,
		}
	)

	msgReceiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	if err := k.MintCoin(ctx, coin); err != nil {
		return nil, err
	}
	if err := k.SendCoinFromModuleToAccount(ctx, msgReceiver, coin); err != nil {
		return nil, err
	}

	k.SetSwap(ctx, swap)
	ctx.EventManager().EmitTypedEvent(
		&v1.EventSwap{
			TxHash:   swap.TxHash,
			Receiver: swap.Receiver,
		},
	)

	return &v1.MsgSwapResponse{}, nil
}
