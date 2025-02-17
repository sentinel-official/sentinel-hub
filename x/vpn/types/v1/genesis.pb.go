// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/vpn/v1/genesis.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	v1 "github.com/sentinel-official/hub/v12/x/deposit/types/v1"
	v11 "github.com/sentinel-official/hub/v12/x/lease/types/v1"
	v3 "github.com/sentinel-official/hub/v12/x/node/types/v3"
	v31 "github.com/sentinel-official/hub/v12/x/plan/types/v3"
	v32 "github.com/sentinel-official/hub/v12/x/provider/types/v3"
	v33 "github.com/sentinel-official/hub/v12/x/session/types/v3"
	v34 "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GenesisState struct {
	Deposit      *v1.GenesisState  `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit,omitempty"`
	Lease        *v11.GenesisState `protobuf:"bytes,2,opt,name=lease,proto3" json:"lease,omitempty"`
	Node         *v3.GenesisState  `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Plan         *v31.GenesisState `protobuf:"bytes,4,opt,name=plan,proto3" json:"plan,omitempty"`
	Provider     *v32.GenesisState `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	Session      *v33.GenesisState `protobuf:"bytes,6,opt,name=session,proto3" json:"session,omitempty"`
	Subscription *v34.GenesisState `protobuf:"bytes,7,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_255b384044b0bea0, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "sentinel.vpn.v1.GenesisState")
}

func init() { proto.RegisterFile("sentinel/vpn/v1/genesis.proto", fileDescriptor_255b384044b0bea0) }

var fileDescriptor_255b384044b0bea0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xbd, 0x4e, 0xeb, 0x30,
	0x00, 0x85, 0x93, 0xdb, 0xbf, 0xab, 0xdc, 0x4a, 0x57, 0x8a, 0x18, 0xa2, 0x4a, 0xb8, 0xb4, 0x43,
	0xc5, 0x82, 0xad, 0x34, 0x62, 0x42, 0x62, 0x60, 0x41, 0x62, 0x04, 0xb1, 0xb0, 0x25, 0xad, 0x9b,
	0x5a, 0x0a, 0xb1, 0x15, 0xbb, 0x16, 0xbc, 0x05, 0x8f, 0xc1, 0xce, 0x4b, 0x74, 0xec, 0xc8, 0x08,
	0xe9, 0x8b, 0x20, 0xbb, 0xf9, 0x6b, 0x5d, 0xb6, 0xe8, 0x9c, 0xf3, 0x9d, 0x24, 0x27, 0xb1, 0x73,
	0xca, 0x71, 0x2a, 0x48, 0x8a, 0x13, 0x24, 0x59, 0x8a, 0xa4, 0x8f, 0x62, 0x9c, 0x62, 0x4e, 0x38,
	0x64, 0x19, 0x15, 0xd4, 0xfd, 0x5f, 0xda, 0x50, 0xb2, 0x14, 0x4a, 0x7f, 0x70, 0x12, 0xd3, 0x98,
	0x6a, 0x0f, 0xa9, 0xab, 0x5d, 0x6c, 0x30, 0xaa, 0x5a, 0xe6, 0x98, 0x51, 0x4e, 0x84, 0xd1, 0x34,
	0x18, 0x56, 0x91, 0x04, 0x87, 0x1c, 0x9b, 0x01, 0x50, 0x05, 0x52, 0x3a, 0xc7, 0x48, 0x06, 0xbf,
	0xfa, 0x2c, 0x09, 0x53, 0xd3, 0x1f, 0xd7, 0x7e, 0x46, 0x25, 0x99, 0xe3, 0xcc, 0xcc, 0xd4, 0xcf,
	0xc9, 0x31, 0xe7, 0x84, 0x1e, 0xa9, 0x99, 0xd4, 0x91, 0x55, 0xc4, 0x67, 0x19, 0x61, 0xe2, 0x58,
	0x6e, 0xfc, 0xd1, 0x72, 0xfa, 0xb7, 0x3b, 0xe5, 0x41, 0x84, 0x02, 0xbb, 0x57, 0x4e, 0xaf, 0x78,
	0x79, 0xcf, 0x3e, 0xb3, 0xcf, 0xff, 0x4d, 0x47, 0xb0, 0x1a, 0xaf, 0x30, 0xa0, 0xf4, 0x61, 0x93,
	0xb9, 0x2f, 0x09, 0xf7, 0xd2, 0xe9, 0xe8, 0x59, 0xbc, 0x3f, 0x1a, 0x1d, 0xd6, 0xa8, 0x96, 0x0d,
	0x70, 0x97, 0x76, 0xa7, 0x4e, 0x5b, 0x8d, 0xe5, 0xb5, 0x34, 0x05, 0x6a, 0x4a, 0xa9, 0x50, 0x06,
	0xfb, 0x90, 0xce, 0x2a, 0x46, 0x0d, 0xe8, 0xb5, 0x0f, 0x19, 0xa5, 0x9a, 0x8c, 0x52, 0xdd, 0x6b,
	0xe7, 0x6f, 0x39, 0xaa, 0xd7, 0xd1, 0xdc, 0xb8, 0xc1, 0x15, 0x8e, 0xc1, 0x56, 0x8c, 0xda, 0xa6,
	0x18, 0xdc, 0xeb, 0x1e, 0x6e, 0x53, 0x18, 0x06, 0x5d, 0x12, 0xee, 0x9d, 0xd3, 0x6f, 0x7e, 0x0a,
	0xaf, 0xa7, 0x1b, 0x26, 0x8d, 0x86, 0x86, 0x6b, 0xd4, 0xec, 0xb1, 0x37, 0x8f, 0xeb, 0x6f, 0x60,
	0xbd, 0xe7, 0xc0, 0x5a, 0xe7, 0xc0, 0xde, 0xe4, 0xc0, 0xfe, 0xca, 0x81, 0xfd, 0xb6, 0x05, 0xd6,
	0x66, 0x0b, 0xac, 0xcf, 0x2d, 0xb0, 0x9e, 0x82, 0x98, 0x88, 0xe5, 0x2a, 0x82, 0x33, 0xfa, 0x8c,
	0xca, 0x3b, 0x5c, 0xd0, 0xc5, 0x82, 0xcc, 0x48, 0x98, 0xa0, 0xe5, 0x2a, 0x42, 0xd2, 0x9f, 0xa2,
	0x17, 0x7d, 0x5c, 0xc4, 0x2b, 0xc3, 0x1c, 0x49, 0x3f, 0xea, 0xea, 0x7f, 0x22, 0xf8, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x9b, 0x5e, 0x22, 0xdf, 0x4e, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Subscription != nil {
		{
			size, err := m.Subscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Session != nil {
		{
			size, err := m.Session.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Provider != nil {
		{
			size, err := m.Provider.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Plan != nil {
		{
			size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Node != nil {
		{
			size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Lease != nil {
		{
			size, err := m.Lease.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Deposit != nil {
		{
			size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Deposit != nil {
		l = m.Deposit.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Lease != nil {
		l = m.Lease.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Plan != nil {
		l = m.Plan.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Provider != nil {
		l = m.Provider.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Session != nil {
		l = m.Session.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Subscription != nil {
		l = m.Subscription.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deposit == nil {
				m.Deposit = &v1.GenesisState{}
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lease == nil {
				m.Lease = &v11.GenesisState{}
			}
			if err := m.Lease.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v3.GenesisState{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Plan == nil {
				m.Plan = &v31.GenesisState{}
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Provider == nil {
				m.Provider = &v32.GenesisState{}
			}
			if err := m.Provider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Session == nil {
				m.Session = &v33.GenesisState{}
			}
			if err := m.Session.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subscription == nil {
				m.Subscription = &v34.GenesisState{}
			}
			if err := m.Subscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
