// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: default.proto

package hedgeproto

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

type Payload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          map[string]string      `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payload) Reset() {
	*x = Payload{}
	mi := &file_default_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_default_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_default_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Payload) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_default_proto protoreflect.FileDescriptor

var file_default_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xab, 0x02, 0x0a, 0x05, 0x48, 0x65, 0x64, 0x67,
	0x65, 0x12, 0x36, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x64, 0x67,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x13,
	0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x09, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x13, 0x2e, 0x68, 0x65,
	0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x6f, 0x53, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x13, 0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x53, 0x6f, 0x53, 0x52, 0x65, 0x61, 0x64, 0x12, 0x13, 0x2e,
	0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x13, 0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x36, 0x0a,
	0x08, 0x53, 0x6f, 0x53, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x64, 0x67,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x13,
	0x2e, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x65, 0x6e,
	0x69, 0x67, 0x68, 0x74, 0x2f, 0x68, 0x65, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_default_proto_rawDescOnce sync.Once
	file_default_proto_rawDescData = file_default_proto_rawDesc
)

func file_default_proto_rawDescGZIP() []byte {
	file_default_proto_rawDescOnce.Do(func() {
		file_default_proto_rawDescData = protoimpl.X.CompressGZIP(file_default_proto_rawDescData)
	})
	return file_default_proto_rawDescData
}

var file_default_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_default_proto_goTypes = []any{
	(*Payload)(nil), // 0: hedgeproto.Payload
	nil,             // 1: hedgeproto.Payload.MetaEntry
}
var file_default_proto_depIdxs = []int32{
	1, // 0: hedgeproto.Payload.meta:type_name -> hedgeproto.Payload.MetaEntry
	0, // 1: hedgeproto.Hedge.Send:input_type -> hedgeproto.Payload
	0, // 2: hedgeproto.Hedge.Broadcast:input_type -> hedgeproto.Payload
	0, // 3: hedgeproto.Hedge.SoSWrite:input_type -> hedgeproto.Payload
	0, // 4: hedgeproto.Hedge.SoSRead:input_type -> hedgeproto.Payload
	0, // 5: hedgeproto.Hedge.SoSClose:input_type -> hedgeproto.Payload
	0, // 6: hedgeproto.Hedge.Send:output_type -> hedgeproto.Payload
	0, // 7: hedgeproto.Hedge.Broadcast:output_type -> hedgeproto.Payload
	0, // 8: hedgeproto.Hedge.SoSWrite:output_type -> hedgeproto.Payload
	0, // 9: hedgeproto.Hedge.SoSRead:output_type -> hedgeproto.Payload
	0, // 10: hedgeproto.Hedge.SoSClose:output_type -> hedgeproto.Payload
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_default_proto_init() }
func file_default_proto_init() {
	if File_default_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_default_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_default_proto_goTypes,
		DependencyIndexes: file_default_proto_depIdxs,
		MessageInfos:      file_default_proto_msgTypes,
	}.Build()
	File_default_proto = out.File
	file_default_proto_rawDesc = nil
	file_default_proto_goTypes = nil
	file_default_proto_depIdxs = nil
}
