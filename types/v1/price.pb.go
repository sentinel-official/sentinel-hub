// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/types/v1/price.proto

package v1

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
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

type Price struct {
	Denom      string                      `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	BaseValue  cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=base_value,json=baseValue,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"base_value"`
	QuoteValue cosmossdk_io_math.Int       `protobuf:"bytes,3,opt,name=quote_value,json=quoteValue,proto3,customtype=cosmossdk.io/math.Int" json:"quote_value"`
}

func (m *Price) Reset()      { *m = Price{} }
func (*Price) ProtoMessage() {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28f60e170a28ffb, []int{0}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Price.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return m.Size()
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Price)(nil), "sentinel.types.v1.Price")
}

func init() { proto.RegisterFile("sentinel/types/v1/price.proto", fileDescriptor_f28f60e170a28ffb) }

var fileDescriptor_f28f60e170a28ffb = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4b, 0xcd, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x28,
	0xca, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x49, 0xeb, 0x81, 0xa5,
	0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1,
	0xd2, 0x42, 0x46, 0x2e, 0xd6, 0x00, 0x90, 0x46, 0x21, 0x11, 0x2e, 0xd6, 0x94, 0xd4, 0xbc, 0xfc,
	0x5c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0xc8, 0x89, 0x8b, 0x2b, 0x29, 0xb1,
	0x38, 0x35, 0xbe, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x24, 0xe5, 0xa4, 0x7c, 0xe2, 0x9e,
	0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0xd2, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xc5, 0x29, 0xd9, 0x7a,
	0x99, 0xf9, 0xfa, 0xb9, 0x89, 0x25, 0x19, 0x7a, 0x3e, 0xa9, 0xe9, 0x89, 0xc9, 0x95, 0x2e, 0xa9,
	0xc9, 0x41, 0x9c, 0x20, 0x6d, 0x61, 0x20, 0x5d, 0x42, 0x76, 0x5c, 0xdc, 0x85, 0xa5, 0xf9, 0x25,
	0x30, 0x43, 0x98, 0xc1, 0x86, 0xc8, 0x42, 0x0d, 0x11, 0xc5, 0x34, 0xc4, 0x33, 0xaf, 0x24, 0x88,
	0x0b, 0xac, 0x03, 0xac, 0xdf, 0x29, 0xf8, 0xc4, 0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x56,
	0x3c, 0x92, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xdd, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0xef, 0x75, 0xf3, 0xd3, 0xd2,
	0x32, 0x93, 0x33, 0x13, 0x73, 0xf4, 0x33, 0x4a, 0x93, 0xf4, 0xcb, 0x0c, 0x8d, 0xe0, 0xa1, 0x95,
	0xc4, 0x06, 0xf6, 0xbf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x58, 0xc0, 0x58, 0x49, 0x01,
	0x00, 0x00,
}

func (m *Price) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Price) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Price) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.QuoteValue.Size()
		i -= size
		if _, err := m.QuoteValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPrice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BaseValue.Size()
		i -= size
		if _, err := m.BaseValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPrice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Price) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	l = m.BaseValue.Size()
	n += 1 + l + sovPrice(uint64(l))
	l = m.QuoteValue.Size()
	n += 1 + l + sovPrice(uint64(l))
	return n
}

func sovPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrice(x uint64) (n int) {
	return sovPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Price) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrice
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
			return fmt.Errorf("proto: Price: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Price: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrice
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
func skipPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrice
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
					return 0, ErrIntOverflowPrice
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
					return 0, ErrIntOverflowPrice
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
				return 0, ErrInvalidLengthPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrice = fmt.Errorf("proto: unexpected end of group")
)