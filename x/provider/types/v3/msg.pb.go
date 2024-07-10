// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/provider/v3/msg.proto

package v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	v2 "github.com/sentinel-official/hub/v12/x/provider/types/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgUpdateParamsRequest struct {
	From   string    `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Params v2.Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParamsRequest) Reset()         { *m = MsgUpdateParamsRequest{} }
func (m *MsgUpdateParamsRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsRequest) ProtoMessage()    {}
func (*MsgUpdateParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df9ce074f854f90c, []int{0}
}
func (m *MsgUpdateParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsRequest.Merge(m, src)
}
func (m *MsgUpdateParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsRequest proto.InternalMessageInfo

type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df9ce074f854f90c, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParamsRequest)(nil), "sentinel.provider.v3.MsgUpdateParamsRequest")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "sentinel.provider.v3.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("sentinel/provider/v3/msg.proto", fileDescriptor_df9ce074f854f90c) }

var fileDescriptor_df9ce074f854f90c = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x13, 0x19, 0x03, 0xe3, 0x41, 0x08, 0x43, 0xe7, 0x90, 0x38, 0x77, 0xda, 0xc1, 0x25,
	0xd8, 0x5e, 0xc4, 0xe3, 0xee, 0x03, 0x99, 0x78, 0xd9, 0xad, 0xdb, 0xb2, 0x34, 0xb0, 0x36, 0x31,
	0x49, 0x83, 0x82, 0x1f, 0xc2, 0x8f, 0xe1, 0x47, 0xe9, 0x71, 0x47, 0x4f, 0xa2, 0xed, 0x17, 0x11,
	0x5b, 0x8b, 0xa0, 0x3d, 0x78, 0xfb, 0xc3, 0xfb, 0xfd, 0xdf, 0x7b, 0x3c, 0x44, 0x2c, 0x4f, 0x9d,
	0x4c, 0xf9, 0x96, 0x69, 0xa3, 0xbc, 0x5c, 0x73, 0xc3, 0x7c, 0xc8, 0x12, 0x2b, 0xa8, 0x36, 0xca,
	0x29, 0xdc, 0x6b, 0x74, 0xda, 0xe8, 0xd4, 0x87, 0x83, 0x9e, 0x50, 0x42, 0x55, 0x00, 0xfb, 0xba,
	0x6a, 0x76, 0x70, 0xde, 0xe2, 0x15, 0x30, 0x1d, 0x99, 0x28, 0xb1, 0x35, 0x32, 0x8a, 0xd1, 0xd1,
	0xcc, 0x8a, 0x3b, 0xbd, 0x8e, 0x1c, 0xbf, 0xa9, 0x84, 0x39, 0xbf, 0xcf, 0xb8, 0x75, 0x18, 0xa3,
	0xce, 0xc6, 0xa8, 0xa4, 0x0f, 0x87, 0x70, 0xbc, 0x3f, 0xaf, 0x6e, 0x7c, 0x8d, 0xba, 0xf5, 0x77,
	0x7f, 0x6f, 0x08, 0xc7, 0x07, 0xc1, 0x29, 0x6d, 0x69, 0x13, 0xd0, 0xda, 0x68, 0xda, 0xc9, 0xdf,
	0xce, 0xc0, 0xfc, 0xfb, 0x63, 0x74, 0x82, 0x8e, 0xff, 0x24, 0x59, 0xad, 0x52, 0xcb, 0x83, 0x27,
	0x84, 0x66, 0x56, 0xdc, 0x72, 0xe3, 0xe5, 0x8a, 0xe3, 0x14, 0x1d, 0xfe, 0x02, 0xf1, 0x45, 0x5b,
	0x4e, 0x48, 0xdb, 0x9b, 0x0f, 0x26, 0xff, 0xa4, 0xeb, 0xf4, 0xe9, 0x22, 0xff, 0x20, 0xe0, 0xa5,
	0x20, 0x20, 0x2f, 0x08, 0xdc, 0x15, 0x04, 0xbe, 0x17, 0x04, 0x3e, 0x97, 0x04, 0xec, 0x4a, 0x02,
	0x5e, 0x4b, 0x02, 0x16, 0x57, 0x42, 0xba, 0x38, 0x5b, 0xd2, 0x95, 0x4a, 0x58, 0x63, 0x3d, 0x51,
	0x9b, 0x8d, 0x5c, 0xc9, 0x68, 0xcb, 0xe2, 0x6c, 0xc9, 0xfc, 0x65, 0xc0, 0x1e, 0x7e, 0x56, 0x76,
	0x8f, 0x9a, 0x5b, 0xe6, 0xc3, 0x65, 0xb7, 0x5a, 0x39, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x94, 0x3e, 0x2a, 0xd6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgUpdateParams(ctx context.Context, in *MsgUpdateParamsRequest, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgUpdateParams(ctx context.Context, in *MsgUpdateParamsRequest, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/sentinel.provider.v3.MsgService/MsgUpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgUpdateParams(context.Context, *MsgUpdateParamsRequest) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgUpdateParams(ctx context.Context, req *MsgUpdateParamsRequest) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpdateParams not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgUpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgUpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.provider.v3.MsgService/MsgUpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgUpdateParams(ctx, req.(*MsgUpdateParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.provider.v3.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgUpdateParams",
			Handler:    _MsgService_MsgUpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/provider/v3/msg.proto",
}

func (m *MsgUpdateParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)