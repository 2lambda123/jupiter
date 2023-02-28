// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: fieldmask/v1/option.proto

package fieldmaskv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FieldMask: field mask rules
type FieldMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In  bool `protobuf:"varint,1,opt,name=in,proto3" json:"in"`
	Out bool `protobuf:"varint,2,opt,name=out,proto3" json:"out"`
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fieldmask_v1_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_v1_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_fieldmask_v1_option_proto_rawDescGZIP(), []int{0}
}

func (x *FieldMask) GetIn() bool {
	if x != nil {
		return x.In
	}
	return false
}

func (x *FieldMask) GetOut() bool {
	if x != nil {
		return x.Out
	}
	return false
}

var file_fieldmask_v1_option_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldMask)(nil),
		Field:         9528,
		Name:          "fieldmask.v1.Option",
		Tag:           "bytes,9528,opt,name=Option",
		Filename:      "fieldmask/v1/option.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional fieldmask.v1.FieldMask Option = 9528;
	E_Option = &file_fieldmask_v1_option_proto_extTypes[0]
)

var File_fieldmask_v1_option_proto protoreflect.FileDescriptor

var file_fieldmask_v1_option_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x09, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x3a, 0x52, 0x0a, 0x06, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x4a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x65,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x2e, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x2f, 0x6a,
	0x75, 0x70, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d,
	0x61, 0x73, 0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fieldmask_v1_option_proto_rawDescOnce sync.Once
	file_fieldmask_v1_option_proto_rawDescData = file_fieldmask_v1_option_proto_rawDesc
)

func file_fieldmask_v1_option_proto_rawDescGZIP() []byte {
	file_fieldmask_v1_option_proto_rawDescOnce.Do(func() {
		file_fieldmask_v1_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_fieldmask_v1_option_proto_rawDescData)
	})
	return file_fieldmask_v1_option_proto_rawDescData
}

var file_fieldmask_v1_option_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fieldmask_v1_option_proto_goTypes = []interface{}{
	(*FieldMask)(nil),                 // 0: fieldmask.v1.FieldMask
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_fieldmask_v1_option_proto_depIdxs = []int32{
	1, // 0: fieldmask.v1.Option:extendee -> google.protobuf.FieldOptions
	0, // 1: fieldmask.v1.Option:type_name -> fieldmask.v1.FieldMask
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fieldmask_v1_option_proto_init() }
func file_fieldmask_v1_option_proto_init() {
	if File_fieldmask_v1_option_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fieldmask_v1_option_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMask); i {
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
			RawDescriptor: file_fieldmask_v1_option_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_fieldmask_v1_option_proto_goTypes,
		DependencyIndexes: file_fieldmask_v1_option_proto_depIdxs,
		MessageInfos:      file_fieldmask_v1_option_proto_msgTypes,
		ExtensionInfos:    file_fieldmask_v1_option_proto_extTypes,
	}.Build()
	File_fieldmask_v1_option_proto = out.File
	file_fieldmask_v1_option_proto_rawDesc = nil
	file_fieldmask_v1_option_proto_goTypes = nil
	file_fieldmask_v1_option_proto_depIdxs = nil
}
