// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: botapi/command/v1/command_service.proto

package command

import (
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

// Telegram User's message
type IncomingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Telegram User's ID
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Incoming message text
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *IncomingMessage) Reset() {
	*x = IncomingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_botapi_command_v1_command_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingMessage) ProtoMessage() {}

func (x *IncomingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_botapi_command_v1_command_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingMessage.ProtoReflect.Descriptor instead.
func (*IncomingMessage) Descriptor() ([]byte, []int) {
	return file_botapi_command_v1_command_service_proto_rawDescGZIP(), []int{0}
}

func (x *IncomingMessage) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IncomingMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Telegram User's ID
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Reply message text from services
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ReplyMessage) Reset() {
	*x = ReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_botapi_command_v1_command_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMessage) ProtoMessage() {}

func (x *ReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_botapi_command_v1_command_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMessage.ProtoReflect.Descriptor instead.
func (*ReplyMessage) Descriptor() ([]byte, []int) {
	return file_botapi_command_v1_command_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyMessage) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReplyMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_botapi_command_v1_command_service_proto protoreflect.FileDescriptor

var file_botapi_command_v1_command_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x6f, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6f, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x3e, 0x0a, 0x0f,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3b, 0x0a, 0x0c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x63, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x62, 0x6f,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62,
	0x2d, 0x7a, 0x61, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x6f, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_botapi_command_v1_command_service_proto_rawDescOnce sync.Once
	file_botapi_command_v1_command_service_proto_rawDescData = file_botapi_command_v1_command_service_proto_rawDesc
)

func file_botapi_command_v1_command_service_proto_rawDescGZIP() []byte {
	file_botapi_command_v1_command_service_proto_rawDescOnce.Do(func() {
		file_botapi_command_v1_command_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_botapi_command_v1_command_service_proto_rawDescData)
	})
	return file_botapi_command_v1_command_service_proto_rawDescData
}

var file_botapi_command_v1_command_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_botapi_command_v1_command_service_proto_goTypes = []interface{}{
	(*IncomingMessage)(nil), // 0: botapi.command.v1.IncomingMessage
	(*ReplyMessage)(nil),    // 1: botapi.command.v1.ReplyMessage
}
var file_botapi_command_v1_command_service_proto_depIdxs = []int32{
	0, // 0: botapi.command.v1.CommandService.GetReply:input_type -> botapi.command.v1.IncomingMessage
	1, // 1: botapi.command.v1.CommandService.GetReply:output_type -> botapi.command.v1.ReplyMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_botapi_command_v1_command_service_proto_init() }
func file_botapi_command_v1_command_service_proto_init() {
	if File_botapi_command_v1_command_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_botapi_command_v1_command_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomingMessage); i {
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
		file_botapi_command_v1_command_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMessage); i {
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
			RawDescriptor: file_botapi_command_v1_command_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_botapi_command_v1_command_service_proto_goTypes,
		DependencyIndexes: file_botapi_command_v1_command_service_proto_depIdxs,
		MessageInfos:      file_botapi_command_v1_command_service_proto_msgTypes,
	}.Build()
	File_botapi_command_v1_command_service_proto = out.File
	file_botapi_command_v1_command_service_proto_rawDesc = nil
	file_botapi_command_v1_command_service_proto_goTypes = nil
	file_botapi_command_v1_command_service_proto_depIdxs = nil
}