// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/params.proto

package v2

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	Deposit                  types.Coin                               `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	ActiveDuration           time.Duration                            `protobuf:"bytes,2,opt,name=active_duration,json=activeDuration,proto3,stdduration" json:"active_duration"`
	MaxGigabytePrices        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=max_gigabyte_prices,json=maxGigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_gigabyte_prices"`
	MinGigabytePrices        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=min_gigabyte_prices,json=minGigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_gigabyte_prices"`
	MaxHourlyPrices          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=max_hourly_prices,json=maxHourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_hourly_prices"`
	MinHourlyPrices          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=min_hourly_prices,json=minHourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_hourly_prices"`
	MaxSubscriptionGigabytes int64                                    `protobuf:"varint,7,opt,name=max_subscription_gigabytes,json=maxSubscriptionGigabytes,proto3" json:"max_subscription_gigabytes,omitempty"`
	MinSubscriptionGigabytes int64                                    `protobuf:"varint,8,opt,name=min_subscription_gigabytes,json=minSubscriptionGigabytes,proto3" json:"min_subscription_gigabytes,omitempty"`
	MaxSubscriptionHours     int64                                    `protobuf:"varint,9,opt,name=max_subscription_hours,json=maxSubscriptionHours,proto3" json:"max_subscription_hours,omitempty"`
	MinSubscriptionHours     int64                                    `protobuf:"varint,10,opt,name=min_subscription_hours,json=minSubscriptionHours,proto3" json:"min_subscription_hours,omitempty"`
	StakingShare             cosmossdk_io_math.LegacyDec              `protobuf:"bytes,11,opt,name=staking_share,json=stakingShare,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"staking_share"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f1279255e9f358, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "sentinel.node.v2.Params")
}

func init() { proto.RegisterFile("sentinel/node/v2/params.proto", fileDescriptor_02f1279255e9f358) }

var fileDescriptor_02f1279255e9f358 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0xed, 0x5f, 0xfa, 0x4b, 0x5a, 0x17, 0x28, 0x84, 0x0a, 0x99, 0x20, 0x9c, 0x08, 0x96,
	0x2c, 0xbd, 0x23, 0xa1, 0x0b, 0x12, 0x53, 0xa8, 0x44, 0x87, 0x0e, 0x55, 0x2a, 0x31, 0xb0, 0x44,
	0x67, 0xe7, 0xe2, 0x3c, 0x4a, 0x7c, 0x67, 0xf9, 0xce, 0x26, 0x11, 0x6f, 0x82, 0x91, 0x97, 0x80,
	0x58, 0x78, 0x1b, 0x19, 0x3b, 0x22, 0x86, 0x16, 0x92, 0x37, 0x82, 0xee, 0x8f, 0x51, 0x9a, 0x02,
	0x13, 0x9d, 0x72, 0xf1, 0xdd, 0xe7, 0xf9, 0x3c, 0xfe, 0x3e, 0xf2, 0x79, 0x8f, 0x05, 0x65, 0x12,
	0x18, 0x9d, 0x62, 0xc6, 0x87, 0x14, 0x17, 0x5d, 0x9c, 0x92, 0x8c, 0x24, 0x02, 0xa5, 0x19, 0x97,
	0xbc, 0x7e, 0xb7, 0xdc, 0x46, 0x6a, 0x1b, 0x15, 0xdd, 0x46, 0x10, 0x71, 0x91, 0x70, 0x81, 0x43,
	0x22, 0x28, 0x2e, 0x3a, 0x21, 0x95, 0xa4, 0x83, 0x23, 0x0e, 0xcc, 0x10, 0x8d, 0xfd, 0x98, 0xc7,
	0x5c, 0x2f, 0xb1, 0x5a, 0xd9, 0xa7, 0x41, 0xcc, 0x79, 0x3c, 0xa5, 0x58, 0xff, 0x0b, 0xf3, 0x11,
	0x1e, 0xe6, 0x19, 0x91, 0xc0, 0x2d, 0xf5, 0xe4, 0x4b, 0xcd, 0xab, 0x9e, 0x6a, 0x71, 0xfd, 0x85,
	0x57, 0x1b, 0xd2, 0x94, 0x0b, 0x90, 0xbe, 0xdb, 0x72, 0xdb, 0xbb, 0xdd, 0x87, 0xc8, 0x28, 0x91,
	0x52, 0x22, 0xab, 0x44, 0xaf, 0x38, 0xb0, 0xde, 0xd6, 0xe2, 0xa2, 0xe9, 0xf4, 0xcb, 0xf3, 0xf5,
	0x13, 0x6f, 0x8f, 0x44, 0x12, 0x0a, 0x3a, 0x28, 0xcb, 0xfb, 0xff, 0xd9, 0x12, 0xc6, 0x8f, 0x4a,
	0x3f, 0x3a, 0xb2, 0x07, 0x7a, 0xdb, 0xaa, 0xc4, 0xc7, 0xcb, 0xa6, 0xdb, 0xbf, 0x63, 0xd8, 0x72,
	0xa7, 0xfe, 0xde, 0xbb, 0x9f, 0x90, 0xd9, 0x20, 0x86, 0x98, 0x84, 0x73, 0x49, 0x07, 0x69, 0x06,
	0x11, 0x15, 0x7e, 0xa5, 0x55, 0xf9, 0x7b, 0x53, 0xcf, 0x54, 0xc5, 0xcf, 0x97, 0xcd, 0x76, 0x0c,
	0x72, 0x9c, 0x87, 0x28, 0xe2, 0x09, 0xb6, 0xa1, 0x99, 0x9f, 0x03, 0x31, 0x9c, 0x60, 0x39, 0x4f,
	0xa9, 0xd0, 0x80, 0xe8, 0xdf, 0x4b, 0xc8, 0xec, 0xb5, 0xd5, 0x9c, 0x6a, 0x8b, 0x96, 0x03, 0xbb,
	0x26, 0xdf, 0xba, 0x09, 0x39, 0xb0, 0x0d, 0xf9, 0x3b, 0x4f, 0x75, 0x34, 0x18, 0xf3, 0x3c, 0x9b,
	0xce, 0x4b, 0xf5, 0xff, 0xff, 0x5e, 0xbd, 0x97, 0x90, 0xd9, 0xb1, 0x96, 0xac, 0x89, 0x81, 0x6d,
	0x88, 0xab, 0x37, 0x21, 0x06, 0x76, 0x45, 0xfc, 0xd2, 0x6b, 0xa8, 0x37, 0x16, 0x79, 0x28, 0xa2,
	0x0c, 0x52, 0x35, 0xff, 0x5f, 0xd9, 0x0b, 0xbf, 0xd6, 0x72, 0xdb, 0x95, 0xbe, 0x9f, 0x90, 0xd9,
	0xd9, 0xda, 0x81, 0x32, 0x34, 0x43, 0x03, 0xfb, 0x13, 0xbd, 0x6d, 0x69, 0x60, 0xbf, 0xa7, 0x0f,
	0xbd, 0x07, 0xd7, 0xdc, 0x2a, 0x01, 0xe1, 0xef, 0x68, 0x72, 0x7f, 0xc3, 0xab, 0x1a, 0x37, 0xd4,
	0xa6, 0xd3, 0x50, 0x9e, 0xa5, 0xae, 0xfa, 0x0c, 0x75, 0xec, 0xdd, 0x16, 0x92, 0x4c, 0x80, 0xc5,
	0x03, 0x31, 0x26, 0x19, 0xf5, 0x77, 0x5b, 0x6e, 0x7b, 0xa7, 0xf7, 0x54, 0x25, 0xf8, 0xed, 0xa2,
	0xf9, 0xc8, 0xe4, 0x25, 0x86, 0x13, 0x04, 0x1c, 0x27, 0x44, 0x8e, 0xd1, 0x09, 0x8d, 0x49, 0x34,
	0x3f, 0xa2, 0x51, 0xff, 0x96, 0x25, 0xcf, 0x14, 0xd8, 0x7b, 0xb3, 0xf8, 0x11, 0x38, 0x9f, 0x96,
	0x81, 0xb3, 0x58, 0x06, 0xee, 0xf9, 0x32, 0x70, 0xbf, 0x2f, 0x03, 0xf7, 0xc3, 0x2a, 0x70, 0xce,
	0x57, 0x81, 0xf3, 0x75, 0x15, 0x38, 0x6f, 0x0f, 0xd7, 0xc6, 0x51, 0x5e, 0x23, 0x07, 0x7c, 0x34,
	0x82, 0x08, 0xc8, 0x14, 0x8f, 0xf3, 0x10, 0x17, 0x9d, 0x2e, 0x9e, 0x99, 0x8b, 0x47, 0x4f, 0x07,
	0x17, 0xdd, 0xb0, 0xaa, 0x3f, 0xd1, 0xe7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x40, 0xda, 0xda,
	0x42, 0x99, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.StakingShare.Size()
		i -= size
		if _, err := m.StakingShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.MinSubscriptionHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSubscriptionHours))
		i--
		dAtA[i] = 0x50
	}
	if m.MaxSubscriptionHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSubscriptionHours))
		i--
		dAtA[i] = 0x48
	}
	if m.MinSubscriptionGigabytes != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSubscriptionGigabytes))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxSubscriptionGigabytes != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSubscriptionGigabytes))
		i--
		dAtA[i] = 0x38
	}
	if len(m.MinHourlyPrices) > 0 {
		for iNdEx := len(m.MinHourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinHourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.MaxHourlyPrices) > 0 {
		for iNdEx := len(m.MaxHourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxHourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MinGigabytePrices) > 0 {
		for iNdEx := len(m.MinGigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinGigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MaxGigabytePrices) > 0 {
		for iNdEx := len(m.MaxGigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxGigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ActiveDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	if len(m.MaxGigabytePrices) > 0 {
		for _, e := range m.MaxGigabytePrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinGigabytePrices) > 0 {
		for _, e := range m.MinGigabytePrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MaxHourlyPrices) > 0 {
		for _, e := range m.MaxHourlyPrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinHourlyPrices) > 0 {
		for _, e := range m.MinHourlyPrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxSubscriptionGigabytes != 0 {
		n += 1 + sovParams(uint64(m.MaxSubscriptionGigabytes))
	}
	if m.MinSubscriptionGigabytes != 0 {
		n += 1 + sovParams(uint64(m.MinSubscriptionGigabytes))
	}
	if m.MaxSubscriptionHours != 0 {
		n += 1 + sovParams(uint64(m.MaxSubscriptionHours))
	}
	if m.MinSubscriptionHours != 0 {
		n += 1 + sovParams(uint64(m.MinSubscriptionHours))
	}
	l = m.StakingShare.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGigabytePrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxGigabytePrices = append(m.MaxGigabytePrices, types.Coin{})
			if err := m.MaxGigabytePrices[len(m.MaxGigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGigabytePrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinGigabytePrices = append(m.MinGigabytePrices, types.Coin{})
			if err := m.MinGigabytePrices[len(m.MinGigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHourlyPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxHourlyPrices = append(m.MaxHourlyPrices, types.Coin{})
			if err := m.MaxHourlyPrices[len(m.MaxHourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinHourlyPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinHourlyPrices = append(m.MinHourlyPrices, types.Coin{})
			if err := m.MinHourlyPrices[len(m.MinHourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubscriptionGigabytes", wireType)
			}
			m.MaxSubscriptionGigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubscriptionGigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSubscriptionGigabytes", wireType)
			}
			m.MinSubscriptionGigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSubscriptionGigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubscriptionHours", wireType)
			}
			m.MaxSubscriptionHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubscriptionHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSubscriptionHours", wireType)
			}
			m.MinSubscriptionHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSubscriptionHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
