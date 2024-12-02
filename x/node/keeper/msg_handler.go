package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) HandleMsgRegisterNode(ctx sdk.Context, msg *v3.MsgRegisterNodeRequest) (*v3.MsgRegisterNodeResponse, error) {
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr := base.NodeAddress(accAddr.Bytes())
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	deposit := k.Deposit(ctx)
	if err := k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	node := v3.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteURL:      msg.RemoteURL,
		Status:         v1base.StatusInactive,
		InactiveAt:     time.Time{},
		StatusAt:       ctx.BlockTime(),
	}

	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreate{
			NodeAddress:    node.Address,
			GigabytePrices: node.GigabytePrices.String(),
			HourlyPrices:   node.HourlyPrices.String(),
			RemoteURL:      node.RemoteURL,
		},
	)

	return &v3.MsgRegisterNodeResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateNodeDetails(ctx sdk.Context, msg *v3.MsgUpdateNodeDetailsRequest) (*v3.MsgUpdateNodeDetailsResponse, error) {
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	node.GigabytePrices = msg.GigabytePrices
	node.HourlyPrices = msg.HourlyPrices
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateDetails{
			NodeAddress:    node.Address,
			GigabytePrices: node.GigabytePrices.String(),
			HourlyPrices:   node.HourlyPrices.String(),
			RemoteURL:      node.RemoteURL,
		},
	)

	return &v3.MsgUpdateNodeDetailsResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateNodeStatus(ctx sdk.Context, msg *v3.MsgUpdateNodeStatusRequest) (*v3.MsgUpdateNodeStatusResponse, error) {
	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if msg.Status.Equal(v1base.StatusInactive) {
		if err := k.NodeInactivePreHook(ctx, nodeAddr); err != nil {
			return nil, err
		}
	}

	if msg.Status.Equal(v1base.StatusActive) {
		if node.Status.Equal(v1base.StatusInactive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}
	if msg.Status.Equal(v1base.StatusInactive) {
		if node.Status.Equal(v1base.StatusActive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}
	}

	k.DeleteNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	node.Status = msg.Status
	node.InactiveAt = time.Time{}
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(v1base.StatusActive) {
		node.InactiveAt = k.GetInactiveAt(ctx)
	}

	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			NodeAddress: node.Address,
			Status:      node.Status,
		},
	)

	return &v3.MsgUpdateNodeStatusResponse{}, nil
}

func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	if msg.Gigabytes != 0 {
		if ok := k.IsValidSessionGigabytes(ctx, msg.Gigabytes); !ok {
			return nil, types.NewErrorInvalidGigabytes(msg.Gigabytes)
		}
	}
	if msg.Hours != 0 {
		if ok := k.IsValidSessionHours(ctx, msg.Hours); !ok {
			return nil, types.NewErrorInvalidHours(msg.Hours)
		}
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	var (
		count      = k.GetSessionCount(ctx)
		inactiveAt = k.GetSessionInactiveAt(ctx)
		basePrice  = sdk.DecCoin{Amount: sdkmath.LegacyZeroDec()}
		session    = &v3.Session{
			ID:            count + 1,
			AccAddress:    accAddr.String(),
			NodeAddress:   nodeAddr.String(),
			Price:         sdk.Coin{Amount: sdkmath.ZeroInt()},
			DownloadBytes: sdkmath.ZeroInt(),
			UploadBytes:   sdkmath.ZeroInt(),
			MaxGigabytes:  msg.Gigabytes,
			Duration:      0,
			MaxHours:      msg.Hours,
			Status:        v1base.StatusActive,
			InactiveAt:    inactiveAt,
			StatusAt:      ctx.BlockTime(),
		}
	)

	if msg.Gigabytes != 0 {
		basePrice, found = node.GigabytePrice(msg.Denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(msg.Denom)
		}
	}
	if msg.Hours != 0 {
		basePrice, found = node.HourlyPrice(msg.Denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(msg.Denom)
		}
	}

	session.Price, err = k.GetQuote(ctx, basePrice)
	if err != nil {
		return nil, err
	}

	deposit := session.DepositAmount()
	if err := k.AddDeposit(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	k.SetSessionCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreateSession{
			ID:           session.ID,
			AccAddress:   session.AccAddress,
			NodeAddress:  session.NodeAddress,
			Price:        session.Price.String(),
			MaxGigabytes: session.MaxGigabytes,
			MaxHours:     session.MaxHours,
		},
	)

	return &v3.MsgStartSessionResponse{
		ID: session.ID,
	}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	minGigabytePrices := k.MinGigabytePrices(ctx)
	minHourlyPrices := k.MinHourlyPrices(ctx)
	k.SetParams(ctx, msg.Params)

	minGigabytePricesModified := !msg.Params.MinGigabytePrices.IsEqual(minGigabytePrices)
	minHourlyPricesModified := !msg.Params.MinHourlyPrices.IsEqual(minHourlyPrices)

	if minGigabytePricesModified {
		minGigabytePrices = k.MinGigabytePrices(ctx)
	}
	if minHourlyPricesModified {
		minHourlyPrices = k.MinHourlyPrices(ctx)
	}

	if minGigabytePricesModified || minHourlyPricesModified {
		k.IterateNodes(ctx, func(_ int, item v3.Node) bool {
			if minGigabytePricesModified {
				for _, coin := range minGigabytePrices {
					amount := item.GigabytePrices.AmountOf(coin.Denom)
					if amount.LT(coin.Amount) {
						item.GigabytePrices = item.GigabytePrices.Sub(
							sdk.NewDecCoins(sdk.NewDecCoinFromDec(coin.Denom, amount)),
						).Add(coin)
					}
				}
			}

			if minHourlyPricesModified {
				for _, coin := range minHourlyPrices {
					amount := item.HourlyPrices.AmountOf(coin.Denom)
					if amount.LT(coin.Amount) {
						item.HourlyPrices = item.HourlyPrices.Sub(
							sdk.NewDecCoins(sdk.NewDecCoinFromDec(coin.Denom, amount)),
						).Add(coin)
					}
				}
			}

			k.SetNode(ctx, item)
			ctx.EventManager().EmitTypedEvent(
				&v3.EventUpdateDetails{
					NodeAddress:    item.Address,
					GigabytePrices: item.GigabytePrices.String(),
					HourlyPrices:   item.HourlyPrices.String(),
					RemoteURL:      item.RemoteURL,
				},
			)

			return false
		})
	}

	return &v3.MsgUpdateParamsResponse{}, nil
}
