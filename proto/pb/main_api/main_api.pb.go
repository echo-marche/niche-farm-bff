// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.2
// source: main_api.proto

package main_api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemCode string `protobuf:"bytes,1,opt,name=system_code,json=systemCode,proto3" json:"system_code,omitempty"`
}

func (x *SampleRequest) Reset() {
	*x = SampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequest) ProtoMessage() {}

func (x *SampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequest.ProtoReflect.Descriptor instead.
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return file_main_api_proto_rawDescGZIP(), []int{0}
}

func (x *SampleRequest) GetSystemCode() string {
	if x != nil {
		return x.SystemCode
	}
	return ""
}

type SampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SampleResponse) Reset() {
	*x = SampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResponse) ProtoMessage() {}

func (x *SampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResponse.ProtoReflect.Descriptor instead.
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return file_main_api_proto_rawDescGZIP(), []int{1}
}

func (x *SampleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SampleResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_main_api_proto protoreflect.FileDescriptor

var file_main_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x0a, 0x05, 0x5b, 0x41,
	0x42, 0x5a, 0x5d, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x3a, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x5a, 0x0a, 0x07, 0x4d,
	0x61, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x12, 0x4f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07,
	0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_api_proto_rawDescOnce sync.Once
	file_main_api_proto_rawDescData = file_main_api_proto_rawDesc
)

func file_main_api_proto_rawDescGZIP() []byte {
	file_main_api_proto_rawDescOnce.Do(func() {
		file_main_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_api_proto_rawDescData)
	})
	return file_main_api_proto_rawDescData
}

var file_main_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_main_api_proto_goTypes = []interface{}{
	(*SampleRequest)(nil),  // 0: main_api.SampleRequest
	(*SampleResponse)(nil), // 1: main_api.SampleResponse
}
var file_main_api_proto_depIdxs = []int32{
	0, // 0: main_api.MainApi.GetSample:input_type -> main_api.SampleRequest
	1, // 1: main_api.MainApi.GetSample:output_type -> main_api.SampleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_api_proto_init() }
func file_main_api_proto_init() {
	if File_main_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleRequest); i {
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
		file_main_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleResponse); i {
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
			RawDescriptor: file_main_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_api_proto_goTypes,
		DependencyIndexes: file_main_api_proto_depIdxs,
		MessageInfos:      file_main_api_proto_msgTypes,
	}.Build()
	File_main_api_proto = out.File
	file_main_api_proto_rawDesc = nil
	file_main_api_proto_goTypes = nil
	file_main_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MainApiClient is the client API for MainApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MainApiClient interface {
	GetSample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type mainApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMainApiClient(cc grpc.ClientConnInterface) MainApiClient {
	return &mainApiClient{cc}
}

func (c *mainApiClient) GetSample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/main_api.MainApi/GetSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainApiServer is the server API for MainApi service.
type MainApiServer interface {
	GetSample(context.Context, *SampleRequest) (*SampleResponse, error)
}

// UnimplementedMainApiServer can be embedded to have forward compatible implementations.
type UnimplementedMainApiServer struct {
}

func (*UnimplementedMainApiServer) GetSample(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSample not implemented")
}

func RegisterMainApiServer(s *grpc.Server, srv MainApiServer) {
	s.RegisterService(&_MainApi_serviceDesc, srv)
}

func _MainApi_GetSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainApiServer).GetSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main_api.MainApi/GetSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainApiServer).GetSample(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MainApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main_api.MainApi",
	HandlerType: (*MainApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSample",
			Handler:    _MainApi_GetSample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main_api.proto",
}
