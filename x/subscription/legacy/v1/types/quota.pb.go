// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v1/quota.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Quota struct {
	Address   string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Allocated github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=allocated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"allocated"`
	Consumed  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=consumed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"consumed"`
}

func (m *Quota) Reset()         { *m = Quota{} }
func (m *Quota) String() string { return proto.CompactTextString(m) }
func (*Quota) ProtoMessage()    {}
func (*Quota) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e59de24126e6d2e, []int{0}
}
func (m *Quota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Quota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Quota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Quota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quota.Merge(m, src)
}
func (m *Quota) XXX_Size() int {
	return m.Size()
}
func (m *Quota) XXX_DiscardUnknown() {
	xxx_messageInfo_Quota.DiscardUnknown(m)
}

var xxx_messageInfo_Quota proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Quota)(nil), "sentinel.subscription.v1.Quota")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v1/quota.proto", fileDescriptor_3e59de24126e6d2e)
}

var fileDescriptor_3e59de24126e6d2e = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4b, 0xcd, 0xd1, 0x2f, 0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc,
	0xcf, 0xd3, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0xcd, 0x2f, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x80, 0xa9, 0xd2, 0x43, 0x56, 0xa5, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e,
	0x0f, 0x56, 0xa4, 0x0f, 0x62, 0x41, 0xd4, 0x2b, 0xed, 0x66, 0xe4, 0x62, 0x0d, 0x04, 0xe9, 0x17,
	0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0x71, 0x85, 0x7c, 0xb8, 0x38, 0x13, 0x73, 0x72, 0xf2, 0x93, 0x13, 0x4b, 0x52, 0x53,
	0x24, 0x98, 0x40, 0x72, 0x4e, 0x7a, 0x27, 0xee, 0xc9, 0x33, 0xdc, 0xba, 0x27, 0xaf, 0x96, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x0c,
	0xa5, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x3c, 0xf3, 0x4a, 0x82,
	0x10, 0x06, 0x08, 0x79, 0x71, 0x71, 0x24, 0xe7, 0xe7, 0x15, 0x97, 0xe6, 0xa6, 0xa6, 0x48, 0x30,
	0x93, 0x65, 0x18, 0x5c, 0xbf, 0x53, 0xe2, 0x89, 0x87, 0x72, 0x0c, 0x2b, 0x1e, 0xc9, 0x31, 0x9c,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x3d, 0x92, 0x99, 0xb0, 0xa0, 0xd1,
	0xcd, 0x4f, 0x4b, 0xcb, 0x4c, 0xce, 0x4c, 0xcc, 0xd1, 0xcf, 0x28, 0x4d, 0xd2, 0xaf, 0x40, 0x0d,
	0xcf, 0x9c, 0xd4, 0xf4, 0xc4, 0xe4, 0x4a, 0x50, 0xb0, 0x82, 0x2d, 0x4c, 0x62, 0x03, 0x87, 0x93,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x38, 0xc7, 0x9f, 0x7f, 0x01, 0x00, 0x00,
}

func (m *Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Quota) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Quota) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Consumed.Size()
		i -= size
		if _, err := m.Consumed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Allocated.Size()
		i -= size
		if _, err := m.Allocated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuota(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuota(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuota(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Quota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuota(uint64(l))
	}
	l = m.Allocated.Size()
	n += 1 + l + sovQuota(uint64(l))
	l = m.Consumed.Size()
	n += 1 + l + sovQuota(uint64(l))
	return n
}

func sovQuota(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuota(x uint64) (n int) {
	return sovQuota(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Quota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Allocated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Consumed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuota
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
func skipQuota(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
				return 0, ErrInvalidLengthQuota
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuota
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuota
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuota        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuota          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuota = fmt.Errorf("proto: unexpected end of group")
)
