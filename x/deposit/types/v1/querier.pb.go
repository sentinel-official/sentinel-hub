// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/deposit/v1/querier.proto

package v1

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryDepositsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDepositsRequest) Reset()         { *m = QueryDepositsRequest{} }
func (m *QueryDepositsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDepositsRequest) ProtoMessage()    {}
func (*QueryDepositsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60968570eeafa82b, []int{0}
}
func (m *QueryDepositsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositsRequest.Merge(m, src)
}
func (m *QueryDepositsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositsRequest proto.InternalMessageInfo

type QueryDepositRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryDepositRequest) Reset()         { *m = QueryDepositRequest{} }
func (m *QueryDepositRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDepositRequest) ProtoMessage()    {}
func (*QueryDepositRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60968570eeafa82b, []int{1}
}
func (m *QueryDepositRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositRequest.Merge(m, src)
}
func (m *QueryDepositRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositRequest proto.InternalMessageInfo

type QueryDepositsResponse struct {
	Deposits   []Deposit           `protobuf:"bytes,1,rep,name=deposits,proto3" json:"deposits"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDepositsResponse) Reset()         { *m = QueryDepositsResponse{} }
func (m *QueryDepositsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDepositsResponse) ProtoMessage()    {}
func (*QueryDepositsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60968570eeafa82b, []int{2}
}
func (m *QueryDepositsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositsResponse.Merge(m, src)
}
func (m *QueryDepositsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositsResponse proto.InternalMessageInfo

type QueryDepositResponse struct {
	Deposit Deposit `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
}

func (m *QueryDepositResponse) Reset()         { *m = QueryDepositResponse{} }
func (m *QueryDepositResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDepositResponse) ProtoMessage()    {}
func (*QueryDepositResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60968570eeafa82b, []int{3}
}
func (m *QueryDepositResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositResponse.Merge(m, src)
}
func (m *QueryDepositResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryDepositsRequest)(nil), "sentinel.deposit.v1.QueryDepositsRequest")
	proto.RegisterType((*QueryDepositRequest)(nil), "sentinel.deposit.v1.QueryDepositRequest")
	proto.RegisterType((*QueryDepositsResponse)(nil), "sentinel.deposit.v1.QueryDepositsResponse")
	proto.RegisterType((*QueryDepositResponse)(nil), "sentinel.deposit.v1.QueryDepositResponse")
}

func init() { proto.RegisterFile("sentinel/deposit/v1/querier.proto", fileDescriptor_60968570eeafa82b) }

var fileDescriptor_60968570eeafa82b = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x82, 0x18, 0x78, 0x70, 0xf1, 0x86, 0x34, 0x55, 0xc3, 0x83, 0x4a, 0xb0, 0x76,
	0x12, 0x7e, 0x4a, 0x39, 0x70, 0x41, 0x1c, 0x26, 0x04, 0x57, 0x28, 0x5c, 0xe0, 0x80, 0xe4, 0xb6,
	0x6f, 0x99, 0xa5, 0x2e, 0xce, 0x62, 0x37, 0x62, 0x42, 0x5c, 0xb8, 0x23, 0x21, 0xc1, 0x07, 0xe0,
	0xc8, 0x17, 0xe0, 0x3b, 0xf4, 0x38, 0x89, 0x0b, 0x27, 0x04, 0x2d, 0x1f, 0x04, 0xc5, 0x71, 0xc2,
	0x82, 0xa2, 0xb5, 0x37, 0xc7, 0x79, 0xff, 0xff, 0xff, 0xf7, 0x9e, 0x6d, 0x7a, 0xcb, 0x60, 0x6c,
	0x55, 0x8c, 0x13, 0x18, 0x63, 0xa2, 0x8d, 0xb2, 0x90, 0x85, 0x70, 0x3c, 0xc5, 0x54, 0x61, 0x2a,
	0x92, 0x54, 0x5b, 0xcd, 0x36, 0xca, 0x12, 0xe1, 0x4b, 0x44, 0x16, 0xb6, 0xf7, 0x46, 0xda, 0x1c,
	0x69, 0x03, 0x43, 0x69, 0xd0, 0xd5, 0x9f, 0x40, 0x16, 0x0e, 0xd1, 0xca, 0x10, 0x12, 0x19, 0xa9,
	0x58, 0x5a, 0xa5, 0xe3, 0xc2, 0xa0, 0xbd, 0x19, 0xe9, 0x48, 0xbb, 0x25, 0xe4, 0x2b, 0xbf, 0xbb,
	0x1d, 0x69, 0x1d, 0x4d, 0x10, 0x64, 0xa2, 0x40, 0xc6, 0xb1, 0xb6, 0x4e, 0x62, 0xfc, 0xdf, 0x46,
	0xae, 0x32, 0xdf, 0x95, 0x74, 0x5e, 0xd3, 0xcd, 0x67, 0x79, 0xf0, 0xa3, 0x62, 0xd7, 0x0c, 0xf0,
	0x78, 0x8a, 0xc6, 0xb2, 0xc7, 0x94, 0xfe, 0x43, 0xd8, 0x22, 0x37, 0x49, 0x77, 0xbd, 0x7f, 0x47,
	0x14, 0xbc, 0x22, 0xe7, 0x15, 0x8e, 0x57, 0x78, 0x5e, 0xf1, 0x54, 0x46, 0xe8, 0xb5, 0x83, 0x33,
	0xca, 0x0e, 0xd0, 0x8d, 0xb3, 0xfe, 0xa5, 0xfd, 0x16, 0x5d, 0x93, 0xe3, 0x71, 0x8a, 0xc6, 0x38,
	0xef, 0x2b, 0x83, 0xf2, 0xb3, 0xf3, 0x85, 0xd0, 0xeb, 0xff, 0x11, 0x99, 0x44, 0xc7, 0x06, 0xd9,
	0x43, 0x7a, 0xd9, 0xb3, 0xe7, 0xa2, 0x0b, 0xdd, 0xf5, 0xfe, 0xb6, 0x68, 0x98, 0xaa, 0xf0, 0xc2,
	0xfd, 0x8b, 0xb3, 0x9f, 0x3b, 0xc1, 0xa0, 0xd2, 0xb0, 0x27, 0xb5, 0x96, 0x5a, 0xae, 0xa5, 0xdd,
	0xa5, 0x2d, 0x15, 0xe1, 0xb5, 0x9e, 0x5e, 0xd4, 0x67, 0x56, 0x01, 0x3e, 0xa0, 0x6b, 0x3e, 0xcc,
	0x0f, 0x6c, 0x15, 0xbe, 0x52, 0xd2, 0xff, 0xd6, 0xa2, 0x57, 0x9d, 0xed, 0x73, 0x4c, 0x33, 0x35,
	0x42, 0xf6, 0x81, 0xd0, 0x6b, 0xb5, 0x49, 0xb0, 0x5e, 0xa3, 0x5f, 0xd3, 0xf9, 0xb5, 0xf7, 0x56,
	0x29, 0x2d, 0xb8, 0x3b, 0xb7, 0xdf, 0x7f, 0xff, 0xf3, 0xa9, 0xb5, 0xc3, 0x6e, 0xc0, 0x39, 0xf7,
	0xc5, 0xb0, 0xcf, 0xc4, 0x03, 0x7a, 0x03, 0xd6, 0x5d, 0x9a, 0x51, 0xd2, 0xf4, 0x56, 0xa8, 0xf4,
	0x30, 0xe0, 0x60, 0x7a, 0x6c, 0xf7, 0x5c, 0x18, 0x78, 0xeb, 0xef, 0xcb, 0xbb, 0xfd, 0x97, 0xb3,
	0xdf, 0x3c, 0xf8, 0x3a, 0xe7, 0xc1, 0x6c, 0xce, 0xc9, 0xe9, 0x9c, 0x93, 0x5f, 0x73, 0x4e, 0x3e,
	0x2e, 0x78, 0x70, 0xba, 0xe0, 0xc1, 0x8f, 0x05, 0x0f, 0x5e, 0xdd, 0x8f, 0x94, 0x3d, 0x9c, 0x0e,
	0xc5, 0x48, 0x1f, 0x55, 0xa6, 0x77, 0xf5, 0xc1, 0x81, 0x1a, 0x29, 0x39, 0x81, 0xc3, 0xe9, 0x10,
	0xb2, 0xb0, 0x0f, 0x6f, 0xaa, 0x1c, 0x7b, 0x92, 0xa0, 0xc9, 0x5f, 0xe2, 0x25, 0xf7, 0x46, 0xee,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x77, 0x27, 0xba, 0xe0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	QueryDeposits(ctx context.Context, in *QueryDepositsRequest, opts ...grpc.CallOption) (*QueryDepositsResponse, error)
	QueryDeposit(ctx context.Context, in *QueryDepositRequest, opts ...grpc.CallOption) (*QueryDepositResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryDeposits(ctx context.Context, in *QueryDepositsRequest, opts ...grpc.CallOption) (*QueryDepositsResponse, error) {
	out := new(QueryDepositsResponse)
	err := c.cc.Invoke(ctx, "/sentinel.deposit.v1.QueryService/QueryDeposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryDeposit(ctx context.Context, in *QueryDepositRequest, opts ...grpc.CallOption) (*QueryDepositResponse, error) {
	out := new(QueryDepositResponse)
	err := c.cc.Invoke(ctx, "/sentinel.deposit.v1.QueryService/QueryDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	QueryDeposits(context.Context, *QueryDepositsRequest) (*QueryDepositsResponse, error)
	QueryDeposit(context.Context, *QueryDepositRequest) (*QueryDepositResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryDeposits(ctx context.Context, req *QueryDepositsRequest) (*QueryDepositsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeposits not implemented")
}
func (*UnimplementedQueryServiceServer) QueryDeposit(ctx context.Context, req *QueryDepositRequest) (*QueryDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeposit not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryDeposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryDeposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.deposit.v1.QueryService/QueryDeposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryDeposits(ctx, req.(*QueryDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.deposit.v1.QueryService/QueryDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryDeposit(ctx, req.(*QueryDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var QueryService_serviceDesc = _QueryService_serviceDesc
var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.deposit.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryDeposits",
			Handler:    _QueryService_QueryDeposits_Handler,
		},
		{
			MethodName: "QueryDeposit",
			Handler:    _QueryService_QueryDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/deposit/v1/querier.proto",
}

func (m *QueryDepositsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDepositsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryDepositRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryDepositsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovQuerier(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryDepositResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deposit.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDepositsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryDepositsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryDepositRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryDepositRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryDepositsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryDepositsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryDepositResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryDepositResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)
