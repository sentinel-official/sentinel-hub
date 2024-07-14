// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/plan/v2/events.proto

package v2

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

type EventCreate struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	ID      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *EventCreate) Reset()         { *m = EventCreate{} }
func (m *EventCreate) String() string { return proto.CompactTextString(m) }
func (*EventCreate) ProtoMessage()    {}
func (*EventCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{0}
}
func (m *EventCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreate.Merge(m, src)
}
func (m *EventCreate) XXX_Size() int {
	return m.Size()
}
func (m *EventCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreate.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreate proto.InternalMessageInfo

type EventLinkNode struct {
	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	NodeAddress string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ID          uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *EventLinkNode) Reset()         { *m = EventLinkNode{} }
func (m *EventLinkNode) String() string { return proto.CompactTextString(m) }
func (*EventLinkNode) ProtoMessage()    {}
func (*EventLinkNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{1}
}
func (m *EventLinkNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventLinkNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventLinkNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventLinkNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLinkNode.Merge(m, src)
}
func (m *EventLinkNode) XXX_Size() int {
	return m.Size()
}
func (m *EventLinkNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLinkNode.DiscardUnknown(m)
}

var xxx_messageInfo_EventLinkNode proto.InternalMessageInfo

type EventUnlinkNode struct {
	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	NodeAddress string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ID          uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *EventUnlinkNode) Reset()         { *m = EventUnlinkNode{} }
func (m *EventUnlinkNode) String() string { return proto.CompactTextString(m) }
func (*EventUnlinkNode) ProtoMessage()    {}
func (*EventUnlinkNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{2}
}
func (m *EventUnlinkNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUnlinkNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUnlinkNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUnlinkNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUnlinkNode.Merge(m, src)
}
func (m *EventUnlinkNode) XXX_Size() int {
	return m.Size()
}
func (m *EventUnlinkNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUnlinkNode.DiscardUnknown(m)
}

var xxx_messageInfo_EventUnlinkNode proto.InternalMessageInfo

type EventUpdateStatus struct {
	Status  v1.Status `protobuf:"varint,1,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
	Address string    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	ID      uint64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *EventUpdateStatus) Reset()         { *m = EventUpdateStatus{} }
func (m *EventUpdateStatus) String() string { return proto.CompactTextString(m) }
func (*EventUpdateStatus) ProtoMessage()    {}
func (*EventUpdateStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{3}
}
func (m *EventUpdateStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateStatus.Merge(m, src)
}
func (m *EventUpdateStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateStatus proto.InternalMessageInfo

type EventCreateSubscription struct {
	Address         string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	ProviderAddress string `protobuf:"bytes,2,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty" yaml:"provider_address"`
	ID              uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	PlanID          uint64 `protobuf:"varint,4,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty" yaml:"plan_id"`
}

func (m *EventCreateSubscription) Reset()         { *m = EventCreateSubscription{} }
func (m *EventCreateSubscription) String() string { return proto.CompactTextString(m) }
func (*EventCreateSubscription) ProtoMessage()    {}
func (*EventCreateSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{4}
}
func (m *EventCreateSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateSubscription.Merge(m, src)
}
func (m *EventCreateSubscription) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateSubscription proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventCreate)(nil), "sentinel.plan.v2.EventCreate")
	proto.RegisterType((*EventLinkNode)(nil), "sentinel.plan.v2.EventLinkNode")
	proto.RegisterType((*EventUnlinkNode)(nil), "sentinel.plan.v2.EventUnlinkNode")
	proto.RegisterType((*EventUpdateStatus)(nil), "sentinel.plan.v2.EventUpdateStatus")
	proto.RegisterType((*EventCreateSubscription)(nil), "sentinel.plan.v2.EventCreateSubscription")
}

func init() { proto.RegisterFile("sentinel/plan/v2/events.proto", fileDescriptor_4222ab50472303c2) }

var fileDescriptor_4222ab50472303c2 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x3f, 0x6f, 0xd4, 0x30,
	0x1c, 0x8d, 0x43, 0x75, 0x55, 0x5d, 0xfa, 0x2f, 0x45, 0xba, 0xa3, 0x80, 0x53, 0x99, 0xa5, 0x03,
	0xc4, 0xba, 0x00, 0x4b, 0x37, 0x8e, 0x82, 0x54, 0x09, 0x21, 0x94, 0x0a, 0x06, 0x96, 0xe2, 0x3b,
	0xbb, 0x57, 0x8b, 0xd4, 0x8e, 0x12, 0x5f, 0x44, 0xbf, 0x05, 0x9f, 0x00, 0x31, 0x32, 0xf0, 0x41,
	0x3a, 0x76, 0x64, 0x8a, 0x20, 0x37, 0xb2, 0xe5, 0x13, 0xa0, 0xd8, 0x49, 0x75, 0xdc, 0x00, 0xdc,
	0xd6, 0xcd, 0xc9, 0x7b, 0xef, 0xe7, 0xe7, 0xdf, 0xfb, 0xfd, 0xe0, 0xbd, 0x8c, 0x4b, 0x2d, 0x24,
	0x8f, 0x49, 0x12, 0x53, 0x49, 0xf2, 0x90, 0xf0, 0x9c, 0x4b, 0x9d, 0x05, 0x49, 0xaa, 0xb4, 0xf2,
	0x36, 0x5b, 0x38, 0xa8, 0xe1, 0x20, 0x0f, 0x77, 0x6e, 0x8d, 0xd5, 0x58, 0x19, 0x90, 0xd4, 0x27,
	0xcb, 0xdb, 0x41, 0x57, 0x65, 0xf4, 0x79, 0xc2, 0x33, 0x92, 0xf7, 0x49, 0xa6, 0xa9, 0x9e, 0x34,
	0x75, 0xf0, 0x7b, 0xb8, 0xfa, 0xbc, 0xae, 0xfb, 0x2c, 0xe5, 0x54, 0x73, 0xef, 0x01, 0x5c, 0xa6,
	0x8c, 0xa5, 0x3c, 0xcb, 0x7a, 0x60, 0x17, 0xec, 0xad, 0x0c, 0xbc, 0xaa, 0xf0, 0xd7, 0xcf, 0xe9,
	0x59, 0xbc, 0x8f, 0x1b, 0x00, 0x47, 0x2d, 0xc5, 0xbb, 0x0f, 0x5d, 0xc1, 0x7a, 0xee, 0x2e, 0xd8,
	0x5b, 0x1a, 0x6c, 0x97, 0x85, 0xef, 0x1e, 0x1e, 0x54, 0x85, 0xbf, 0x62, 0xe9, 0x82, 0xe1, 0xc8,
	0x15, 0x0c, 0x7f, 0x06, 0x70, 0xcd, 0x5c, 0xf1, 0x52, 0xc8, 0x0f, 0xaf, 0x14, 0x5b, 0xf4, 0x92,
	0x7d, 0x78, 0x53, 0x2a, 0xc6, 0x8f, 0x5b, 0x89, 0x6b, 0x24, 0xdd, 0xaa, 0xf0, 0xb7, 0xad, 0x64,
	0x16, 0xc5, 0xd1, 0x6a, 0xfd, 0xf9, 0xf4, 0x0f, 0x83, 0x37, 0xfe, 0x6e, 0xf0, 0x0b, 0x80, 0x1b,
	0xc6, 0xe0, 0x1b, 0x19, 0x5f, 0x53, 0x8b, 0xdf, 0x00, 0xdc, 0xb2, 0x16, 0x13, 0x46, 0x35, 0x3f,
	0x32, 0x09, 0x7a, 0x07, 0xb0, 0x63, 0xb3, 0x34, 0x1e, 0xd7, 0xc3, 0xdb, 0xc1, 0xd5, 0x50, 0x98,
	0xb0, 0x83, 0xbc, 0x1f, 0x58, 0xea, 0x60, 0xab, 0x2a, 0xfc, 0x35, 0x5b, 0xd3, 0x4a, 0x70, 0xd4,
	0x68, 0x67, 0x9f, 0xea, 0xfe, 0x6f, 0xe4, 0xff, 0xb0, 0xfb, 0x0b, 0xc0, 0xee, 0xcc, 0x54, 0x1d,
	0x4d, 0x86, 0xd9, 0x28, 0x15, 0x89, 0x16, 0x4a, 0x2e, 0xd8, 0xd9, 0x17, 0x70, 0x33, 0x49, 0x55,
	0x2e, 0x18, 0x4f, 0xe7, 0xba, 0x7b, 0xa7, 0x2a, 0xfc, 0xae, 0x95, 0xcd, 0x33, 0x70, 0xb4, 0xd1,
	0xfe, 0x5a, 0xa4, 0xcb, 0xde, 0x13, 0xb8, 0x5c, 0x2f, 0xd3, 0xb1, 0x60, 0xbd, 0x25, 0xc3, 0xbc,
	0x5b, 0x16, 0x7e, 0xe7, 0x75, 0x4c, 0xa5, 0x61, 0x37, 0x26, 0x1b, 0x0a, 0x8e, 0x3a, 0xf5, 0xe9,
	0x90, 0x0d, 0xde, 0x5e, 0xfc, 0x44, 0xce, 0xd7, 0x12, 0x39, 0x17, 0x25, 0x02, 0x97, 0x25, 0x02,
	0x3f, 0x4a, 0x04, 0x3e, 0x4d, 0x91, 0x73, 0x39, 0x45, 0xce, 0xf7, 0x29, 0x72, 0xde, 0x3d, 0x1e,
	0x0b, 0x7d, 0x3a, 0x19, 0x06, 0x23, 0x75, 0x46, 0xda, 0x88, 0x1e, 0xaa, 0x93, 0x13, 0x31, 0x12,
	0x34, 0x26, 0xa7, 0x93, 0x21, 0xc9, 0xfb, 0x21, 0xf9, 0x68, 0x37, 0xbd, 0xd9, 0xd3, 0x70, 0xd8,
	0x31, 0x1b, 0xfa, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xd0, 0x3a, 0x50, 0x0a, 0x04,
	0x00, 0x00,
}

func (m *EventCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventLinkNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventLinkNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventLinkNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUnlinkNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUnlinkNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUnlinkNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PlanID))
		i--
		dAtA[i] = 0x20
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProviderAddress) > 0 {
		i -= len(m.ProviderAddress)
		copy(dAtA[i:], m.ProviderAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ProviderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *EventCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	return n
}

func (m *EventLinkNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	return n
}

func (m *EventUnlinkNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	return n
}

func (m *EventUpdateStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	return n
}

func (m *EventCreateSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ProviderAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	if m.PlanID != 0 {
		n += 1 + sovEvents(uint64(m.PlanID))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreate) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
func (m *EventLinkNode) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventLinkNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventLinkNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
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
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
func (m *EventUnlinkNode) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventUnlinkNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUnlinkNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
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
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
func (m *EventUpdateStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventUpdateStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
func (m *EventCreateSubscription) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventCreateSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddress", wireType)
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
			m.ProviderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanID", wireType)
			}
			m.PlanID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
