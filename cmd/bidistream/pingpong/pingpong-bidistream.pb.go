// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: cmd/bidistream/pingpong/pingpong-bidistream.proto

// this will be package of the generated code

package pingpong

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

// Message structure
type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field number 1-15 use 1 byte, while field 16th - 2047th use 2 bytes
	// So, the first 15 fields should be reserved for fields that are used oftenly
	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ping) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pong) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cmd_bidistream_pingpong_pingpong_bidistream_proto protoreflect.FileDescriptor

var file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6d, 0x64, 0x2f, 0x62, 0x69, 0x64, 0x69, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f,
	0x6e, 0x67, 0x2d, 0x62, 0x69, 0x64, 0x69, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x30, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x30, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x45, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x39, 0x0a,
	0x13, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e,
	0x50, 0x6f, 0x6e, 0x67, 0x28, 0x01, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x6e, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x31, 0x30, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescOnce sync.Once
	file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescData = file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDesc
)

func file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescGZIP() []byte {
	file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescOnce.Do(func() {
		file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescData)
	})
	return file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDescData
}

var file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_bidistream_pingpong_pingpong_bidistream_proto_goTypes = []interface{}{
	(*Ping)(nil), // 0: pingpong.Ping
	(*Pong)(nil), // 1: pingpong.Pong
}
var file_cmd_bidistream_pingpong_pingpong_bidistream_proto_depIdxs = []int32{
	0, // 0: pingpong.PingPong.BidirectionalStream:input_type -> pingpong.Ping
	1, // 1: pingpong.PingPong.BidirectionalStream:output_type -> pingpong.Pong
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_bidistream_pingpong_pingpong_bidistream_proto_init() }
func file_cmd_bidistream_pingpong_pingpong_bidistream_proto_init() {
	if File_cmd_bidistream_pingpong_pingpong_bidistream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_bidistream_pingpong_pingpong_bidistream_proto_goTypes,
		DependencyIndexes: file_cmd_bidistream_pingpong_pingpong_bidistream_proto_depIdxs,
		MessageInfos:      file_cmd_bidistream_pingpong_pingpong_bidistream_proto_msgTypes,
	}.Build()
	File_cmd_bidistream_pingpong_pingpong_bidistream_proto = out.File
	file_cmd_bidistream_pingpong_pingpong_bidistream_proto_rawDesc = nil
	file_cmd_bidistream_pingpong_pingpong_bidistream_proto_goTypes = nil
	file_cmd_bidistream_pingpong_pingpong_bidistream_proto_depIdxs = nil
}
