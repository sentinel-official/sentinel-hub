// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/session.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/sentinel-official/hub/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type Session struct {
	ID             uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionID uint64          `protobuf:"varint,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	NodeAddress    string          `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	Address        string          `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Bandwidth      types.Bandwidth `protobuf:"bytes,5,opt,name=bandwidth,proto3" json:"bandwidth"`
	Duration       time.Duration   `protobuf:"bytes,6,opt,name=duration,proto3,stdduration" json:"duration"`
	ExpiryAt       time.Time       `protobuf:"bytes,7,opt,name=expiry_at,json=expiryAt,proto3,stdtime" json:"expiry_at"`
	Status         types.Status    `protobuf:"varint,8,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt       time.Time       `protobuf:"bytes,9,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa3a03f78d221e53, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Session)(nil), "sentinel.session.v2.Session")
}

func init() { proto.RegisterFile("sentinel/session/v2/session.proto", fileDescriptor_aa3a03f78d221e53) }

var fileDescriptor_aa3a03f78d221e53 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0xc7, 0x93, 0x5e, 0xe9, 0x8b, 0x0f, 0x15, 0xc9, 0x20, 0x94, 0xab, 0x90, 0xd3, 0x63, 0xea,
	0x82, 0x4d, 0xcb, 0xc8, 0x00, 0x8d, 0xba, 0x74, 0x4d, 0x99, 0x58, 0xaa, 0xa4, 0x76, 0x53, 0x4b,
	0x6d, 0x1c, 0xc5, 0x4e, 0xb9, 0xfb, 0x16, 0x37, 0xf2, 0x11, 0xf8, 0x28, 0x1d, 0x6f, 0x83, 0xa9,
	0x40, 0xfa, 0x45, 0x50, 0xec, 0x3a, 0x77, 0xa2, 0x2c, 0x6c, 0x8f, 0x9f, 0xff, 0xef, 0xf9, 0x3f,
	0x2f, 0x09, 0xb8, 0x96, 0x2c, 0x55, 0x3c, 0x65, 0x1b, 0x22, 0x99, 0x94, 0x5c, 0xa4, 0x64, 0x37,
	0xb6, 0x21, 0xce, 0x72, 0xa1, 0x04, 0x7c, 0x6e, 0x11, 0x6c, 0xf3, 0xbb, 0x71, 0xff, 0x45, 0x22,
	0x12, 0xa1, 0x75, 0x52, 0x45, 0x06, 0xed, 0xa3, 0x44, 0x88, 0x64, 0xc3, 0x88, 0x7e, 0xc5, 0xc5,
	0x8a, 0xd0, 0x22, 0x8f, 0x54, 0x6d, 0xd5, 0xf7, 0xff, 0xd6, 0x15, 0xdf, 0x32, 0xa9, 0xa2, 0x6d,
	0x76, 0x02, 0x1e, 0xc6, 0x51, 0xb7, 0x19, 0x93, 0x64, 0x37, 0x22, 0x71, 0x94, 0xd2, 0x2f, 0x9c,
	0xaa, 0xb5, 0xed, 0x71, 0x8e, 0x48, 0x15, 0xa9, 0x42, 0x1a, 0xfd, 0xf5, 0xf7, 0x0b, 0xd0, 0x9e,
	0x9b, 0x41, 0xe1, 0x4b, 0xd0, 0xe0, 0xd4, 0x73, 0x07, 0xee, 0xb0, 0x19, 0xb4, 0xca, 0x83, 0xdf,
	0x98, 0x4d, 0xc3, 0x06, 0xa7, 0xf0, 0x3d, 0x78, 0x26, 0x8b, 0x58, 0x2e, 0x73, 0x9e, 0x55, 0xd3,
	0x2d, 0x38, 0xf5, 0x1a, 0x1a, 0x82, 0xe5, 0xc1, 0xef, 0xcd, 0x1f, 0x49, 0xb3, 0x69, 0xd8, 0x7b,
	0x8c, 0xce, 0x28, 0xbc, 0x06, 0x4f, 0x53, 0x41, 0xd9, 0x22, 0xa2, 0x34, 0x67, 0x52, 0x7a, 0x17,
	0x03, 0x77, 0xd8, 0x0d, 0x2f, 0xab, 0xdc, 0xc4, 0xa4, 0xa0, 0x07, 0xda, 0x56, 0x6d, 0x6a, 0xd5,
	0x3e, 0xe1, 0x47, 0xd0, 0xad, 0x17, 0xf2, 0x9e, 0x0c, 0xdc, 0xe1, 0xe5, 0xf8, 0x15, 0xae, 0x0f,
	0xac, 0x37, 0xc2, 0xbb, 0x11, 0x0e, 0x2c, 0x13, 0x34, 0xf7, 0x07, 0xdf, 0x09, 0x1f, 0x8a, 0xe0,
	0x07, 0xd0, 0xb1, 0x57, 0xf5, 0x5a, 0xda, 0xe0, 0x0a, 0x9b, 0xb3, 0x62, 0x7b, 0x56, 0x3c, 0x3d,
	0x01, 0x41, 0xa7, 0xaa, 0xfe, 0xfa, 0xd3, 0x77, 0xc3, 0xba, 0x08, 0x4e, 0x40, 0x97, 0xdd, 0x64,
	0x3c, 0xbf, 0x5d, 0x44, 0xca, 0x6b, 0x6b, 0x87, 0xfe, 0x99, 0xc3, 0x27, 0xfb, 0x61, 0x8c, 0xc5,
	0x9d, 0xb6, 0x30, 0x65, 0x13, 0x05, 0x47, 0xa0, 0x65, 0x6e, 0xee, 0x75, 0x06, 0xee, 0xb0, 0x37,
	0xbe, 0xfa, 0xc7, 0x0a, 0x73, 0x0d, 0x84, 0x27, 0xb0, 0xea, 0x6a, 0xa2, 0xaa, 0x6b, 0xf7, 0x7f,
	0xba, 0x9a, 0xb2, 0x89, 0x0a, 0xc2, 0xfd, 0x6f, 0xe4, 0x7c, 0x2b, 0x91, 0xb3, 0x2f, 0x91, 0x7b,
	0x5f, 0x22, 0xf7, 0x57, 0x89, 0xdc, 0xbb, 0x23, 0x72, 0xee, 0x8f, 0xc8, 0xf9, 0x71, 0x44, 0xce,
	0xe7, 0xb7, 0x09, 0x57, 0xeb, 0x22, 0xc6, 0x4b, 0xb1, 0x25, 0x76, 0xa2, 0x37, 0x62, 0xb5, 0xe2,
	0x4b, 0x1e, 0x6d, 0xc8, 0xba, 0x88, 0xc9, 0x4d, 0xfd, 0x9f, 0xeb, 0x41, 0xe3, 0x96, 0xee, 0xfd,
	0xee, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb6, 0xd2, 0xf3, 0x08, 0x03, 0x00, 0x00,
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSession(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	if m.Status != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiryAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiryAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSession(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintSession(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSession(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintSession(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SubscriptionID != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	offset -= sovSession(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSession(uint64(m.ID))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovSession(uint64(m.SubscriptionID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = m.Bandwidth.Size()
	n += 1 + l + sovSession(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovSession(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiryAt)
	n += 1 + l + sovSession(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSession(uint64(m.Status))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSession(uint64(l))
	return n
}

func sovSession(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriptionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiryAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
				return 0, ErrInvalidLengthSession
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSession
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSession
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSession        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSession = fmt.Errorf("proto: unexpected end of group")
)
