// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xde, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x3e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c,
	0x52, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x6d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*Currency)(nil),           // 0: pb.Currency
	(*DeleteCurrencyReq)(nil),  // 1: pb.DeleteCurrencyReq
	(*ListCurrenciesReq)(nil),  // 2: pb.ListCurrenciesReq
	(*RateCurrencyReq)(nil),    // 3: pb.RateCurrencyReq
	(*RatingSumStreamReq)(nil), // 4: pb.RatingSumStreamReq
	(*DeleteCurrencyRes)(nil),  // 5: pb.DeleteCurrencyRes
	(*CurrencyList)(nil),       // 6: pb.CurrencyList
	(*RatingSumRes)(nil),       // 7: pb.RatingSumRes
}
var file_service_proto_depIdxs = []int32{
	0, // 0: pb.CurrencyService.CreateCurrency:input_type -> pb.Currency
	0, // 1: pb.CurrencyService.UpdateCurrency:input_type -> pb.Currency
	1, // 2: pb.CurrencyService.DeleteCurrency:input_type -> pb.DeleteCurrencyReq
	2, // 3: pb.CurrencyService.ListCurrencies:input_type -> pb.ListCurrenciesReq
	3, // 4: pb.CurrencyService.RateCurrency:input_type -> pb.RateCurrencyReq
	4, // 5: pb.CurrencyService.RatingSumStream:input_type -> pb.RatingSumStreamReq
	0, // 6: pb.CurrencyService.CreateCurrency:output_type -> pb.Currency
	0, // 7: pb.CurrencyService.UpdateCurrency:output_type -> pb.Currency
	5, // 8: pb.CurrencyService.DeleteCurrency:output_type -> pb.DeleteCurrencyRes
	6, // 9: pb.CurrencyService.ListCurrencies:output_type -> pb.CurrencyList
	7, // 10: pb.CurrencyService.RateCurrency:output_type -> pb.RatingSumRes
	7, // 11: pb.CurrencyService.RatingSumStream:output_type -> pb.RatingSumRes
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_currency_proto_init()
	file_rating_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	CreateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error)
	UpdateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error)
	DeleteCurrency(ctx context.Context, in *DeleteCurrencyReq, opts ...grpc.CallOption) (*DeleteCurrencyRes, error)
	ListCurrencies(ctx context.Context, in *ListCurrenciesReq, opts ...grpc.CallOption) (*CurrencyList, error)
	RateCurrency(ctx context.Context, in *RateCurrencyReq, opts ...grpc.CallOption) (*RatingSumRes, error)
	RatingSumStream(ctx context.Context, in *RatingSumStreamReq, opts ...grpc.CallOption) (CurrencyService_RatingSumStreamClient, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) CreateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/pb.CurrencyService/CreateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) UpdateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/pb.CurrencyService/UpdateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) DeleteCurrency(ctx context.Context, in *DeleteCurrencyReq, opts ...grpc.CallOption) (*DeleteCurrencyRes, error) {
	out := new(DeleteCurrencyRes)
	err := c.cc.Invoke(ctx, "/pb.CurrencyService/DeleteCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) ListCurrencies(ctx context.Context, in *ListCurrenciesReq, opts ...grpc.CallOption) (*CurrencyList, error) {
	out := new(CurrencyList)
	err := c.cc.Invoke(ctx, "/pb.CurrencyService/ListCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) RateCurrency(ctx context.Context, in *RateCurrencyReq, opts ...grpc.CallOption) (*RatingSumRes, error) {
	out := new(RatingSumRes)
	err := c.cc.Invoke(ctx, "/pb.CurrencyService/RateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) RatingSumStream(ctx context.Context, in *RatingSumStreamReq, opts ...grpc.CallOption) (CurrencyService_RatingSumStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CurrencyService_serviceDesc.Streams[0], "/pb.CurrencyService/RatingSumStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyServiceRatingSumStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyService_RatingSumStreamClient interface {
	Recv() (*RatingSumRes, error)
	grpc.ClientStream
}

type currencyServiceRatingSumStreamClient struct {
	grpc.ClientStream
}

func (x *currencyServiceRatingSumStreamClient) Recv() (*RatingSumRes, error) {
	m := new(RatingSumRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
type CurrencyServiceServer interface {
	CreateCurrency(context.Context, *Currency) (*Currency, error)
	UpdateCurrency(context.Context, *Currency) (*Currency, error)
	DeleteCurrency(context.Context, *DeleteCurrencyReq) (*DeleteCurrencyRes, error)
	ListCurrencies(context.Context, *ListCurrenciesReq) (*CurrencyList, error)
	RateCurrency(context.Context, *RateCurrencyReq) (*RatingSumRes, error)
	RatingSumStream(*RatingSumStreamReq, CurrencyService_RatingSumStreamServer) error
}

// UnimplementedCurrencyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (*UnimplementedCurrencyServiceServer) CreateCurrency(context.Context, *Currency) (*Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) UpdateCurrency(context.Context, *Currency) (*Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) DeleteCurrency(context.Context, *DeleteCurrencyReq) (*DeleteCurrencyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) ListCurrencies(context.Context, *ListCurrenciesReq) (*CurrencyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrencies not implemented")
}
func (*UnimplementedCurrencyServiceServer) RateCurrency(context.Context, *RateCurrencyReq) (*RatingSumRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) RatingSumStream(*RatingSumStreamReq, CurrencyService_RatingSumStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RatingSumStream not implemented")
}

func RegisterCurrencyServiceServer(s *grpc.Server, srv CurrencyServiceServer) {
	s.RegisterService(&_CurrencyService_serviceDesc, srv)
}

func _CurrencyService_CreateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).CreateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrencyService/CreateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).CreateCurrency(ctx, req.(*Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_UpdateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).UpdateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrencyService/UpdateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).UpdateCurrency(ctx, req.(*Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_DeleteCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCurrencyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).DeleteCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrencyService/DeleteCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).DeleteCurrency(ctx, req.(*DeleteCurrencyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_ListCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrenciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).ListCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrencyService/ListCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).ListCurrencies(ctx, req.(*ListCurrenciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_RateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateCurrencyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).RateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrencyService/RateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).RateCurrency(ctx, req.(*RateCurrencyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_RatingSumStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RatingSumStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyServiceServer).RatingSumStream(m, &currencyServiceRatingSumStreamServer{stream})
}

type CurrencyService_RatingSumStreamServer interface {
	Send(*RatingSumRes) error
	grpc.ServerStream
}

type currencyServiceRatingSumStreamServer struct {
	grpc.ServerStream
}

func (x *currencyServiceRatingSumStreamServer) Send(m *RatingSumRes) error {
	return x.ServerStream.SendMsg(m)
}

var _CurrencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCurrency",
			Handler:    _CurrencyService_CreateCurrency_Handler,
		},
		{
			MethodName: "UpdateCurrency",
			Handler:    _CurrencyService_UpdateCurrency_Handler,
		},
		{
			MethodName: "DeleteCurrency",
			Handler:    _CurrencyService_DeleteCurrency_Handler,
		},
		{
			MethodName: "ListCurrencies",
			Handler:    _CurrencyService_ListCurrencies_Handler,
		},
		{
			MethodName: "RateCurrency",
			Handler:    _CurrencyService_RateCurrency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RatingSumStream",
			Handler:       _CurrencyService_RatingSumStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
