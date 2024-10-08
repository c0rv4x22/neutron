// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/transfer/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("neutron/transfer/v1/query.proto", fileDescriptor_560cfedb574fdf6b) }

var fileDescriptor_560cfedb574fdf6b = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x4b, 0x03, 0x31,
	0x18, 0xc6, 0x7b, 0x42, 0x0b, 0xc6, 0x45, 0x32, 0x96, 0x12, 0xa1, 0x74, 0xd0, 0xaa, 0x89, 0xd7,
	0x56, 0x3f, 0x80, 0x38, 0x38, 0x38, 0xa8, 0x38, 0xb9, 0x48, 0xee, 0x8c, 0x77, 0x07, 0x6d, 0x92,
	0x26, 0xb9, 0x62, 0x91, 0x2e, 0x7e, 0x02, 0xa1, 0x5f, 0xc1, 0x41, 0xc4, 0x0f, 0xe2, 0x58, 0x70,
	0x71, 0x94, 0xd6, 0x0f, 0x22, 0x97, 0xbb, 0xb3, 0x0a, 0xfe, 0xe9, 0x6d, 0x49, 0x78, 0x9e, 0xe7,
	0xfd, 0xe5, 0x7d, 0x13, 0xb0, 0xc6, 0x59, 0x6c, 0x94, 0xe0, 0xc4, 0x28, 0xca, 0xf5, 0x15, 0x53,
	0x64, 0xe0, 0x92, 0x7e, 0xcc, 0xd4, 0x10, 0x4b, 0x25, 0x8c, 0x80, 0xab, 0x99, 0x00, 0xe7, 0x82,
	0x6a, 0x2d, 0x10, 0x22, 0xe8, 0x32, 0x42, 0x65, 0x44, 0x28, 0xe7, 0xc2, 0x50, 0x13, 0x09, 0xae,
	0x53, 0x7d, 0x75, 0x3d, 0xf2, 0x7c, 0x42, 0xa5, 0xec, 0x46, 0x7e, 0x7a, 0xfe, 0x5b, 0x72, 0xeb,
	0xbe, 0x0c, 0xca, 0x27, 0xc9, 0x1e, 0x3e, 0x39, 0x00, 0x1c, 0x30, 0x2e, 0x7a, 0x67, 0x8a, 0xfa,
	0x0c, 0x76, 0x70, 0xe4, 0xf9, 0xf8, 0x6b, 0xc6, 0x67, 0x71, 0x3c, 0x70, 0xb1, 0xf5, 0xcc, 0xe5,
	0xa7, 0xac, 0x1f, 0x33, 0x6d, 0xaa, 0xbb, 0x05, 0x5d, 0x5a, 0x0a, 0xae, 0x59, 0xdd, 0xbd, 0x7d,
	0x79, 0x1f, 0x2f, 0x6d, 0xc2, 0x0d, 0x92, 0x81, 0x7f, 0x07, 0xbe, 0x4c, 0x1c, 0x17, 0x26, 0xb1,
	0x68, 0x72, 0x13, 0x52, 0x1d, 0x8e, 0xe0, 0x83, 0x03, 0x56, 0xe6, 0x49, 0x1a, 0x16, 0xab, 0xac,
	0x73, 0xe0, 0xbd, 0xa2, 0xb6, 0x8c, 0xb8, 0x69, 0x89, 0x1b, 0xb0, 0xfe, 0x3f, 0x31, 0x1c, 0x3b,
	0xa0, 0x72, 0x4c, 0x15, 0xed, 0x69, 0xb8, 0xb3, 0x40, 0xb9, 0x54, 0x9a, 0x03, 0xba, 0x05, 0x1c,
	0x19, 0x5b, 0xc3, 0xb2, 0x21, 0x58, 0xfb, 0x99, 0x4d, 0xa6, 0x28, 0x8f, 0x0e, 0x58, 0xb6, 0x37,
	0x3b, 0xa4, 0x3a, 0x84, 0xed, 0x45, 0xfb, 0x90, 0xa8, 0x73, 0xb6, 0x4e, 0x31, 0x53, 0x86, 0xd7,
	0xb2, 0x78, 0x5b, 0xb0, 0xf9, 0x57, 0xeb, 0x92, 0x21, 0x27, 0xc3, 0xb6, 0x2d, 0x1c, 0xed, 0x1f,
	0x3d, 0x4f, 0x91, 0x33, 0x99, 0x22, 0xe7, 0x6d, 0x8a, 0x9c, 0xbb, 0x19, 0x2a, 0x4d, 0x66, 0xa8,
	0xf4, 0x3a, 0x43, 0xa5, 0xf3, 0x56, 0x10, 0x99, 0x30, 0xf6, 0xb0, 0x2f, 0x7a, 0x24, 0xfb, 0x25,
	0xdb, 0x42, 0x05, 0xf9, 0x9a, 0x0c, 0x3a, 0xe4, 0x7a, 0x9e, 0x6f, 0x86, 0x92, 0x69, 0xaf, 0x62,
	0xdf, 0x7e, 0xfb, 0x23, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x46, 0x5b, 0xdc, 0x78, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// DenomTrace queries a denomination trace information.
	DenomTrace(ctx context.Context, in *types.QueryDenomTraceRequest, opts ...grpc.CallOption) (*types.QueryDenomTraceResponse, error)
	// DenomTraces queries all denomination traces.
	DenomTraces(ctx context.Context, in *types.QueryDenomTracesRequest, opts ...grpc.CallOption) (*types.QueryDenomTracesResponse, error)
	// Params queries all parameters of the ibc-transfer module.
	Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error)
	// DenomHash queries a denomination hash information.
	DenomHash(ctx context.Context, in *types.QueryDenomHashRequest, opts ...grpc.CallOption) (*types.QueryDenomHashResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DenomTrace(ctx context.Context, in *types.QueryDenomTraceRequest, opts ...grpc.CallOption) (*types.QueryDenomTraceResponse, error) {
	out := new(types.QueryDenomTraceResponse)
	err := c.cc.Invoke(ctx, "/neutron.transfer.Query/DenomTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomTraces(ctx context.Context, in *types.QueryDenomTracesRequest, opts ...grpc.CallOption) (*types.QueryDenomTracesResponse, error) {
	out := new(types.QueryDenomTracesResponse)
	err := c.cc.Invoke(ctx, "/neutron.transfer.Query/DenomTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	out := new(types.QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/neutron.transfer.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomHash(ctx context.Context, in *types.QueryDenomHashRequest, opts ...grpc.CallOption) (*types.QueryDenomHashResponse, error) {
	out := new(types.QueryDenomHashResponse)
	err := c.cc.Invoke(ctx, "/neutron.transfer.Query/DenomHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DenomTrace queries a denomination trace information.
	DenomTrace(context.Context, *types.QueryDenomTraceRequest) (*types.QueryDenomTraceResponse, error)
	// DenomTraces queries all denomination traces.
	DenomTraces(context.Context, *types.QueryDenomTracesRequest) (*types.QueryDenomTracesResponse, error)
	// Params queries all parameters of the ibc-transfer module.
	Params(context.Context, *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
	// DenomHash queries a denomination hash information.
	DenomHash(context.Context, *types.QueryDenomHashRequest) (*types.QueryDenomHashResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DenomTrace(ctx context.Context, req *types.QueryDenomTraceRequest) (*types.QueryDenomTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTrace not implemented")
}
func (*UnimplementedQueryServer) DenomTraces(ctx context.Context, req *types.QueryDenomTracesRequest) (*types.QueryDenomTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTraces not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DenomHash(ctx context.Context, req *types.QueryDenomHashRequest) (*types.QueryDenomHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomHash not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.transfer.Query/DenomTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTrace(ctx, req.(*types.QueryDenomTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.transfer.Query/DenomTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTraces(ctx, req.(*types.QueryDenomTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.transfer.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*types.QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.transfer.Query/DenomHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomHash(ctx, req.(*types.QueryDenomHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neutron.transfer.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomTrace",
			Handler:    _Query_DenomTrace_Handler,
		},
		{
			MethodName: "DenomTraces",
			Handler:    _Query_DenomTraces_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomHash",
			Handler:    _Query_DenomHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neutron/transfer/v1/query.proto",
}
