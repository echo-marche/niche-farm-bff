// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.2
// source: sendmail.proto

package sendmail

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string `protobuf:"bytes,1,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress   string `protobuf:"bytes,2,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sendmail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_sendmail_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *EmailRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

type EmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sendmail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_sendmail_proto_rawDescGZIP(), []int{1}
}

func (x *EmailResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_sendmail_proto protoreflect.FileDescriptor

var file_sendmail_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67,
	0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x27, 0x0a, 0x0d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x55, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x49, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x61, 0x69, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sendmail_proto_rawDescOnce sync.Once
	file_sendmail_proto_rawDescData = file_sendmail_proto_rawDesc
)

func file_sendmail_proto_rawDescGZIP() []byte {
	file_sendmail_proto_rawDescOnce.Do(func() {
		file_sendmail_proto_rawDescData = protoimpl.X.CompressGZIP(file_sendmail_proto_rawDescData)
	})
	return file_sendmail_proto_rawDescData
}

var file_sendmail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sendmail_proto_goTypes = []interface{}{
	(*EmailRequest)(nil),  // 0: sendmail.EmailRequest
	(*EmailResponse)(nil), // 1: sendmail.EmailResponse
}
var file_sendmail_proto_depIdxs = []int32{
	0, // 0: sendmail.Sendmail.SendUserRegistration:input_type -> sendmail.EmailRequest
	1, // 1: sendmail.Sendmail.SendUserRegistration:output_type -> sendmail.EmailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sendmail_proto_init() }
func file_sendmail_proto_init() {
	if File_sendmail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sendmail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRequest); i {
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
		file_sendmail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailResponse); i {
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
			RawDescriptor: file_sendmail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sendmail_proto_goTypes,
		DependencyIndexes: file_sendmail_proto_depIdxs,
		MessageInfos:      file_sendmail_proto_msgTypes,
	}.Build()
	File_sendmail_proto = out.File
	file_sendmail_proto_rawDesc = nil
	file_sendmail_proto_goTypes = nil
	file_sendmail_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SendmailClient is the client API for Sendmail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendmailClient interface {
	SendUserRegistration(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
}

type sendmailClient struct {
	cc grpc.ClientConnInterface
}

func NewSendmailClient(cc grpc.ClientConnInterface) SendmailClient {
	return &sendmailClient{cc}
}

func (c *sendmailClient) SendUserRegistration(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/sendmail.Sendmail/SendUserRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendmailServer is the server API for Sendmail service.
type SendmailServer interface {
	SendUserRegistration(context.Context, *EmailRequest) (*EmailResponse, error)
}

// UnimplementedSendmailServer can be embedded to have forward compatible implementations.
type UnimplementedSendmailServer struct {
}

func (*UnimplementedSendmailServer) SendUserRegistration(context.Context, *EmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserRegistration not implemented")
}

func RegisterSendmailServer(s *grpc.Server, srv SendmailServer) {
	s.RegisterService(&_Sendmail_serviceDesc, srv)
}

func _Sendmail_SendUserRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendmailServer).SendUserRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sendmail.Sendmail/SendUserRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendmailServer).SendUserRegistration(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sendmail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sendmail.Sendmail",
	HandlerType: (*SendmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUserRegistration",
			Handler:    _Sendmail_SendUserRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sendmail.proto",
}
