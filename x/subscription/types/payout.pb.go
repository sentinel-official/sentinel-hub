// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/payout.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Payout struct {
	ID          uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address     string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	NodeAddress string     `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	Hours       int64      `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	Price       types.Coin `protobuf:"bytes,5,opt,name=price,proto3" json:"price"`
	Timestamp   time.Time  `protobuf:"bytes,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *Payout) Reset()         { *m = Payout{} }
func (m *Payout) String() string { return proto.CompactTextString(m) }
func (*Payout) ProtoMessage()    {}
func (*Payout) Descriptor() ([]byte, []int) {
	return fileDescriptor_f20b0a92fdfecfe5, []int{0}
}
func (m *Payout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payout.Merge(m, src)
}
func (m *Payout) XXX_Size() int {
	return m.Size()
}
func (m *Payout) XXX_DiscardUnknown() {
	xxx_messageInfo_Payout.DiscardUnknown(m)
}

var xxx_messageInfo_Payout proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Payout)(nil), "sentinel.subscription.v2.Payout")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v2/payout.proto", fileDescriptor_f20b0a92fdfecfe5)
}

var fileDescriptor_f20b0a92fdfecfe5 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x6a, 0xe3, 0x40,
	0x14, 0x85, 0x35, 0xb2, 0xad, 0x5d, 0x8f, 0xb7, 0x12, 0x66, 0xd1, 0xba, 0x18, 0x69, 0x17, 0x16,
	0xd4, 0xec, 0x0c, 0xf6, 0xe2, 0x07, 0x88, 0x92, 0x26, 0x5d, 0x10, 0x81, 0x40, 0x9a, 0xa0, 0x9f,
	0xb1, 0x3c, 0x60, 0xe9, 0x0a, 0xcd, 0xc8, 0xc4, 0x6f, 0xe1, 0xc7, 0xc8, 0xa3, 0xb8, 0x74, 0x99,
	0xca, 0x49, 0xe4, 0x3e, 0xcf, 0x10, 0x24, 0x59, 0xf9, 0xe9, 0xee, 0xb9, 0xf7, 0x3b, 0x70, 0x0e,
	0x17, 0xff, 0x95, 0x3c, 0x53, 0x22, 0xe3, 0x2b, 0x26, 0xcb, 0x50, 0x46, 0x85, 0xc8, 0x95, 0x80,
	0x8c, 0xad, 0x67, 0x2c, 0x0f, 0x36, 0x50, 0x2a, 0x9a, 0x17, 0xa0, 0xc0, 0xb4, 0x3a, 0x8c, 0x7e,
	0xc6, 0xe8, 0x7a, 0x36, 0x21, 0x11, 0xc8, 0x14, 0x24, 0x0b, 0x03, 0xc9, 0xd9, 0x7a, 0x1a, 0x72,
	0x15, 0x4c, 0x59, 0x04, 0x22, 0x6b, 0x9d, 0x93, 0x71, 0x02, 0x09, 0x34, 0x23, 0xab, 0xa7, 0xd3,
	0xd6, 0x4e, 0x00, 0x92, 0x15, 0x67, 0x8d, 0x0a, 0xcb, 0x05, 0x53, 0x22, 0xe5, 0x52, 0x05, 0x69,
	0xde, 0x02, 0x7f, 0x5e, 0x11, 0x36, 0xae, 0x9a, 0x04, 0xe6, 0x4f, 0xac, 0x8b, 0xd8, 0x42, 0x0e,
	0x72, 0xfb, 0x9e, 0x51, 0x1d, 0x6c, 0xfd, 0xf2, 0xc2, 0xd7, 0x45, 0x6c, 0x5a, 0xf8, 0x5b, 0x10,
	0xc7, 0x05, 0x97, 0xd2, 0xd2, 0x1d, 0xe4, 0x0e, 0xfd, 0x4e, 0x9a, 0xbf, 0xf1, 0x8f, 0x0c, 0x62,
	0x7e, 0xd7, 0x9d, 0x7b, 0xcd, 0x79, 0x54, 0xef, 0xce, 0x4e, 0xc8, 0x18, 0x0f, 0x96, 0x50, 0x16,
	0xd2, 0xea, 0x3b, 0xc8, 0xed, 0xf9, 0xad, 0x30, 0xe7, 0x78, 0x90, 0x17, 0x22, 0xe2, 0xd6, 0xc0,
	0x41, 0xee, 0x68, 0xf6, 0x8b, 0xb6, 0xe5, 0x68, 0x5d, 0x8e, 0x9e, 0xca, 0xd1, 0x73, 0x10, 0x99,
	0xd7, 0xdf, 0x1d, 0x6c, 0xcd, 0x6f, 0x69, 0xd3, 0xc3, 0xc3, 0xf7, 0xfc, 0x96, 0xd1, 0x58, 0x27,
	0xb4, 0x6d, 0x48, 0xbb, 0x86, 0xf4, 0xba, 0x23, 0xbc, 0xef, 0xb5, 0x77, 0xfb, 0x64, 0x23, 0xff,
	0xc3, 0xe6, 0xdd, 0xec, 0x5e, 0x88, 0xf6, 0x50, 0x11, 0x6d, 0x57, 0x11, 0xb4, 0xaf, 0x08, 0x7a,
	0xae, 0x08, 0xda, 0x1e, 0x89, 0xb6, 0x3f, 0x12, 0xed, 0xf1, 0x48, 0xb4, 0xdb, 0x79, 0x22, 0xd4,
	0xb2, 0x0c, 0x69, 0x04, 0x29, 0xeb, 0xde, 0xf1, 0x0f, 0x16, 0x0b, 0x11, 0x89, 0x60, 0xc5, 0x96,
	0x65, 0xc8, 0xee, 0xbf, 0x3e, 0x51, 0x6d, 0x72, 0x2e, 0x43, 0xa3, 0x49, 0xf0, 0xff, 0x2d, 0x00,
	0x00, 0xff, 0xff, 0xab, 0xad, 0x12, 0x0f, 0xea, 0x01, 0x00, 0x00,
}

func (m *Payout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPayout(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPayout(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Hours != 0 {
		i = encodeVarintPayout(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintPayout(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPayout(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintPayout(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPayout(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayout(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Payout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPayout(uint64(m.ID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPayout(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovPayout(uint64(l))
	}
	if m.Hours != 0 {
		n += 1 + sovPayout(uint64(m.Hours))
	}
	l = m.Price.Size()
	n += 1 + l + sovPayout(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovPayout(uint64(l))
	return n
}

func sovPayout(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPayout(x uint64) (n int) {
	return sovPayout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayout
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
			return fmt.Errorf("proto: Payout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
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
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
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
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
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
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
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
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayout
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
func skipPayout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayout
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
					return 0, ErrIntOverflowPayout
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
					return 0, ErrIntOverflowPayout
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
				return 0, ErrInvalidLengthPayout
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayout
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayout
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayout        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayout          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayout = fmt.Errorf("proto: unexpected end of group")
)