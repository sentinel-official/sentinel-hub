// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v3/session.proto

package v3

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	v1 "github.com/sentinel-official/hub/v12/types/v1"
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
	ID             uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccAddress     string                `protobuf:"bytes,2,opt,name=acc_address,json=accAddress,proto3" json:"acc_address,omitempty"`
	NodeAddress    string                `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	SubscriptionID uint64                `protobuf:"varint,4,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	DownloadBytes  cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=download_bytes,json=downloadBytes,proto3,customtype=cosmossdk.io/math.Int" json:"download_bytes"`
	UploadBytes    cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=upload_bytes,json=uploadBytes,proto3,customtype=cosmossdk.io/math.Int" json:"upload_bytes"`
	Duration       time.Duration         `protobuf:"bytes,7,opt,name=duration,proto3,stdduration" json:"duration"`
	Status         v1.Status             `protobuf:"varint,8,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt       time.Time             `protobuf:"bytes,9,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
	InactiveAt     time.Time             `protobuf:"bytes,10,opt,name=inactive_at,json=inactiveAt,proto3,stdtime" json:"inactive_at"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed776360b0097ec9, []int{0}
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
	proto.RegisterType((*Session)(nil), "sentinel.subscription.v3.Session")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v3/session.proto", fileDescriptor_ed776360b0097ec9)
}

var fileDescriptor_ed776360b0097ec9 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x8f, 0xd3, 0x3e,
	0x18, 0xc6, 0x93, 0xfe, 0xfb, 0xef, 0xb5, 0xee, 0x51, 0x24, 0x0b, 0x50, 0xae, 0x12, 0x49, 0x61,
	0x40, 0x5d, 0xb0, 0xd5, 0x76, 0x04, 0x09, 0x1a, 0x95, 0xa1, 0x6b, 0xca, 0xc4, 0x40, 0xe5, 0xc4,
	0x6e, 0x6a, 0xd1, 0xc4, 0x51, 0xed, 0x04, 0xee, 0x5b, 0xdc, 0xc8, 0x47, 0xb8, 0x8f, 0xd2, 0xf1,
	0x46, 0xc4, 0x50, 0x20, 0xfd, 0x22, 0x28, 0x71, 0x13, 0x05, 0x58, 0x60, 0x73, 0xfc, 0xfe, 0x9e,
	0xe7, 0x79, 0xed, 0xd7, 0x01, 0xcf, 0x24, 0x8b, 0x15, 0x8f, 0xd9, 0x0e, 0xcb, 0xd4, 0x97, 0xc1,
	0x9e, 0x27, 0x8a, 0x8b, 0x18, 0x67, 0x33, 0x2c, 0x99, 0x94, 0x5c, 0xc4, 0x28, 0xd9, 0x0b, 0x25,
	0xa0, 0x55, 0x71, 0xa8, 0xc9, 0xa1, 0x6c, 0x36, 0x7c, 0x10, 0x8a, 0x50, 0x94, 0x10, 0x2e, 0x56,
	0x9a, 0x1f, 0xda, 0xa1, 0x10, 0xe1, 0x8e, 0xe1, 0xf2, 0xcb, 0x4f, 0x37, 0x98, 0xa6, 0x7b, 0xa2,
	0x6a, 0xbf, 0xa1, 0xf3, 0x7b, 0x5d, 0xf1, 0x88, 0x49, 0x45, 0xa2, 0xa4, 0x32, 0xa8, 0x1b, 0x53,
	0xd7, 0x09, 0x93, 0x38, 0x9b, 0x60, 0xa9, 0x88, 0x4a, 0xa5, 0xae, 0x3f, 0xbd, 0x6d, 0x83, 0x8b,
	0x95, 0x6e, 0x11, 0x3e, 0x02, 0x2d, 0x4e, 0x2d, 0x73, 0x64, 0x8e, 0xdb, 0x6e, 0x27, 0x3f, 0x3a,
	0xad, 0xe5, 0xc2, 0x6b, 0x71, 0x0a, 0x1d, 0xd0, 0x27, 0x41, 0xb0, 0x26, 0x94, 0xee, 0x99, 0x94,
	0x56, 0x6b, 0x64, 0x8e, 0x7b, 0x1e, 0x20, 0x41, 0x30, 0xd7, 0x3b, 0xf0, 0x09, 0xb8, 0x8c, 0x05,
	0x65, 0x35, 0xf1, 0x5f, 0x49, 0xf4, 0x8b, 0xbd, 0x0a, 0x79, 0x01, 0xee, 0x37, 0x4f, 0xbc, 0xe6,
	0xd4, 0x6a, 0x97, 0x41, 0x30, 0x3f, 0x3a, 0x83, 0x55, 0xa3, 0xb4, 0x5c, 0x78, 0x83, 0x26, 0xba,
	0xa4, 0x70, 0x01, 0x06, 0x54, 0x7c, 0x8c, 0x77, 0x82, 0xd0, 0xb5, 0x7f, 0xad, 0x98, 0xb4, 0xfe,
	0x2f, 0x12, 0xdc, 0xc7, 0x87, 0xa3, 0x63, 0x7c, 0x3d, 0x3a, 0x0f, 0x03, 0x21, 0x23, 0x21, 0x25,
	0xfd, 0x80, 0xb8, 0xc0, 0x11, 0x51, 0x5b, 0xb4, 0x8c, 0x95, 0x77, 0xaf, 0x12, 0xb9, 0x85, 0x06,
	0xbe, 0x06, 0x97, 0x69, 0xd2, 0xf0, 0xe8, 0xfc, 0x8d, 0x47, 0x5f, 0x4b, 0xb4, 0xc3, 0x2b, 0xd0,
	0xad, 0xee, 0xdf, 0xba, 0x18, 0x99, 0xe3, 0xfe, 0xf4, 0x0a, 0xe9, 0x01, 0xa0, 0x6a, 0x00, 0x68,
	0x71, 0x06, 0xdc, 0x6e, 0x61, 0xfc, 0xf9, 0x9b, 0x63, 0x7a, 0xb5, 0x08, 0x4e, 0x40, 0x47, 0xdf,
	0xbe, 0xd5, 0x1d, 0x99, 0xe3, 0xc1, 0xf4, 0x0a, 0xd5, 0xef, 0xa1, 0x1c, 0x0f, 0xca, 0x26, 0x68,
	0x55, 0x02, 0xde, 0x19, 0x84, 0x73, 0xd0, 0xd3, 0xab, 0x35, 0x51, 0x56, 0xaf, 0x0c, 0x1d, 0xfe,
	0x11, 0xfa, 0xb6, 0x9a, 0xba, 0x4e, 0xbd, 0x29, 0x53, 0xb5, 0x6c, 0xae, 0xe0, 0x1b, 0xd0, 0xe7,
	0x31, 0x09, 0x14, 0xcf, 0x58, 0x61, 0x02, 0xfe, 0xc1, 0x04, 0x54, 0xc2, 0xb9, 0x72, 0xdf, 0x1f,
	0x7e, 0xd8, 0xc6, 0x6d, 0x6e, 0x1b, 0x87, 0xdc, 0x36, 0xef, 0x72, 0xdb, 0xfc, 0x9e, 0xdb, 0xe6,
	0xcd, 0xc9, 0x36, 0xee, 0x4e, 0xb6, 0xf1, 0xe5, 0x64, 0x1b, 0xef, 0x5e, 0x86, 0x5c, 0x6d, 0x53,
	0x1f, 0x05, 0x22, 0xc2, 0xd5, 0xc1, 0x9e, 0x8b, 0xcd, 0x86, 0x07, 0x9c, 0xec, 0xf0, 0x36, 0xf5,
	0x71, 0x36, 0x99, 0xe2, 0x4f, 0xbf, 0xfe, 0x23, 0xe7, 0x77, 0x39, 0xf3, 0x3b, 0x65, 0x27, 0xb3,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xb2, 0xea, 0x05, 0x4c, 0x03, 0x00, 0x00,
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
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.InactiveAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSession(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSession(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	if m.Status != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintSession(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x3a
	{
		size := m.UploadBytes.Size()
		i -= size
		if _, err := m.UploadBytes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.DownloadBytes.Size()
		i -= size
		if _, err := m.DownloadBytes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.SubscriptionID != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintSession(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccAddress) > 0 {
		i -= len(m.AccAddress)
		copy(dAtA[i:], m.AccAddress)
		i = encodeVarintSession(dAtA, i, uint64(len(m.AccAddress)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.AccAddress)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovSession(uint64(m.SubscriptionID))
	}
	l = m.DownloadBytes.Size()
	n += 1 + l + sovSession(uint64(l))
	l = m.UploadBytes.Size()
	n += 1 + l + sovSession(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovSession(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSession(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSession(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccAddress", wireType)
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
			m.AccAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownloadBytes", wireType)
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
			if err := m.DownloadBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadBytes", wireType)
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
			if err := m.UploadBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
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
				m.Status |= v1.Status(b&0x7F) << shift
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveAt", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.InactiveAt, dAtA[iNdEx:postIndex]); err != nil {
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