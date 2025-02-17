// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v1/events.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	v1 "github.com/sentinel-official/hub/v12/types/v1"
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

type EventAddQuota struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventAddQuota) Reset()         { *m = EventAddQuota{} }
func (m *EventAddQuota) String() string { return proto.CompactTextString(m) }
func (*EventAddQuota) ProtoMessage()    {}
func (*EventAddQuota) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d3264879ff947c, []int{0}
}
func (m *EventAddQuota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAddQuota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAddQuota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAddQuota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddQuota.Merge(m, src)
}
func (m *EventAddQuota) XXX_Size() int {
	return m.Size()
}
func (m *EventAddQuota) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddQuota.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddQuota proto.InternalMessageInfo

type EventSetStatus struct {
	Id     uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Status v1.Status `protobuf:"varint,2,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
}

func (m *EventSetStatus) Reset()         { *m = EventSetStatus{} }
func (m *EventSetStatus) String() string { return proto.CompactTextString(m) }
func (*EventSetStatus) ProtoMessage()    {}
func (*EventSetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d3264879ff947c, []int{1}
}
func (m *EventSetStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetStatus.Merge(m, src)
}
func (m *EventSetStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventSetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetStatus proto.InternalMessageInfo

type EventSubscribe struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty" yaml:"node"`
	Plan uint64 `protobuf:"varint,3,opt,name=plan,proto3" json:"plan,omitempty" yaml:"plan"`
}

func (m *EventSubscribe) Reset()         { *m = EventSubscribe{} }
func (m *EventSubscribe) String() string { return proto.CompactTextString(m) }
func (*EventSubscribe) ProtoMessage()    {}
func (*EventSubscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d3264879ff947c, []int{2}
}
func (m *EventSubscribe) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSubscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSubscribe.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSubscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSubscribe.Merge(m, src)
}
func (m *EventSubscribe) XXX_Size() int {
	return m.Size()
}
func (m *EventSubscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSubscribe.DiscardUnknown(m)
}

var xxx_messageInfo_EventSubscribe proto.InternalMessageInfo

type EventUpdateQuota struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventUpdateQuota) Reset()         { *m = EventUpdateQuota{} }
func (m *EventUpdateQuota) String() string { return proto.CompactTextString(m) }
func (*EventUpdateQuota) ProtoMessage()    {}
func (*EventUpdateQuota) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d3264879ff947c, []int{3}
}
func (m *EventUpdateQuota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateQuota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateQuota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateQuota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateQuota.Merge(m, src)
}
func (m *EventUpdateQuota) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateQuota) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateQuota.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateQuota proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventAddQuota)(nil), "sentinel.subscription.v1.EventAddQuota")
	proto.RegisterType((*EventSetStatus)(nil), "sentinel.subscription.v1.EventSetStatus")
	proto.RegisterType((*EventSubscribe)(nil), "sentinel.subscription.v1.EventSubscribe")
	proto.RegisterType((*EventUpdateQuota)(nil), "sentinel.subscription.v1.EventUpdateQuota")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v1/events.proto", fileDescriptor_11d3264879ff947c)
}

var fileDescriptor_11d3264879ff947c = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0xed, 0x2c, 0x64, 0x44, 0x23, 0xd9, 0x66, 0x76, 0xf0, 0x02, 0x93, 0x83, 0xc6, 0x20,
	0x87, 0xcd, 0xc2, 0xd9, 0x6d, 0xec, 0x32, 0xb3, 0x3d, 0xc0, 0x1c, 0x7a, 0x29, 0xa5, 0xc5, 0x8e,
	0x94, 0x44, 0xe0, 0x58, 0x26, 0x92, 0x4d, 0xf3, 0x16, 0x7d, 0x8c, 0x3e, 0x4a, 0x8e, 0x39, 0xf6,
	0x64, 0x5a, 0xe7, 0x0d, 0xfc, 0x04, 0xc5, 0x92, 0x1d, 0xd2, 0x53, 0x2e, 0xbd, 0x09, 0x7d, 0xbf,
	0xef, 0xf7, 0x97, 0x3e, 0x3e, 0xf0, 0x4d, 0xd0, 0x44, 0xb2, 0x84, 0xc6, 0x58, 0x64, 0x91, 0x98,
	0x6f, 0x58, 0x2a, 0x19, 0x4f, 0x70, 0xee, 0x61, 0x9a, 0xd3, 0x44, 0x0a, 0x37, 0xdd, 0x70, 0xc9,
	0x2d, 0xbb, 0xc5, 0xdc, 0x53, 0xcc, 0xcd, 0xbd, 0xd1, 0xa7, 0x25, 0x5f, 0x72, 0x05, 0xe1, 0xfa,
	0xa4, 0xf9, 0x11, 0x3c, 0x6a, 0xe5, 0x36, 0xa5, 0xa2, 0xf6, 0x09, 0x19, 0xca, 0xac, 0xf1, 0xa1,
	0x2b, 0x30, 0xf8, 0x57, 0xfb, 0xff, 0x10, 0xf2, 0x3f, 0xe3, 0x32, 0xb4, 0xbe, 0x80, 0x0e, 0x23,
	0xb6, 0x39, 0x36, 0x27, 0x5d, 0x7f, 0x50, 0x15, 0x4e, 0x7f, 0x1b, 0xae, 0xe3, 0x5f, 0x88, 0x11,
	0x14, 0x74, 0x18, 0xb1, 0xbe, 0x83, 0xb7, 0x21, 0x21, 0x1b, 0x2a, 0x84, 0xdd, 0x19, 0x9b, 0x93,
	0xbe, 0x6f, 0x55, 0x85, 0x33, 0xd4, 0x4c, 0x53, 0x40, 0x41, 0x8b, 0xa0, 0x0c, 0x0c, 0x95, 0x7d,
	0x46, 0xe5, 0x4c, 0xa5, 0x9e, 0xd3, 0xff, 0x05, 0x3d, 0xfd, 0x3c, 0x65, 0x1f, 0x4e, 0x3f, 0xbb,
	0xc7, 0xff, 0xaa, 0xf7, 0xbb, 0xb9, 0xe7, 0x6a, 0x93, 0xff, 0xb1, 0x2a, 0x9c, 0x81, 0xee, 0xd6,
	0x2d, 0x28, 0x68, 0x7a, 0xd1, 0xb6, 0x8d, 0xd5, 0x23, 0x8a, 0xe8, 0xb9, 0xd8, 0xaf, 0xa0, 0x9b,
	0x70, 0x42, 0x9b, 0x2f, 0xbd, 0xaf, 0x0a, 0xe7, 0x9d, 0x06, 0xea, 0x5b, 0x14, 0xa8, 0x62, 0x0d,
	0xa5, 0x71, 0x98, 0xd8, 0x6f, 0x94, 0xe5, 0x04, 0xaa, 0x6f, 0x51, 0xa0, 0x8a, 0xe8, 0x06, 0x7c,
	0x50, 0xd1, 0x17, 0x29, 0x09, 0x25, 0x7d, 0xfd, 0x91, 0xfa, 0xd7, 0xbb, 0x27, 0x68, 0xdc, 0x97,
	0xd0, 0xd8, 0x95, 0xd0, 0xdc, 0x97, 0xd0, 0x7c, 0x2c, 0xa1, 0x79, 0x77, 0x80, 0xc6, 0xfe, 0x00,
	0x8d, 0x87, 0x03, 0x34, 0x2e, 0x7f, 0x2f, 0x99, 0x5c, 0x65, 0x91, 0x3b, 0xe7, 0x6b, 0xdc, 0x4e,
	0xef, 0x07, 0x5f, 0x2c, 0xd8, 0x9c, 0x85, 0x31, 0x5e, 0x65, 0x11, 0xce, 0xbd, 0x29, 0xbe, 0x7d,
	0xb9, 0x67, 0xed, 0x76, 0x44, 0x3d, 0xb5, 0x17, 0x3f, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4a,
	0x4b, 0x2e, 0xa8, 0x90, 0x02, 0x00, 0x00,
}

func (m *EventAddQuota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAddQuota) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAddQuota) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventSubscribe) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSubscribe) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSubscribe) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Plan != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Plan))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateQuota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateQuota) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateQuota) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventAddQuota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventSubscribe) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Plan != 0 {
		n += 1 + sovEvents(uint64(m.Plan))
	}
	return n
}

func (m *EventUpdateQuota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventAddQuota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAddQuota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAddQuota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSubscribe) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSubscribe: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSubscribe: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			m.Plan = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Plan |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateQuota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateQuota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateQuota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
