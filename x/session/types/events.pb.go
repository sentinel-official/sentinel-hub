// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/sentinel-official/hub/types"
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

type EventStart struct {
	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	NodeAddress    string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ID             uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	PlanID         uint64 `protobuf:"varint,4,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty" yaml:"plan_id"`
	SubscriptionID uint64 `protobuf:"varint,5,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty" yaml:"subscription_id"`
}

func (m *EventStart) Reset()         { *m = EventStart{} }
func (m *EventStart) String() string { return proto.CompactTextString(m) }
func (*EventStart) ProtoMessage()    {}
func (*EventStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_b22d66b3bd957014, []int{0}
}
func (m *EventStart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStart.Merge(m, src)
}
func (m *EventStart) XXX_Size() int {
	return m.Size()
}
func (m *EventStart) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStart.DiscardUnknown(m)
}

var xxx_messageInfo_EventStart proto.InternalMessageInfo

type EventUpdateDetails struct {
	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	NodeAddress    string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ID             uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	PlanID         uint64 `protobuf:"varint,4,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty" yaml:"plan_id"`
	SubscriptionID uint64 `protobuf:"varint,5,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty" yaml:"subscription_id"`
}

func (m *EventUpdateDetails) Reset()         { *m = EventUpdateDetails{} }
func (m *EventUpdateDetails) String() string { return proto.CompactTextString(m) }
func (*EventUpdateDetails) ProtoMessage()    {}
func (*EventUpdateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b22d66b3bd957014, []int{1}
}
func (m *EventUpdateDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateDetails.Merge(m, src)
}
func (m *EventUpdateDetails) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateDetails proto.InternalMessageInfo

type EventUpdateStatus struct {
	Status         types.Status `protobuf:"varint,1,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
	Address        string       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	NodeAddress    string       `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ID             uint64       `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	PlanID         uint64       `protobuf:"varint,5,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty" yaml:"plan_id"`
	SubscriptionID uint64       `protobuf:"varint,6,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty" yaml:"subscription_id"`
}

func (m *EventUpdateStatus) Reset()         { *m = EventUpdateStatus{} }
func (m *EventUpdateStatus) String() string { return proto.CompactTextString(m) }
func (*EventUpdateStatus) ProtoMessage()    {}
func (*EventUpdateStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b22d66b3bd957014, []int{2}
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

func init() {
	proto.RegisterType((*EventStart)(nil), "sentinel.session.v2.EventStart")
	proto.RegisterType((*EventUpdateDetails)(nil), "sentinel.session.v2.EventUpdateDetails")
	proto.RegisterType((*EventUpdateStatus)(nil), "sentinel.session.v2.EventUpdateStatus")
}

func init() { proto.RegisterFile("sentinel/session/v2/events.proto", fileDescriptor_b22d66b3bd957014) }

var fileDescriptor_b22d66b3bd957014 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xb1, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0xed, 0x6b, 0xea, 0xaa, 0x07, 0x04, 0xf5, 0x82, 0x20, 0x54, 0xe8, 0x1c, 0x1d, 0x4b,
	0x87, 0xe2, 0xa3, 0x41, 0x2c, 0xdd, 0x88, 0xc2, 0x90, 0x0d, 0x39, 0xea, 0xc2, 0x52, 0x5d, 0x72,
	0xd7, 0xf4, 0x24, 0xd7, 0x67, 0xe5, 0x2e, 0x16, 0x7d, 0x0b, 0x1e, 0x03, 0x46, 0xde, 0xa2, 0x63,
	0x47, 0x26, 0x0b, 0x9c, 0x37, 0xc8, 0xc0, 0x8c, 0x7c, 0x3e, 0x47, 0x16, 0x03, 0xa2, 0xea, 0xca,
	0x76, 0xbe, 0xef, 0xff, 0xfb, 0x6c, 0xfd, 0xfc, 0xe9, 0x83, 0x03, 0x2d, 0x52, 0x23, 0x53, 0x91,
	0x50, 0x2d, 0xb4, 0x96, 0x2a, 0xa5, 0xf9, 0x90, 0x8a, 0x5c, 0xa4, 0x46, 0x47, 0xd9, 0x52, 0x19,
	0x85, 0x7a, 0x4d, 0x22, 0x72, 0x89, 0x28, 0x1f, 0x1e, 0x3e, 0x59, 0xa8, 0x85, 0xb2, 0x75, 0x5a,
	0x9d, 0xea, 0xe8, 0x21, 0xde, 0x36, 0x33, 0xd7, 0x99, 0xd0, 0x34, 0x3f, 0xa1, 0xda, 0x30, 0xb3,
	0x72, 0xad, 0xc8, 0x57, 0x00, 0xe1, 0xfb, 0xaa, 0xf7, 0xd4, 0xb0, 0xa5, 0x41, 0xc7, 0x70, 0x8f,
	0x71, 0xbe, 0x14, 0x5a, 0xf7, 0xfd, 0x81, 0x7f, 0xb4, 0x3f, 0x42, 0x9b, 0x22, 0xec, 0x5e, 0xb3,
	0xab, 0xe4, 0x94, 0xb8, 0x02, 0x89, 0x9b, 0x08, 0x3a, 0x85, 0x0f, 0x53, 0xc5, 0xc5, 0x79, 0x83,
	0x00, 0x8b, 0x3c, 0xdb, 0x14, 0x61, 0xaf, 0x46, 0xda, 0x55, 0x12, 0x3f, 0xa8, 0x1e, 0xdf, 0x39,
	0xf6, 0x25, 0x04, 0x92, 0xf7, 0x77, 0x06, 0xfe, 0x51, 0x67, 0xd4, 0x2b, 0x8b, 0x10, 0x4c, 0xc6,
	0x9b, 0x22, 0xdc, 0xaf, 0x39, 0xc9, 0x49, 0x0c, 0x24, 0x47, 0x6f, 0xe1, 0x5e, 0x96, 0xb0, 0xf4,
	0x5c, 0xf2, 0x7e, 0xc7, 0x26, 0x5f, 0x94, 0x45, 0x18, 0x7c, 0x48, 0x58, 0x6a, 0xd3, 0xee, 0xc3,
	0x5c, 0x84, 0xc4, 0x41, 0x75, 0x9a, 0x70, 0x74, 0x06, 0x1f, 0xeb, 0xd5, 0x4c, 0xcf, 0x97, 0x32,
	0x33, 0x52, 0x59, 0x7c, 0xd7, 0xe2, 0xc7, 0x65, 0x11, 0x76, 0xa7, 0xad, 0x92, 0x6d, 0xf3, 0xb4,
	0x6e, 0xf3, 0x07, 0x42, 0xe2, 0x6e, 0xfb, 0x66, 0xc2, 0xc9, 0x37, 0x00, 0x91, 0x75, 0x75, 0x96,
	0x71, 0x66, 0xc4, 0x58, 0x18, 0x26, 0x13, 0xfd, 0xdf, 0xd9, 0xdf, 0x9c, 0xfd, 0x02, 0xf0, 0xa0,
	0xe5, 0x6c, 0x6a, 0x67, 0x0f, 0x8d, 0x61, 0x50, 0x4f, 0xa1, 0x35, 0xd6, 0x1d, 0x3e, 0x8f, 0xb6,
	0x13, 0x6d, 0xc7, 0x34, 0xca, 0x4f, 0xa2, 0x3a, 0x3a, 0x3a, 0xd8, 0x14, 0xe1, 0x23, 0xf7, 0x32,
	0x7b, 0x43, 0x62, 0xc7, 0xb6, 0xc5, 0x83, 0xbb, 0x8b, 0xdf, 0xb9, 0xb3, 0xf8, 0xce, 0x3f, 0x8b,
	0xdf, 0xbd, 0x9f, 0xf8, 0xe0, 0xfe, 0xe2, 0x47, 0xf1, 0xcd, 0x4f, 0xec, 0x7d, 0x29, 0xb1, 0x77,
	0x53, 0x62, 0xff, 0xb6, 0xc4, 0xfe, 0x8f, 0x12, 0xfb, 0x9f, 0xd7, 0xd8, 0xbb, 0x5d, 0x63, 0xef,
	0xfb, 0x1a, 0x7b, 0x1f, 0x5f, 0x2f, 0xa4, 0xb9, 0x5c, 0xcd, 0xa2, 0xb9, 0xba, 0xa2, 0x8d, 0xfe,
	0x57, 0xea, 0xe2, 0x42, 0xce, 0x25, 0x4b, 0xe8, 0xe5, 0x6a, 0x46, 0x3f, 0x6d, 0x37, 0x90, 0xfd,
	0x2b, 0xb3, 0xc0, 0xee, 0x8c, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x71, 0x03, 0x06,
	0xa2, 0x04, 0x00, 0x00,
}

func (m *EventStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubscriptionID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x28
	}
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

func (m *EventUpdateDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubscriptionID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x28
	}
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
	if m.SubscriptionID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x30
	}
	if m.PlanID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PlanID))
		i--
		dAtA[i] = 0x28
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
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
func (m *EventStart) Size() (n int) {
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
	if m.PlanID != 0 {
		n += 1 + sovEvents(uint64(m.PlanID))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovEvents(uint64(m.SubscriptionID))
	}
	return n
}

func (m *EventUpdateDetails) Size() (n int) {
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
	if m.PlanID != 0 {
		n += 1 + sovEvents(uint64(m.PlanID))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovEvents(uint64(m.SubscriptionID))
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
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	if m.PlanID != 0 {
		n += 1 + sovEvents(uint64(m.PlanID))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovEvents(uint64(m.SubscriptionID))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventStart) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStart: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
func (m *EventUpdateDetails) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventUpdateDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateDetails: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				m.Status |= types.Status(b&0x7F) << shift
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
		case 4:
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
		case 5:
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
