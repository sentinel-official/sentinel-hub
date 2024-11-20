package expected

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type DepositKeeper interface {
	AddDeposit(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
	SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
}

type LeaseKeeper interface {
	NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error
}

type SessionKeeper interface {
	DeleteSession(ctx sdk.Context, id uint64)
	DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	DeleteSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64)
	GetCount(ctx sdk.Context) uint64
	GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool)
	NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error
	SetCount(ctx sdk.Context, count uint64)
	SetSession(ctx sdk.Context, session sessiontypes.Session)
	SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64)
	SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64)
	StatusChangeDelay(ctx sdk.Context) time.Duration
}
