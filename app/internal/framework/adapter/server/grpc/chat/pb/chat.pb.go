// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: chat.proto

package pb

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

// 채팅 메시지 요청
type MsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request int32  `protobuf:"varint,1,opt,name=request,proto3" json:"request,omitempty"`
	RoomIdx int32  `protobuf:"varint,2,opt,name=room_idx,json=roomIdx,proto3" json:"room_idx,omitempty"`
	UserIdx int32  `protobuf:"varint,3,opt,name=user_idx,json=userIdx,proto3" json:"user_idx,omitempty"`
	Body    string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MsgReq) Reset() {
	*x = MsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReq) ProtoMessage() {}

func (x *MsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReq.ProtoReflect.Descriptor instead.
func (*MsgReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MsgReq) GetRequest() int32 {
	if x != nil {
		return x.Request
	}
	return 0
}

func (x *MsgReq) GetRoomIdx() int32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *MsgReq) GetUserIdx() int32 {
	if x != nil {
		return x.UserIdx
	}
	return 0
}

func (x *MsgReq) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// 채팅 메시지 응답
type MsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int32  `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	RoomIdx  int32  `protobuf:"varint,2,opt,name=room_idx,json=roomIdx,proto3" json:"room_idx,omitempty"`
	UserIdx  int32  `protobuf:"varint,3,opt,name=user_idx,json=userIdx,proto3" json:"user_idx,omitempty"`
	Body     string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MsgRes) Reset() {
	*x = MsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRes) ProtoMessage() {}

func (x *MsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRes.ProtoReflect.Descriptor instead.
func (*MsgRes) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MsgRes) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *MsgRes) GetRoomIdx() int32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *MsgRes) GetUserIdx() int32 {
	if x != nil {
		return x.UserIdx
	}
	return 0
}

func (x *MsgRes) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x6c, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x78, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x6e, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x78, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x32, 0x3c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0c,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(*MsgReq)(nil), // 0: chat.MsgReq
	(*MsgRes)(nil), // 1: chat.MsgRes
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.ChatService:input_type -> chat.MsgReq
	1, // 1: chat.ChatService.ChatService:output_type -> chat.MsgRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReq); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRes); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	ChatService(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatServiceClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ChatService(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat.ChatService/ChatService", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatServiceClient{stream}
	return x, nil
}

type ChatService_ChatServiceClient interface {
	Send(*MsgReq) error
	Recv() (*MsgRes, error)
	grpc.ClientStream
}

type chatServiceChatServiceClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatServiceClient) Send(m *MsgReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatServiceClient) Recv() (*MsgRes, error) {
	m := new(MsgRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	ChatService(ChatService_ChatServiceServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) ChatService(ChatService_ChatServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatService not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_ChatService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ChatService(&chatServiceChatServiceServer{stream})
}

type ChatService_ChatServiceServer interface {
	Send(*MsgRes) error
	Recv() (*MsgReq, error)
	grpc.ServerStream
}

type chatServiceChatServiceServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatServiceServer) Send(m *MsgRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatServiceServer) Recv() (*MsgReq, error) {
	m := new(MsgReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatService",
			Handler:       _ChatService_ChatService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}