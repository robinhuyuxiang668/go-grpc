// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.9.0
// source: both_stream.proto

package both_stream_grpc

import (
	context "context"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_both_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_both_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_both_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_both_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_both_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_both_stream_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_both_stream_proto protoreflect.FileDescriptor

var file_both_stream_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x62, 0x6f, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x32, 0x51, 0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x68, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x43, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x74,
	0x68, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6f, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x62, 0x6f, 0x74, 0x68,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_both_stream_proto_rawDescOnce sync.Once
	file_both_stream_proto_rawDescData = file_both_stream_proto_rawDesc
)

func file_both_stream_proto_rawDescGZIP() []byte {
	file_both_stream_proto_rawDescOnce.Do(func() {
		file_both_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_both_stream_proto_rawDescData)
	})
	return file_both_stream_proto_rawDescData
}

var file_both_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_both_stream_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: both_stream_grpc.Request
	(*Response)(nil), // 1: both_stream_grpc.Response
}
var file_both_stream_proto_depIdxs = []int32{
	0, // 0: both_stream_grpc.BothStream.Chat:input_type -> both_stream_grpc.Request
	1, // 1: both_stream_grpc.BothStream.Chat:output_type -> both_stream_grpc.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_both_stream_proto_init() }
func file_both_stream_proto_init() {
	if File_both_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_both_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_both_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_both_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_both_stream_proto_goTypes,
		DependencyIndexes: file_both_stream_proto_depIdxs,
		MessageInfos:      file_both_stream_proto_msgTypes,
	}.Build()
	File_both_stream_proto = out.File
	file_both_stream_proto_rawDesc = nil
	file_both_stream_proto_goTypes = nil
	file_both_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BothStreamClient is the client API for BothStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BothStreamClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (BothStream_ChatClient, error)
}

type bothStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewBothStreamClient(cc grpc.ClientConnInterface) BothStreamClient {
	return &bothStreamClient{cc}
}

func (c *bothStreamClient) Chat(ctx context.Context, opts ...grpc.CallOption) (BothStream_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BothStream_serviceDesc.Streams[0], "/both_stream_grpc.BothStream/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &bothStreamChatClient{stream}
	return x, nil
}

type BothStream_ChatClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type bothStreamChatClient struct {
	grpc.ClientStream
}

func (x *bothStreamChatClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bothStreamChatClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BothStreamServer is the server API for BothStream service.
type BothStreamServer interface {
	Chat(BothStream_ChatServer) error
}

// UnimplementedBothStreamServer can be embedded to have forward compatible implementations.
type UnimplementedBothStreamServer struct {
}

func (*UnimplementedBothStreamServer) Chat(BothStream_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}

func RegisterBothStreamServer(s *grpc.Server, srv BothStreamServer) {
	s.RegisterService(&_BothStream_serviceDesc, srv)
}

func _BothStream_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BothStreamServer).Chat(&bothStreamChatServer{stream})
}

type BothStream_ChatServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type bothStreamChatServer struct {
	grpc.ServerStream
}

func (x *bothStreamChatServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bothStreamChatServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BothStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "both_stream_grpc.BothStream",
	HandlerType: (*BothStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _BothStream_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "both_stream.proto",
}
