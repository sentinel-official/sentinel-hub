// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/session/types
// ALIASGEN: github.com/sentinel-official/hub/x/session/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/session/querier
package session

import (
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/querier"
	"github.com/sentinel-official/hub/x/session/types"
)

const (
	AttributeKeyCount            = types.AttributeKeyCount
	AttributeKeyID               = types.AttributeKeyID
	AttributeKeySubscription     = types.AttributeKeySubscription
	AttributeKeyAddress          = types.AttributeKeyAddress
	ModuleName                   = types.ModuleName
	QuerierRoute                 = types.QuerierRoute
	DefaultInactiveDuration      = types.DefaultInactiveDuration
	QuerySession                 = types.QuerySession
	QuerySessions                = types.QuerySessions
	QuerySessionsForSubscription = types.QuerySessionsForSubscription
	QuerySessionsForNode         = types.QuerySessionsForNode
	QuerySessionsForAddress      = types.QuerySessionsForAddress
	QueryOngoingSession          = types.QueryOngoingSession
)

var (
	// functions aliases
	RegisterCodec                         = types.RegisterCodec
	NewGenesisState                       = types.NewGenesisState
	DefaultGenesisState                   = types.DefaultGenesisState
	SessionKey                            = types.SessionKey
	GetSessionForSubscriptionKeyPrefix    = types.GetSessionForSubscriptionKeyPrefix
	SessionForSubscriptionKey             = types.SessionForSubscriptionKey
	GetSessionForNodeKeyPrefix            = types.GetSessionForNodeKeyPrefix
	SessionForNodeKey                     = types.SessionForNodeKey
	GetSessionForAddressKeyPrefix         = types.GetSessionForAddressKeyPrefix
	SessionForAddressKey                  = types.SessionForAddressKey
	GetOngoingSessionPrefix               = types.GetOngoingSessionPrefix
	OngoingSessionKey                     = types.OngoingSessionKey
	GetActiveSessionAtKeyPrefix           = types.GetActiveSessionAtKeyPrefix
	ActiveSessionAtKey                    = types.ActiveSessionAtKey
	IDFromSessionForSubscriptionKey       = types.IDFromSessionForSubscriptionKey
	IDFromSessionForNodeKey               = types.IDFromSessionForNodeKey
	IDFromSessionForAddressKey            = types.IDFromSessionForAddressKey
	IDFromActiveSessionAtKey              = types.IDFromActiveSessionAtKey
	NewMsgUpsert                          = types.NewMsgUpsert
	NewParams                             = types.NewParams
	DefaultParams                         = types.DefaultParams
	ParamsKeyTable                        = types.ParamsKeyTable
	NewQuerySessionParams                 = types.NewQuerySessionParams
	NewQuerySessionsParams                = types.NewQuerySessionsParams
	NewQuerySessionsForSubscriptionParams = types.NewQuerySessionsForSubscriptionParams
	NewQuerySessionsForNodeParams         = types.NewQuerySessionsForNodeParams
	NewQuerySessionsForAddressParams      = types.NewQuerySessionsForAddressParams
	NewQueryOngoingSessionParams          = types.NewQueryOngoingSessionParams
	NewKeeper                             = keeper.NewKeeper
	Querier                               = querier.Querier

	// variable aliases
	ModuleCdc                       = types.ModuleCdc
	ErrorMarshal                    = types.ErrorMarshal
	ErrorUnmarshal                  = types.ErrorUnmarshal
	ErrorUnknownMsgType             = types.ErrorUnknownMsgType
	ErrorUnknownQueryType           = types.ErrorUnknownQueryType
	ErrorInvalidField               = types.ErrorInvalidField
	ErrorSubscriptionDoesNotExit    = types.ErrorSubscriptionDoesNotExit
	ErrorInvalidSubscriptionStatus  = types.ErrorInvalidSubscriptionStatus
	ErrorUnauthorized               = types.ErrorUnauthorized
	ErrorQuotaDoesNotExist          = types.ErrorQuotaDoesNotExist
	ErrorInvalidBandwidth           = types.ErrorInvalidBandwidth
	EventTypeSetCount               = types.EventTypeSetCount
	EventTypeSetActive              = types.EventTypeSetActive
	EventTypeUpdate                 = types.EventTypeUpdate
	ParamsSubspace                  = types.ParamsSubspace
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	EventModuleName                 = types.EventModuleName
	CountKey                        = types.CountKey
	SessionKeyPrefix                = types.SessionKeyPrefix
	SessionForSubscriptionKeyPrefix = types.SessionForSubscriptionKeyPrefix
	SessionForNodeKeyPrefix         = types.SessionForNodeKeyPrefix
	SessionForAddressKeyPrefix      = types.SessionForAddressKeyPrefix
	OngoingSessionKeyPrefix         = types.OngoingSessionKeyPrefix
	ActiveSessionAtKeyPrefix        = types.ActiveSessionAtKeyPrefix
	KeyInactiveDuration             = types.KeyInactiveDuration
)

type (
	GenesisState                       = types.GenesisState
	MsgUpsert                          = types.MsgUpsert
	Params                             = types.Params
	QuerySessionParams                 = types.QuerySessionParams
	QuerySessionsParams                = types.QuerySessionsParams
	QuerySessionsForSubscriptionParams = types.QuerySessionsForSubscriptionParams
	QuerySessionsForNodeParams         = types.QuerySessionsForNodeParams
	QuerySessionsForAddressParams      = types.QuerySessionsForAddressParams
	QueryOngoingSessionParams          = types.QueryOngoingSessionParams
	Session                            = types.Session
	Sessions                           = types.Sessions
	Keeper                             = keeper.Keeper
)
