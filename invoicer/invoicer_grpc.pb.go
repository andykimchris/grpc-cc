// Definition of our API. Add a service with RPC methods

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: invoicer.proto

package invoicer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Invoicer_Create_FullMethodName            = "/Invoicer/Create"
	Invoicer_SumNums_FullMethodName           = "/Invoicer/SumNums"
	Invoicer_ExchangeConverter_FullMethodName = "/Invoicer/ExchangeConverter"
	Invoicer_UploadInvoices_FullMethodName    = "/Invoicer/UploadInvoices"
	Invoicer_ListInvoices_FullMethodName      = "/Invoicer/ListInvoices"
	Invoicer_ChatWithClient_FullMethodName    = "/Invoicer/ChatWithClient"
)

// InvoicerClient is the client API for Invoicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoicerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	SumNums(ctx context.Context, in *MultipleAmounts, opts ...grpc.CallOption) (*SumsResponse, error)
	ExchangeConverter(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
	// client side streaming
	UploadInvoices(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[InvoiceRequest, UploadSummaryResponse], error)
	// server side streaming
	ListInvoices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InvoiceRequest], error)
	// bidirectional streaming
	ChatWithClient(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error)
}

type invoicerClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicerClient(cc grpc.ClientConnInterface) InvoicerClient {
	return &invoicerClient{cc}
}

func (c *invoicerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Invoicer_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicerClient) SumNums(ctx context.Context, in *MultipleAmounts, opts ...grpc.CallOption) (*SumsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumsResponse)
	err := c.cc.Invoke(ctx, Invoicer_SumNums_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicerClient) ExchangeConverter(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, Invoicer_ExchangeConverter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicerClient) UploadInvoices(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[InvoiceRequest, UploadSummaryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Invoicer_ServiceDesc.Streams[0], Invoicer_UploadInvoices_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[InvoiceRequest, UploadSummaryResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Invoicer_UploadInvoicesClient = grpc.ClientStreamingClient[InvoiceRequest, UploadSummaryResponse]

func (c *invoicerClient) ListInvoices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InvoiceRequest], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Invoicer_ServiceDesc.Streams[1], Invoicer_ListInvoices_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, InvoiceRequest]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Invoicer_ListInvoicesClient = grpc.ServerStreamingClient[InvoiceRequest]

func (c *invoicerClient) ChatWithClient(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Invoicer_ServiceDesc.Streams[2], Invoicer_ChatWithClient_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ChatMessage, ChatMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Invoicer_ChatWithClientClient = grpc.BidiStreamingClient[ChatMessage, ChatMessage]

// InvoicerServer is the server API for Invoicer service.
// All implementations must embed UnimplementedInvoicerServer
// for forward compatibility.
type InvoicerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	SumNums(context.Context, *MultipleAmounts) (*SumsResponse, error)
	ExchangeConverter(context.Context, *ExchangeRequest) (*ExchangeResponse, error)
	// client side streaming
	UploadInvoices(grpc.ClientStreamingServer[InvoiceRequest, UploadSummaryResponse]) error
	// server side streaming
	ListInvoices(*Empty, grpc.ServerStreamingServer[InvoiceRequest]) error
	// bidirectional streaming
	ChatWithClient(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error
	mustEmbedUnimplementedInvoicerServer()
}

// UnimplementedInvoicerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInvoicerServer struct{}

func (UnimplementedInvoicerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInvoicerServer) SumNums(context.Context, *MultipleAmounts) (*SumsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumNums not implemented")
}
func (UnimplementedInvoicerServer) ExchangeConverter(context.Context, *ExchangeRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeConverter not implemented")
}
func (UnimplementedInvoicerServer) UploadInvoices(grpc.ClientStreamingServer[InvoiceRequest, UploadSummaryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadInvoices not implemented")
}
func (UnimplementedInvoicerServer) ListInvoices(*Empty, grpc.ServerStreamingServer[InvoiceRequest]) error {
	return status.Errorf(codes.Unimplemented, "method ListInvoices not implemented")
}
func (UnimplementedInvoicerServer) ChatWithClient(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method ChatWithClient not implemented")
}
func (UnimplementedInvoicerServer) mustEmbedUnimplementedInvoicerServer() {}
func (UnimplementedInvoicerServer) testEmbeddedByValue()                  {}

// UnsafeInvoicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoicerServer will
// result in compilation errors.
type UnsafeInvoicerServer interface {
	mustEmbedUnimplementedInvoicerServer()
}

func RegisterInvoicerServer(s grpc.ServiceRegistrar, srv InvoicerServer) {
	// If the following call pancis, it indicates UnimplementedInvoicerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Invoicer_ServiceDesc, srv)
}

func _Invoicer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invoicer_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoicer_SumNums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultipleAmounts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).SumNums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invoicer_SumNums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).SumNums(ctx, req.(*MultipleAmounts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoicer_ExchangeConverter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).ExchangeConverter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invoicer_ExchangeConverter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).ExchangeConverter(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoicer_UploadInvoices_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InvoicerServer).UploadInvoices(&grpc.GenericServerStream[InvoiceRequest, UploadSummaryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Invoicer_UploadInvoicesServer = grpc.ClientStreamingServer[InvoiceRequest, UploadSummaryResponse]

func _Invoicer_ListInvoices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InvoicerServer).ListInvoices(m, &grpc.GenericServerStream[Empty, InvoiceRequest]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Invoicer_ListInvoicesServer = grpc.ServerStreamingServer[InvoiceRequest]

func _Invoicer_ChatWithClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InvoicerServer).ChatWithClient(&grpc.GenericServerStream[ChatMessage, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Invoicer_ChatWithClientServer = grpc.BidiStreamingServer[ChatMessage, ChatMessage]

// Invoicer_ServiceDesc is the grpc.ServiceDesc for Invoicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Invoicer",
	HandlerType: (*InvoicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Invoicer_Create_Handler,
		},
		{
			MethodName: "SumNums",
			Handler:    _Invoicer_SumNums_Handler,
		},
		{
			MethodName: "ExchangeConverter",
			Handler:    _Invoicer_ExchangeConverter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadInvoices",
			Handler:       _Invoicer_UploadInvoices_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListInvoices",
			Handler:       _Invoicer_ListInvoices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChatWithClient",
			Handler:       _Invoicer_ChatWithClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "invoicer.proto",
}
