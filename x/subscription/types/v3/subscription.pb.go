// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v3/subscription.proto

package v3

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	v1 "github.com/sentinel-official/hub/v12/types/v1"
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

type Subscription struct {
	ID         uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccAddress string     `protobuf:"bytes,2,opt,name=acc_address,json=accAddress,proto3" json:"acc_address,omitempty"`
	PlanID     uint64     `protobuf:"varint,3,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Price      types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	Status     v1.Status  `protobuf:"varint,5,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	InactiveAt time.Time  `protobuf:"bytes,6,opt,name=inactive_at,json=inactiveAt,proto3,stdtime" json:"inactive_at"`
	RenewalAt  time.Time  `protobuf:"bytes,7,opt,name=renewal_at,json=renewalAt,proto3,stdtime" json:"renewal_at"`
	StatusAt   time.Time  `protobuf:"bytes,8,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922b6eb0c5993a9, []int{0}
}
func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return m.Size()
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Subscription)(nil), "sentinel.subscription.v3.Subscription")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v3/subscription.proto", fileDescriptor_d922b6eb0c5993a9)
}

var fileDescriptor_d922b6eb0c5993a9 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3f, 0x8f, 0xd3, 0x30,
	0x1c, 0x8d, 0x7b, 0xbd, 0x5c, 0xeb, 0x22, 0x06, 0x0b, 0xa1, 0x5c, 0x07, 0xa7, 0x82, 0xa5, 0x12,
	0xc2, 0x56, 0x5a, 0xb1, 0xb1, 0xa4, 0x77, 0x0c, 0xdd, 0x50, 0x8e, 0x89, 0x81, 0xca, 0x71, 0xdc,
	0x9c, 0xa5, 0x34, 0x8e, 0x62, 0x27, 0xc0, 0xb7, 0xb8, 0x2f, 0x81, 0xc4, 0x47, 0xe9, 0x78, 0x23,
	0x53, 0x81, 0xf4, 0x8b, 0xa0, 0xfc, 0x3b, 0xd1, 0xb1, 0x9b, 0xed, 0xdf, 0x7b, 0xcf, 0xef, 0x3d,
	0xfd, 0xe0, 0x1b, 0x2d, 0x52, 0x23, 0x53, 0x91, 0x50, 0x5d, 0x84, 0x9a, 0xe7, 0x32, 0x33, 0x52,
	0xa5, 0xb4, 0x5c, 0x9e, 0xdc, 0x49, 0x96, 0x2b, 0xa3, 0x90, 0xd3, 0x83, 0xc9, 0xc9, 0xb0, 0x5c,
	0x4e, 0x31, 0x57, 0x7a, 0xa7, 0x34, 0x0d, 0x99, 0x16, 0xb4, 0xf4, 0x42, 0x61, 0x98, 0x47, 0xb9,
	0x92, 0x1d, 0x73, 0xfa, 0x22, 0x56, 0xb1, 0x6a, 0x8e, 0xb4, 0x3e, 0x75, 0xaf, 0x6e, 0xac, 0x54,
	0x9c, 0x08, 0xda, 0xdc, 0xc2, 0x62, 0x4b, 0x8d, 0xdc, 0x09, 0x6d, 0xd8, 0x2e, 0xeb, 0x00, 0xf8,
	0xc9, 0x9d, 0xf9, 0x9e, 0x09, 0x4d, 0x4b, 0x8f, 0x6a, 0xc3, 0x4c, 0xa1, 0xdb, 0xf9, 0xab, 0x1f,
	0x17, 0xf0, 0xd9, 0xdd, 0x7f, 0x56, 0xd0, 0x4b, 0x38, 0x90, 0x91, 0x03, 0x66, 0x60, 0x3e, 0x5c,
	0xd9, 0xd5, 0xc1, 0x1d, 0xac, 0x6f, 0x83, 0x81, 0x8c, 0x90, 0x0b, 0x27, 0x8c, 0xf3, 0x0d, 0x8b,
	0xa2, 0x5c, 0x68, 0xed, 0x0c, 0x66, 0x60, 0x3e, 0x0e, 0x20, 0xe3, 0xdc, 0x6f, 0x5f, 0xd0, 0x6b,
	0x78, 0x95, 0x25, 0x2c, 0xdd, 0xc8, 0xc8, 0xb9, 0x68, 0xd8, 0xb0, 0x3a, 0xb8, 0xf6, 0xc7, 0x84,
	0xa5, 0xeb, 0xdb, 0xc0, 0xae, 0x47, 0xeb, 0x08, 0xbd, 0x83, 0x97, 0x59, 0x2e, 0xb9, 0x70, 0x86,
	0x33, 0x30, 0x9f, 0x2c, 0xae, 0x49, 0x9b, 0x9a, 0xd4, 0xa9, 0x49, 0x97, 0x9a, 0xdc, 0x28, 0x99,
	0xae, 0x86, 0xfb, 0x83, 0x6b, 0x05, 0x2d, 0x1a, 0x79, 0xd0, 0x6e, 0x5d, 0x3b, 0x97, 0x33, 0x30,
	0x7f, 0xbe, 0xb8, 0x26, 0x4f, 0x3d, 0x36, 0xb1, 0x48, 0xe9, 0x91, 0xbb, 0x06, 0x10, 0x74, 0x40,
	0xf4, 0x01, 0x4e, 0x64, 0xca, 0xb8, 0x91, 0xa5, 0xd8, 0x30, 0xe3, 0xd8, 0xcd, 0x7f, 0x53, 0xd2,
	0xf6, 0x45, 0xfa, 0xbe, 0xc8, 0xa7, 0xbe, 0xaf, 0xd5, 0xa8, 0xfe, 0xf0, 0xe1, 0xb7, 0x0b, 0x02,
	0xd8, 0x13, 0x7d, 0x83, 0x6e, 0x20, 0xcc, 0x45, 0x2a, 0xbe, 0xb2, 0xa4, 0x56, 0xb9, 0x3a, 0x43,
	0x65, 0xdc, 0xf1, 0x7c, 0x83, 0x7c, 0x38, 0x6e, 0x5d, 0xd5, 0x1a, 0xa3, 0x33, 0x34, 0x46, 0x2d,
	0xcd, 0x37, 0xab, 0x2f, 0xfb, 0xbf, 0xd8, 0xfa, 0x59, 0x61, 0x6b, 0x5f, 0x61, 0xf0, 0x58, 0x61,
	0xf0, 0xa7, 0xc2, 0xe0, 0xe1, 0x88, 0xad, 0xc7, 0x23, 0xb6, 0x7e, 0x1d, 0xb1, 0xf5, 0xf9, 0x7d,
	0x2c, 0xcd, 0x7d, 0x11, 0x12, 0xae, 0x76, 0xb4, 0x6f, 0xe7, 0xad, 0xda, 0x6e, 0x25, 0x97, 0x2c,
	0xa1, 0xf7, 0x45, 0x48, 0x4b, 0x6f, 0x41, 0xbf, 0x9d, 0x6e, 0x69, 0xb7, 0x14, 0xcb, 0xd0, 0x6e,
	0x7c, 0x2c, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x3f, 0x7d, 0x1f, 0xce, 0x02, 0x00, 0x00,
}

func (m *Subscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSubscription(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.RenewalAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.RenewalAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSubscription(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.InactiveAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintSubscription(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x32
	if m.Status != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSubscription(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PlanID != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AccAddress) > 0 {
		i -= len(m.AccAddress)
		copy(dAtA[i:], m.AccAddress)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.AccAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubscription(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubscription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Subscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSubscription(uint64(m.ID))
	}
	l = len(m.AccAddress)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanID != 0 {
		n += 1 + sovSubscription(uint64(m.PlanID))
	}
	l = m.Price.Size()
	n += 1 + l + sovSubscription(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSubscription(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt)
	n += 1 + l + sovSubscription(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.RenewalAt)
	n += 1 + l + sovSubscription(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSubscription(uint64(l))
	return n
}

func sovSubscription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubscription(x uint64) (n int) {
	return sovSubscription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Subscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: Subscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return fmt.Errorf("proto: wrong wireType = %d for field AccAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanID", wireType)
			}
			m.PlanID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= v1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.InactiveAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenewalAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.RenewalAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func skipSubscription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
				return 0, ErrInvalidLengthSubscription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubscription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubscription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubscription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubscription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubscription = fmt.Errorf("proto: unexpected end of group")
)
