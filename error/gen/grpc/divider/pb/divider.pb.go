// Code generated with goa v3.2.3, DO NOT EDIT.
//
// divider protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: divider.proto

package dividerpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IntegerDivideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Left operand
	A int32 `protobuf:"zigzag32,1,opt,name=a,proto3" json:"a,omitempty"`
	// Right operand
	B int32 `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *IntegerDivideRequest) Reset() {
	*x = IntegerDivideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegerDivideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegerDivideRequest) ProtoMessage() {}

func (x *IntegerDivideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_divider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegerDivideRequest.ProtoReflect.Descriptor instead.
func (*IntegerDivideRequest) Descriptor() ([]byte, []int) {
	return file_divider_proto_rawDescGZIP(), []int{0}
}

func (x *IntegerDivideRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *IntegerDivideRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type IntegerDivideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field int32 `protobuf:"zigzag32,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *IntegerDivideResponse) Reset() {
	*x = IntegerDivideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegerDivideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegerDivideResponse) ProtoMessage() {}

func (x *IntegerDivideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_divider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegerDivideResponse.ProtoReflect.Descriptor instead.
func (*IntegerDivideResponse) Descriptor() ([]byte, []int) {
	return file_divider_proto_rawDescGZIP(), []int{1}
}

func (x *IntegerDivideResponse) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

type DivideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Left operand
	A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	// Right operand
	B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *DivideRequest) Reset() {
	*x = DivideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideRequest) ProtoMessage() {}

func (x *DivideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_divider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideRequest.ProtoReflect.Descriptor instead.
func (*DivideRequest) Descriptor() ([]byte, []int) {
	return file_divider_proto_rawDescGZIP(), []int{2}
}

func (x *DivideRequest) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *DivideRequest) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

type DivideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field float64 `protobuf:"fixed64,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *DivideResponse) Reset() {
	*x = DivideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideResponse) ProtoMessage() {}

func (x *DivideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_divider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideResponse.ProtoReflect.Descriptor instead.
func (*DivideResponse) Descriptor() ([]byte, []int) {
	return file_divider_proto_rawDescGZIP(), []int{3}
}

func (x *DivideResponse) GetField() float64 {
	if x != nil {
		return x.Field
	}
	return 0
}

var File_divider_proto protoreflect.FileDescriptor

var file_divider_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x14, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x61, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x62, 0x22, 0x2d, 0x0a, 0x15,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x2b, 0x0a, 0x0d, 0x44,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x62, 0x22, 0x26, 0x0a, 0x0e, 0x44, 0x69, 0x76, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x94, 0x01, 0x0a, 0x07, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0d,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x1d, 0x2e,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x44,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x44, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x64, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_divider_proto_rawDescOnce sync.Once
	file_divider_proto_rawDescData = file_divider_proto_rawDesc
)

func file_divider_proto_rawDescGZIP() []byte {
	file_divider_proto_rawDescOnce.Do(func() {
		file_divider_proto_rawDescData = protoimpl.X.CompressGZIP(file_divider_proto_rawDescData)
	})
	return file_divider_proto_rawDescData
}

var file_divider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_divider_proto_goTypes = []interface{}{
	(*IntegerDivideRequest)(nil),  // 0: divider.IntegerDivideRequest
	(*IntegerDivideResponse)(nil), // 1: divider.IntegerDivideResponse
	(*DivideRequest)(nil),         // 2: divider.DivideRequest
	(*DivideResponse)(nil),        // 3: divider.DivideResponse
}
var file_divider_proto_depIdxs = []int32{
	0, // 0: divider.Divider.IntegerDivide:input_type -> divider.IntegerDivideRequest
	2, // 1: divider.Divider.Divide:input_type -> divider.DivideRequest
	1, // 2: divider.Divider.IntegerDivide:output_type -> divider.IntegerDivideResponse
	3, // 3: divider.Divider.Divide:output_type -> divider.DivideResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_divider_proto_init() }
func file_divider_proto_init() {
	if File_divider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_divider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegerDivideRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_divider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegerDivideResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_divider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_divider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_divider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_divider_proto_goTypes,
		DependencyIndexes: file_divider_proto_depIdxs,
		MessageInfos:      file_divider_proto_msgTypes,
	}.Build()
	File_divider_proto = out.File
	file_divider_proto_rawDesc = nil
	file_divider_proto_goTypes = nil
	file_divider_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DividerClient is the client API for Divider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DividerClient interface {
	// IntegerDivide implements integer_divide.
	IntegerDivide(ctx context.Context, in *IntegerDivideRequest, opts ...grpc.CallOption) (*IntegerDivideResponse, error)
	// Divide implements divide.
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error)
}

type dividerClient struct {
	cc grpc.ClientConnInterface
}

func NewDividerClient(cc grpc.ClientConnInterface) DividerClient {
	return &dividerClient{cc}
}

func (c *dividerClient) IntegerDivide(ctx context.Context, in *IntegerDivideRequest, opts ...grpc.CallOption) (*IntegerDivideResponse, error) {
	out := new(IntegerDivideResponse)
	err := c.cc.Invoke(ctx, "/divider.Divider/IntegerDivide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dividerClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error) {
	out := new(DivideResponse)
	err := c.cc.Invoke(ctx, "/divider.Divider/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DividerServer is the server API for Divider service.
type DividerServer interface {
	// IntegerDivide implements integer_divide.
	IntegerDivide(context.Context, *IntegerDivideRequest) (*IntegerDivideResponse, error)
	// Divide implements divide.
	Divide(context.Context, *DivideRequest) (*DivideResponse, error)
}

// UnimplementedDividerServer can be embedded to have forward compatible implementations.
type UnimplementedDividerServer struct {
}

func (*UnimplementedDividerServer) IntegerDivide(context.Context, *IntegerDivideRequest) (*IntegerDivideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntegerDivide not implemented")
}
func (*UnimplementedDividerServer) Divide(context.Context, *DivideRequest) (*DivideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}

func RegisterDividerServer(s *grpc.Server, srv DividerServer) {
	s.RegisterService(&_Divider_serviceDesc, srv)
}

func _Divider_IntegerDivide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegerDivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DividerServer).IntegerDivide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/divider.Divider/IntegerDivide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DividerServer).IntegerDivide(ctx, req.(*IntegerDivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Divider_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DividerServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/divider.Divider/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DividerServer).Divide(ctx, req.(*DivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Divider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "divider.Divider",
	HandlerType: (*DividerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IntegerDivide",
			Handler:    _Divider_IntegerDivide_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Divider_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "divider.proto",
}
