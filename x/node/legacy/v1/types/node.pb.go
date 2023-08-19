// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v1/node.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types1 "github.com/sentinel-official/hub/types"
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

type Node struct {
	Address   string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Provider  string                                   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Price     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	RemoteURL string                                   `protobuf:"bytes,4,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
	Status    types1.Status                            `protobuf:"varint,5,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt  time.Time                                `protobuf:"bytes,6,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db436afdcec2fda, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Node)(nil), "sentinel.node.v1.Node")
}

func init() { proto.RegisterFile("sentinel/node/v1/node.proto", fileDescriptor_3db436afdcec2fda) }

var fileDescriptor_3db436afdcec2fda = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0xbb, 0x8e, 0xd4, 0x30,
	0x14, 0x8d, 0xf7, 0x31, 0xcc, 0x78, 0x05, 0x42, 0x11, 0x45, 0x76, 0x90, 0x9c, 0x11, 0x55, 0x0a,
	0xd6, 0x26, 0x43, 0x41, 0xbd, 0x43, 0x8b, 0x28, 0x02, 0x5b, 0x40, 0xb3, 0x72, 0x12, 0x4f, 0xd6,
	0x22, 0xc9, 0x8d, 0x6c, 0x27, 0x62, 0xff, 0x62, 0x3f, 0x03, 0xf1, 0x1b, 0x34, 0x53, 0x6e, 0x49,
	0xb5, 0x0b, 0x99, 0x1f, 0x41, 0xb1, 0x93, 0x11, 0xd5, 0x7d, 0x9d, 0x7b, 0x75, 0xce, 0xb1, 0xf1,
	0x4b, 0x2d, 0x6a, 0x23, 0x6b, 0x51, 0xb2, 0x1a, 0x72, 0xc1, 0xba, 0xd8, 0x46, 0xda, 0x28, 0x30,
	0xe0, 0x3f, 0x9f, 0x86, 0xd4, 0x36, 0xbb, 0x78, 0x49, 0x32, 0xd0, 0x15, 0x68, 0x96, 0x72, 0x3d,
	0x80, 0x53, 0x61, 0x78, 0xcc, 0x32, 0x90, 0xb5, 0xdb, 0x58, 0xbe, 0x28, 0xa0, 0x00, 0x9b, 0xb2,
	0x21, 0x1b, 0xbb, 0x61, 0x01, 0x50, 0x94, 0x82, 0xd9, 0x2a, 0x6d, 0xb7, 0xcc, 0xc8, 0x4a, 0x68,
	0xc3, 0xab, 0x66, 0x04, 0x90, 0x03, 0x0b, 0x73, 0xdb, 0x08, 0x3d, 0xd0, 0xd0, 0x86, 0x9b, 0x56,
	0xbb, 0xf9, 0xab, 0x5f, 0x47, 0xf8, 0xe4, 0x23, 0xe4, 0xc2, 0x0f, 0xf0, 0x13, 0x9e, 0xe7, 0x4a,
	0x68, 0x1d, 0xa0, 0x15, 0x8a, 0x16, 0xc9, 0x54, 0xfa, 0x4b, 0x3c, 0x6f, 0x14, 0x74, 0x32, 0x17,
	0x2a, 0x38, 0xb2, 0xa3, 0x43, 0xed, 0x73, 0x7c, 0xda, 0x28, 0x99, 0x89, 0xe0, 0x78, 0x75, 0x1c,
	0x9d, 0xad, 0xcf, 0xa9, 0x53, 0x41, 0x07, 0x15, 0x74, 0x54, 0x41, 0xdf, 0x83, 0xac, 0x37, 0x6f,
	0x76, 0x0f, 0xa1, 0xf7, 0xf3, 0x31, 0x8c, 0x0a, 0x69, 0x6e, 0xda, 0x94, 0x66, 0x50, 0xb1, 0x51,
	0xb2, 0x0b, 0x17, 0x3a, 0xff, 0xe6, 0x48, 0xda, 0x05, 0x9d, 0xb8, 0xcb, 0xfe, 0x6b, 0x8c, 0x95,
	0xa8, 0xc0, 0x88, 0xeb, 0x56, 0x95, 0xc1, 0xc9, 0x40, 0x60, 0xf3, 0xb4, 0x7f, 0x08, 0x17, 0x89,
	0xed, 0x5e, 0x25, 0x1f, 0x92, 0x85, 0x03, 0x5c, 0xa9, 0xd2, 0x8f, 0xf1, 0xcc, 0xe9, 0x0b, 0x4e,
	0x57, 0x28, 0x7a, 0xb6, 0x3e, 0xa7, 0x07, 0xa7, 0xdd, 0xed, 0x2e, 0xa6, 0x9f, 0x2c, 0x20, 0x19,
	0x81, 0xfe, 0x25, 0x5e, 0xb8, 0xec, 0x9a, 0x9b, 0x60, 0xb6, 0x42, 0xd1, 0xd9, 0x7a, 0x49, 0x9d,
	0xaf, 0x74, 0xf2, 0x95, 0x7e, 0x9e, 0x7c, 0xdd, 0xcc, 0x07, 0x21, 0x77, 0x8f, 0x21, 0x4a, 0xe6,
	0x6e, 0xed, 0xd2, 0x6c, 0xbe, 0xec, 0xfe, 0x12, 0xef, 0x47, 0x4f, 0xbc, 0x5d, 0x4f, 0xd0, 0x7d,
	0x4f, 0xd0, 0x9f, 0x9e, 0xa0, 0xbb, 0x3d, 0xf1, 0xee, 0xf7, 0xc4, 0xfb, 0xbd, 0x27, 0xde, 0xd7,
	0x77, 0xff, 0xc9, 0x9e, 0x18, 0x5d, 0xc0, 0x76, 0x2b, 0x33, 0xc9, 0x4b, 0x76, 0xd3, 0xa6, 0xec,
	0xbb, 0xfb, 0x27, 0xa5, 0x28, 0x78, 0x76, 0x3b, 0xbc, 0x93, 0xe5, 0x9b, 0xce, 0x2c, 0x85, 0xb7,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x7f, 0xda, 0xcc, 0x4f, 0x02, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintNode(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if m.Status != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RemoteURL) > 0 {
		i -= len(m.RemoteURL)
		copy(dAtA[i:], m.RemoteURL)
		i = encodeVarintNode(dAtA, i, uint64(len(m.RemoteURL)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	l = len(m.RemoteURL)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovNode(uint64(m.Status))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovNode(uint64(l))
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
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
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)