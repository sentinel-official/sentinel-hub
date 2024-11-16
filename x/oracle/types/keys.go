package types

import (
	ibcicqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v7/types"
)

const (
	ModuleName = "oracle"
	StoreKey   = ModuleName
	Version    = ibcicqtypes.Version
)

var (
	ParamsKey = []byte{0x00}

	AssetKeyPrefix = []byte{0x10}
)

func AssetKey(denom string) []byte {
	return append(AssetKeyPrefix, []byte(denom)...)
}
