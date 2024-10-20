// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: pb/logger.proto

package pb

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

type LogEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level       string            `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Message     string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ServiceName string            `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Data        map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LogEventRequest) Reset() {
	*x = LogEventRequest{}
	mi := &file_pb_logger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEventRequest) ProtoMessage() {}

func (x *LogEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_logger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEventRequest.ProtoReflect.Descriptor instead.
func (*LogEventRequest) Descriptor() ([]byte, []int) {
	return file_pb_logger_proto_rawDescGZIP(), []int{0}
}

func (x *LogEventRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogEventRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEventRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *LogEventRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogEventResponse) Reset() {
	*x = LogEventResponse{}
	mi := &file_pb_logger_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEventResponse) ProtoMessage() {}

func (x *LogEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_logger_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEventResponse.ProtoReflect.Descriptor instead.
func (*LogEventResponse) Descriptor() ([]byte, []int) {
	return file_pb_logger_proto_rawDescGZIP(), []int{1}
}

func (x *LogEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LogEventResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_logger_proto protoreflect.FileDescriptor

var file_pb_logger_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x22, 0xd4, 0x01, 0x0a, 0x0f, 0x4c, 0x6f,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x46, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x50, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c,
	0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_logger_proto_rawDescOnce sync.Once
	file_pb_logger_proto_rawDescData = file_pb_logger_proto_rawDesc
)

func file_pb_logger_proto_rawDescGZIP() []byte {
	file_pb_logger_proto_rawDescOnce.Do(func() {
		file_pb_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_logger_proto_rawDescData)
	})
	return file_pb_logger_proto_rawDescData
}

var file_pb_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_logger_proto_goTypes = []any{
	(*LogEventRequest)(nil),  // 0: logger.LogEventRequest
	(*LogEventResponse)(nil), // 1: logger.LogEventResponse
	nil,                      // 2: logger.LogEventRequest.DataEntry
}
var file_pb_logger_proto_depIdxs = []int32{
	2, // 0: logger.LogEventRequest.data:type_name -> logger.LogEventRequest.DataEntry
	0, // 1: logger.LoggerService.LogEvent:input_type -> logger.LogEventRequest
	1, // 2: logger.LoggerService.LogEvent:output_type -> logger.LogEventResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_logger_proto_init() }
func file_pb_logger_proto_init() {
	if File_pb_logger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_logger_proto_goTypes,
		DependencyIndexes: file_pb_logger_proto_depIdxs,
		MessageInfos:      file_pb_logger_proto_msgTypes,
	}.Build()
	File_pb_logger_proto = out.File
	file_pb_logger_proto_rawDesc = nil
	file_pb_logger_proto_goTypes = nil
	file_pb_logger_proto_depIdxs = nil
}
