// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/plan/v2/plan.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types1 "github.com/sentinel-official/hub/types"
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

type Plan struct {
	ID              uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProviderAddress string                                   `protobuf:"bytes,2,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty"`
	Duration        time.Duration                            `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration"`
	Gigabytes       int64                                    `protobuf:"varint,4,opt,name=gigabytes,proto3" json:"gigabytes,omitempty"`
	Prices          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=prices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"prices"`
	Status          types1.Status                            `protobuf:"varint,6,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt        time.Time                                `protobuf:"bytes,7,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42569898bb92ddb, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Plan)(nil), "sentinel.plan.v2.Plan")
}

func init() { proto.RegisterFile("sentinel/plan/v2/plan.proto", fileDescriptor_f42569898bb92ddb) }

var fileDescriptor_f42569898bb92ddb = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x8d, 0xdb, 0x12, 0xa6, 0x46, 0x82, 0x51, 0x84, 0x50, 0xa6, 0x20, 0x27, 0x62, 0x15, 0x16,
	0xb5, 0x49, 0x39, 0x00, 0x6a, 0x99, 0x0d, 0x2b, 0x50, 0x60, 0xc5, 0x66, 0xe4, 0x24, 0x6e, 0xc6,
	0x22, 0x8d, 0xa3, 0xd8, 0x89, 0x98, 0x35, 0x17, 0x98, 0x25, 0x47, 0x40, 0x9c, 0xa4, 0xcb, 0x59,
	0xb2, 0x9a, 0x42, 0x7a, 0x11, 0x54, 0x3b, 0x2e, 0x88, 0x59, 0xf9, 0xe7, 0xfd, 0xff, 0x7e, 0xde,
	0x7b, 0x36, 0x7c, 0x2a, 0x59, 0xa5, 0x78, 0xc5, 0x4a, 0x52, 0x97, 0xb4, 0x22, 0xdd, 0x42, 0x9f,
	0xb8, 0x6e, 0x84, 0x12, 0xde, 0xa9, 0x6d, 0x62, 0x0d, 0x76, 0x8b, 0x19, 0xca, 0x84, 0xdc, 0x08,
	0x49, 0x52, 0x2a, 0x19, 0xe9, 0xe2, 0x94, 0x29, 0x1a, 0x93, 0x4c, 0xf0, 0x81, 0x31, 0x7b, 0x5c,
	0x88, 0x42, 0xe8, 0x92, 0x1c, 0xaa, 0x01, 0x45, 0x85, 0x10, 0x45, 0xc9, 0x88, 0xfe, 0x4a, 0xdb,
	0x35, 0xc9, 0xdb, 0x86, 0x2a, 0x2e, 0x2c, 0x2b, 0xf8, 0xbf, 0xaf, 0xf8, 0x86, 0x49, 0x45, 0x37,
	0xb5, 0x5d, 0x70, 0x54, 0xa9, 0xae, 0x6a, 0x26, 0x49, 0x17, 0x13, 0xa9, 0xa8, 0x6a, 0xa5, 0xe9,
	0x3f, 0xff, 0x3a, 0x86, 0x93, 0xf7, 0x25, 0xad, 0xbc, 0x27, 0x70, 0xc4, 0x73, 0x1f, 0x84, 0x20,
	0x9a, 0xac, 0xdc, 0xfe, 0x36, 0x18, 0xbd, 0x3d, 0x4f, 0x46, 0x3c, 0xf7, 0x5e, 0xc0, 0xd3, 0xba,
	0x11, 0x1d, 0xcf, 0x59, 0x73, 0x41, 0xf3, 0xbc, 0x61, 0x52, 0xfa, 0xa3, 0x10, 0x44, 0xd3, 0xe4,
	0x91, 0xc5, 0x97, 0x06, 0xf6, 0x5e, 0xc3, 0x13, 0x2b, 0xcf, 0x1f, 0x87, 0x20, 0x7a, 0xb0, 0x38,
	0xc3, 0x46, 0x1f, 0xb6, 0xfa, 0xf0, 0xf9, 0x30, 0xb0, 0x3a, 0xd9, 0xde, 0x06, 0xce, 0xb7, 0x5d,
	0x00, 0x92, 0x23, 0xc9, 0x7b, 0x06, 0xa7, 0x05, 0x2f, 0x68, 0x7a, 0xa5, 0x98, 0xf4, 0x27, 0x21,
	0x88, 0xc6, 0xc9, 0x5f, 0xc0, 0xcb, 0xa0, 0x5b, 0x37, 0x3c, 0x63, 0xd2, 0xbf, 0x17, 0x8e, 0xf5,
	0x72, 0x13, 0x29, 0x3e, 0x44, 0x8a, 0x87, 0x48, 0xf1, 0x1b, 0xc1, 0xab, 0xd5, 0xcb, 0xc3, 0xf2,
	0x1f, 0xbb, 0x20, 0x2a, 0xb8, 0xba, 0x6c, 0x53, 0x9c, 0x89, 0x0d, 0x19, 0xf2, 0x37, 0xc7, 0x5c,
	0xe6, 0x9f, 0x4d, 0x22, 0x9a, 0x20, 0x93, 0x61, 0xb5, 0x17, 0x43, 0xd7, 0xe4, 0xe3, 0xbb, 0x21,
	0x88, 0x1e, 0x2e, 0xce, 0xf0, 0xf1, 0x26, 0xcd, 0x78, 0x17, 0xe3, 0x0f, 0x7a, 0x20, 0x19, 0x06,
	0xbd, 0x25, 0x9c, 0x9a, 0xea, 0x82, 0x2a, 0xff, 0xbe, 0xf6, 0x3d, 0xbb, 0xe3, 0xfb, 0xa3, 0xbd,
	0x17, 0x63, 0xfc, 0x5a, 0x1b, 0x37, 0xb4, 0xa5, 0x5a, 0xbd, 0xdb, 0xfe, 0x46, 0xce, 0xf7, 0x1e,
	0x39, 0xdb, 0x1e, 0x81, 0x9b, 0x1e, 0x81, 0x5f, 0x3d, 0x02, 0xd7, 0x7b, 0xe4, 0xdc, 0xec, 0x91,
	0xf3, 0x73, 0x8f, 0x9c, 0x4f, 0xf3, 0x7f, 0x9c, 0x58, 0x45, 0x73, 0xb1, 0x5e, 0xf3, 0x8c, 0xd3,
	0x92, 0x5c, 0xb6, 0x29, 0xf9, 0x62, 0xde, 0xa1, 0x56, 0x99, 0xba, 0xfa, 0xc7, 0xaf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x92, 0x13, 0x3c, 0x9e, 0xa5, 0x02, 0x00, 0x00,
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPlan(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if m.Status != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Gigabytes != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Gigabytes))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintPlan(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.ProviderAddress) > 0 {
		i -= len(m.ProviderAddress)
		copy(dAtA[i:], m.ProviderAddress)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.ProviderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPlan(uint64(m.ID))
	}
	l = len(m.ProviderAddress)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovPlan(uint64(l))
	if m.Gigabytes != 0 {
		n += 1 + sovPlan(uint64(m.Gigabytes))
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovPlan(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovPlan(uint64(l))
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gigabytes", wireType)
			}
			m.Gigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, types.Coin{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
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
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
