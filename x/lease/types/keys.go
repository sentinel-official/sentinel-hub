package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/sentinel-official/hub/v12/types"
)

const (
	ModuleName = "lease"
)

var (
	CountKey                    = []byte{0x00}
	ParamsKey                   = []byte{0x01}
	LeaseKeyPrefix              = []byte{0x10}
	LeaseForNodeKeyPrefix       = []byte{0x11}
	LeaseForProviderKeyPrefix   = []byte{0x12}
	LeaseForInactiveAtKeyPrefix = []byte{0x13}
	LeaseForPayoutAtKeyPrefix   = []byte{0x14}
	LeaseForRenewalAtKeyPrefix  = []byte{0x15}
)

func LeaseKey(id uint64) []byte {
	return append(LeaseKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForNodeKeyPrefix(addr base.NodeAddress) []byte {
	return append(LeaseForNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetLeaseForNodeByProviderKeyPrefix(nodeAddr base.NodeAddress, provAddr base.ProvAddress) []byte {
	return append(GetLeaseForNodeKeyPrefix(nodeAddr), sdkaddress.MustLengthPrefix(provAddr.Bytes())...)
}

func LeaseForNodeByProviderKey(nodeAddr base.NodeAddress, provAddr base.ProvAddress, id uint64) []byte {
	return append(GetLeaseForNodeByProviderKeyPrefix(nodeAddr, provAddr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForProviderKeyPrefix(addr base.ProvAddress) []byte {
	return append(LeaseForProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetLeaseForProviderByNodeKeyPrefix(provAddr base.ProvAddress, nodeAddr base.NodeAddress) (key []byte) {
	return append(GetLeaseForProviderKeyPrefix(provAddr), sdkaddress.MustLengthPrefix(nodeAddr.Bytes())...)
}

func LeaseForProviderByNodeKey(provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) []byte {
	return append(GetLeaseForProviderByNodeKeyPrefix(provAddr, nodeAddr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForInactiveAtKeyPrefix(at time.Time) []byte {
	return append(LeaseForInactiveAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func LeaseForInactiveAtKey(at time.Time, id uint64) []byte {
	return append(GetLeaseForInactiveAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForPayoutAtKeyPrefix(at time.Time) []byte {
	return append(LeaseForPayoutAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func LeaseForPayoutAtKey(at time.Time, id uint64) []byte {
	return append(GetLeaseForPayoutAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForRenewalAtKeyPrefix(at time.Time) []byte {
	return append(LeaseForRenewalAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func LeaseForRenewalAtKey(at time.Time, id uint64) []byte {
	return append(GetLeaseForRenewalAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func IDFromLeaseForNodeByProviderKey(key []byte) uint64 {
	// prefix (1 byte) | nodeAddrLen (1 byte) | nodeAddr (nodeAddrLen bytes) | provAddrLen (1 byte) | provAddr (provAddrLen bytes) | id (8 bytes)

	nodeAddrLen, provAddrLen := int(key[1]), int(key[2+int(key[1])])
	if len(key) != 11+nodeAddrLen+provAddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+nodeAddrLen+provAddrLen))
	}

	return sdk.BigEndianToUint64(key[3+nodeAddrLen+provAddrLen:])
}

func IDFromLeaseForProviderByNodeKey(key []byte) uint64 {
	// prefix (1 byte) | provAddrLen(1 byte) | provAddr (provAddrLen bytes) | nodeAddrLen(1 byte) | nodeAddr (nodeAddrLen bytes) | id (8 bytes)

	provAddrLen, nodeAddrLen := int(key[1]), int(key[2+int(key[1])])
	if len(key) != 11+provAddrLen+nodeAddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+provAddrLen+nodeAddrLen))
	}

	return sdk.BigEndianToUint64(key[3+provAddrLen+nodeAddrLen:])
}

func IDFromLeaseForInactiveAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}

func IDFromLeaseForPayoutAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}

func IDFromLeaseForRenewalAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
