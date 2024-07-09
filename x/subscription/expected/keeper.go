package expected

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v2"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	AddDeposit(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
}

type ProviderKeeper interface {
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
}

type NodeKeeper interface {
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
	GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool)
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool)
}

type SessionKeeper interface {
	GetCount(ctx sdk.Context) uint64
	GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool)
	SetCount(ctx sdk.Context, count uint64)
	SetSession(ctx sdk.Context, session sessiontypes.Session)
	SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	SetSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64)
	SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64)
	SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64)
	SetSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64)
	StatusChangeDelay(ctx sdk.Context) time.Duration
	SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error
}
