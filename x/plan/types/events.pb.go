// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/plan/v1/events.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types1 "github.com/sentinel-official/hub/types"
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

type EventAddPlan struct {
	From     string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Provider string                                   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       uint64                                   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Price    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	Validity time.Duration                            `protobuf:"bytes,5,opt,name=validity,proto3,stdduration" json:"validity"`
	Bytes    github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,6,opt,name=bytes,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bytes"`
}

func (m *EventAddPlan) Reset()         { *m = EventAddPlan{} }
func (m *EventAddPlan) String() string { return proto.CompactTextString(m) }
func (*EventAddPlan) ProtoMessage()    {}
func (*EventAddPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{0}
}
func (m *EventAddPlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAddPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAddPlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAddPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddPlan.Merge(m, src)
}
func (m *EventAddPlan) XXX_Size() int {
	return m.Size()
}
func (m *EventAddPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddPlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddPlan proto.InternalMessageInfo

type EventSetPlanStatus struct {
	From     string        `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Provider string        `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       uint64        `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Status   types1.Status `protobuf:"varint,4,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
}

func (m *EventSetPlanStatus) Reset()         { *m = EventSetPlanStatus{} }
func (m *EventSetPlanStatus) String() string { return proto.CompactTextString(m) }
func (*EventSetPlanStatus) ProtoMessage()    {}
func (*EventSetPlanStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{1}
}
func (m *EventSetPlanStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetPlanStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetPlanStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetPlanStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetPlanStatus.Merge(m, src)
}
func (m *EventSetPlanStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventSetPlanStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetPlanStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetPlanStatus proto.InternalMessageInfo

type EventAddNodeForPlan struct {
	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventAddNodeForPlan) Reset()         { *m = EventAddNodeForPlan{} }
func (m *EventAddNodeForPlan) String() string { return proto.CompactTextString(m) }
func (*EventAddNodeForPlan) ProtoMessage()    {}
func (*EventAddNodeForPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{2}
}
func (m *EventAddNodeForPlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAddNodeForPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAddNodeForPlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAddNodeForPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddNodeForPlan.Merge(m, src)
}
func (m *EventAddNodeForPlan) XXX_Size() int {
	return m.Size()
}
func (m *EventAddNodeForPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddNodeForPlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddNodeForPlan proto.InternalMessageInfo

type EventRemoveNodeForPlan struct {
	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventRemoveNodeForPlan) Reset()         { *m = EventRemoveNodeForPlan{} }
func (m *EventRemoveNodeForPlan) String() string { return proto.CompactTextString(m) }
func (*EventRemoveNodeForPlan) ProtoMessage()    {}
func (*EventRemoveNodeForPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{3}
}
func (m *EventRemoveNodeForPlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRemoveNodeForPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRemoveNodeForPlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRemoveNodeForPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRemoveNodeForPlan.Merge(m, src)
}
func (m *EventRemoveNodeForPlan) XXX_Size() int {
	return m.Size()
}
func (m *EventRemoveNodeForPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRemoveNodeForPlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventRemoveNodeForPlan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventAddPlan)(nil), "sentinel.plan.v1.EventAddPlan")
	proto.RegisterType((*EventSetPlanStatus)(nil), "sentinel.plan.v1.EventSetPlanStatus")
	proto.RegisterType((*EventAddNodeForPlan)(nil), "sentinel.plan.v1.EventAddNodeForPlan")
	proto.RegisterType((*EventRemoveNodeForPlan)(nil), "sentinel.plan.v1.EventRemoveNodeForPlan")
}

func init() { proto.RegisterFile("sentinel/plan/v1/events.proto", fileDescriptor_fed806c731daf270) }

var fileDescriptor_fed806c731daf270 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xe3, 0x4c, 0x5b, 0x3a, 0x06, 0x8d, 0x50, 0x40, 0x28, 0x53, 0x09, 0x37, 0xea, 0x02,
	0x65, 0x53, 0x9b, 0x0c, 0x07, 0x40, 0x94, 0x01, 0x89, 0x0d, 0xa0, 0x8c, 0xc4, 0x82, 0x9d, 0x13,
	0xbb, 0x19, 0x8b, 0x34, 0x8e, 0x62, 0x27, 0xa2, 0x27, 0x60, 0xcb, 0x92, 0x23, 0x20, 0x4e, 0xc0,
	0x11, 0xba, 0x9c, 0x25, 0x62, 0x31, 0x03, 0xed, 0x45, 0x50, 0xec, 0xa4, 0x02, 0x09, 0x09, 0x24,
	0x58, 0xc5, 0xcf, 0xf6, 0x7b, 0xff, 0xff, 0xbe, 0xe7, 0xc0, 0xbb, 0x8a, 0x17, 0x5a, 0x14, 0x3c,
	0x27, 0x65, 0x4e, 0x0b, 0xd2, 0x44, 0x84, 0x37, 0xbc, 0xd0, 0x0a, 0x97, 0x95, 0xd4, 0xd2, 0xbb,
	0xd9, 0x1f, 0xe3, 0xf6, 0x18, 0x37, 0xd1, 0x04, 0xa5, 0x52, 0xad, 0xa4, 0x22, 0x09, 0x55, 0x9c,
	0x34, 0x51, 0xc2, 0x35, 0x8d, 0x48, 0x2a, 0x45, 0x61, 0x33, 0x26, 0xb7, 0x33, 0x99, 0x49, 0xb3,
	0x24, 0xed, 0xaa, 0xdb, 0x45, 0x99, 0x94, 0x59, 0xce, 0x89, 0x89, 0x92, 0x7a, 0x49, 0x58, 0x5d,
	0x51, 0x2d, 0x64, 0x9f, 0x85, 0xf6, 0x36, 0xf4, 0xba, 0xe4, 0xaa, 0xf5, 0xa1, 0x34, 0xd5, 0x75,
	0xe7, 0x63, 0xf6, 0xd9, 0x85, 0x37, 0x9e, 0xb4, 0xc6, 0x1e, 0x31, 0xf6, 0x32, 0xa7, 0x85, 0xe7,
	0xc1, 0xc1, 0xb2, 0x92, 0x2b, 0x1f, 0x04, 0x20, 0x3c, 0x8c, 0xcd, 0xda, 0x9b, 0xc0, 0x71, 0x59,
	0xc9, 0x46, 0x30, 0x5e, 0xf9, 0xae, 0xd9, 0xdf, 0xc7, 0xde, 0x11, 0x74, 0x05, 0xf3, 0x0f, 0x02,
	0x10, 0x0e, 0x62, 0x57, 0x30, 0x8f, 0xc2, 0x61, 0x59, 0x89, 0x94, 0xfb, 0x83, 0xe0, 0x20, 0xbc,
	0x7e, 0x72, 0x8c, 0x6d, 0x5b, 0xb8, 0x6d, 0x0b, 0x77, 0x6d, 0xe1, 0xc7, 0x52, 0x14, 0x8b, 0xfb,
	0x9b, 0xcb, 0xa9, 0xf3, 0xe9, 0x6a, 0x1a, 0x66, 0x42, 0x9f, 0xd7, 0x09, 0x4e, 0xe5, 0x8a, 0x74,
	0x0c, 0xec, 0x67, 0xae, 0xd8, 0x1b, 0x6b, 0xdb, 0x24, 0xa8, 0xd8, 0x56, 0xf6, 0x1e, 0xc2, 0x71,
	0x43, 0x73, 0xc1, 0x84, 0x5e, 0xfb, 0xc3, 0x00, 0x18, 0x15, 0x8b, 0x01, 0xf7, 0x18, 0xf0, 0x69,
	0x87, 0x61, 0x31, 0x6e, 0x55, 0x3e, 0x5c, 0x4d, 0x41, 0xbc, 0x4f, 0xf2, 0x4e, 0xe1, 0x30, 0x59,
	0x6b, 0xae, 0xfc, 0x51, 0xdb, 0xcc, 0x02, 0xb7, 0x57, 0xbe, 0x5e, 0x4e, 0xef, 0xfd, 0x85, 0x91,
	0x67, 0x85, 0x8e, 0x6d, 0xf2, 0xec, 0x1d, 0x80, 0x9e, 0x41, 0x77, 0xc6, 0x75, 0x8b, 0xee, 0xcc,
	0x70, 0xfd, 0x67, 0x80, 0x11, 0x1c, 0xd9, 0x09, 0xf9, 0x83, 0x00, 0x84, 0x47, 0x27, 0xc7, 0x78,
	0xff, 0x54, 0xac, 0x85, 0x26, 0xc2, 0x56, 0x2a, 0xee, 0x2e, 0xce, 0x24, 0xbc, 0xd5, 0xcf, 0xf0,
	0xb9, 0x64, 0xfc, 0xa9, 0xac, 0xfe, 0xcb, 0x28, 0x7d, 0x78, 0x8d, 0x32, 0x56, 0x71, 0x65, 0xad,
	0x1c, 0xc6, 0x7d, 0x38, 0x7b, 0x05, 0xef, 0x18, 0xc1, 0x98, 0xaf, 0x64, 0xc3, 0xff, 0xa4, 0x69,
	0xeb, 0xba, 0xbf, 0xab, 0x7b, 0xf0, 0x4b, 0xdd, 0xc5, 0x8b, 0xcd, 0x77, 0xe4, 0x7c, 0xdc, 0x22,
	0x67, 0xb3, 0x45, 0xe0, 0x62, 0x8b, 0xc0, 0xb7, 0x2d, 0x02, 0xef, 0x77, 0xc8, 0xb9, 0xd8, 0x21,
	0xe7, 0xcb, 0x0e, 0x39, 0xaf, 0xe7, 0x3f, 0xcd, 0xa8, 0xe7, 0x32, 0x97, 0xcb, 0xa5, 0x48, 0x05,
	0xcd, 0xc9, 0x79, 0x9d, 0x90, 0xb7, 0xf6, 0x87, 0x33, 0xac, 0x92, 0x91, 0x79, 0x10, 0x0f, 0x7e,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xf8, 0x53, 0x6b, 0x8e, 0x03, 0x00, 0x00,
}

func (m *EventAddPlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAddPlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAddPlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Bytes.Size()
		i -= size
		if _, err := m.Bytes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Validity, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Validity):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvents(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetPlanStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetPlanStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetPlanStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAddNodeForPlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAddNodeForPlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAddNodeForPlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRemoveNodeForPlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRemoveNodeForPlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRemoveNodeForPlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
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
func (m *EventAddPlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Validity)
	n += 1 + l + sovEvents(uint64(l))
	l = m.Bytes.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventSetPlanStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventAddNodeForPlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRemoveNodeForPlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
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
func (m *EventAddPlan) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventAddPlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAddPlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Validity, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
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
			if err := m.Bytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *EventSetPlanStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventSetPlanStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetPlanStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
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
				m.Status |= types1.Status(b&0x7F) << shift
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
func (m *EventAddNodeForPlan) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventAddNodeForPlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAddNodeForPlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
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
func (m *EventRemoveNodeForPlan) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventRemoveNodeForPlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRemoveNodeForPlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
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
