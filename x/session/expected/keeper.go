package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
}

type NodeKeeper interface {
	HandleInactiveSession(ctx sdk.Context, id uint64) error
}

type SubscriptionKeeper interface {
	HandleInactiveSession(ctx sdk.Context, id uint64) error
}
