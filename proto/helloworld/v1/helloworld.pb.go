// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: helloworld/v1/helloworld.proto

package helloworldv1

import (
	_ "github.com/douyu/jupiter/proto/fieldmask/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type
type Type int32

const (
	// TYPE_UNSPECIFIED ...
	Type_TYPE_UNSPECIFIED Type = 0
	// TYPE_Filter ...
	Type_TYPE_Filter Type = 1
	// TYPE_Prune ...
	Type_TYPE_Prune Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_Filter",
		2: "TYPE_Prune",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_Filter":      1,
		"TYPE_Prune":       2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_v1_helloworld_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_helloworld_v1_helloworld_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}

// Sex
type Sex int32

const (
	// SEX_UNSPECIFIED ...
	Sex_SEX_UNSPECIFIED Sex = 0
	// SEX_MALE ...
	Sex_SEX_MALE Sex = 1
	// SEX_FEMALE ...
	Sex_SEX_FEMALE Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "SEX_UNSPECIFIED",
		1: "SEX_MALE",
		2: "SEX_FEMALE",
	}
	Sex_value = map[string]int32{
		"SEX_UNSPECIFIED": 0,
		"SEX_MALE":        1,
		"SEX_FEMALE":      2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_v1_helloworld_proto_enumTypes[1].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_helloworld_v1_helloworld_proto_enumTypes[1]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1}
}

// The request message containing the user's name.
type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name ...
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name" param:"name" uri:"name"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error
	Error uint32 `protobuf:"varint,1,opt,name=error,proto3" json:"error"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// data ...
	Data *SayHelloResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SayHelloResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SayHelloResponse) GetData() *SayHelloResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// The request message containing the user's name.
type SayHiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name ...
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
}

func (x *SayHiRequest) Reset() {
	*x = SayHiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHiRequest) ProtoMessage() {}

func (x *SayHiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHiRequest.ProtoReflect.Descriptor instead.
func (*SayHiRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *SayHiRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type SayHiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error
	Error uint32 `protobuf:"varint,1,opt,name=error,proto3" json:"error"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// data ...
	Data *SayHiResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *SayHiResponse) Reset() {
	*x = SayHiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHiResponse) ProtoMessage() {}

func (x *SayHiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHiResponse.ProtoReflect.Descriptor instead.
func (*SayHiResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *SayHiResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SayHiResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SayHiResponse) GetData() *SayHiResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type SayGoodByeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Age        uint64                 `protobuf:"varint,2,opt,name=age,proto3" json:"age"`
	Type       Type                   `protobuf:"varint,3,opt,name=type,proto3,enum=helloworld.v1.Type" json:"type"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"updateMask"`
}

func (x *SayGoodByeRequest) Reset() {
	*x = SayGoodByeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodByeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodByeRequest) ProtoMessage() {}

func (x *SayGoodByeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodByeRequest.ProtoReflect.Descriptor instead.
func (*SayGoodByeRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{4}
}

func (x *SayGoodByeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayGoodByeRequest) GetAge() uint64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *SayGoodByeRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *SayGoodByeRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type SayGoodByeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error 错误代码
	Error uint32 `protobuf:"varint,1,opt,name=error,proto3" json:"error"`
	// msg 错误信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// data 响应数据
	Data *SayGoodByeResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *SayGoodByeResponse) Reset() {
	*x = SayGoodByeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodByeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodByeResponse) ProtoMessage() {}

func (x *SayGoodByeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodByeResponse.ProtoReflect.Descriptor instead.
func (*SayGoodByeResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{5}
}

func (x *SayGoodByeResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SayGoodByeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SayGoodByeResponse) GetData() *SayGoodByeResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type OtherHelloMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
}

func (x *OtherHelloMessage) Reset() {
	*x = OtherHelloMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherHelloMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherHelloMessage) ProtoMessage() {}

func (x *OtherHelloMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherHelloMessage.ProtoReflect.Descriptor instead.
func (*OtherHelloMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{6}
}

func (x *OtherHelloMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OtherHelloMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Data is the data to be sent.
type SayHelloResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// age_number is the age number.
	AgeNumber uint64 `protobuf:"varint,2,opt,name=age_number,json=ageNumber,proto3" json:"ageNumber"`
	// sex is the user's sex
	Sex Sex `protobuf:"varint,3,opt,name=sex,proto3,enum=helloworld.v1.Sex" json:"sex"`
	// metadata is the user's metadata
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SayHelloResponse_Data) Reset() {
	*x = SayHelloResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse_Data) ProtoMessage() {}

func (x *SayHelloResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse_Data.ProtoReflect.Descriptor instead.
func (*SayHelloResponse_Data) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SayHelloResponse_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayHelloResponse_Data) GetAgeNumber() uint64 {
	if x != nil {
		return x.AgeNumber
	}
	return 0
}

func (x *SayHelloResponse_Data) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_SEX_UNSPECIFIED
}

func (x *SayHelloResponse_Data) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Data is the data to be sent.
type SayHiResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// age_number is the age number.
	AgeNumber uint64 `protobuf:"varint,2,opt,name=age_number,json=ageNumber,proto3" json:"ageNumber"`
}

func (x *SayHiResponse_Data) Reset() {
	*x = SayHiResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHiResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHiResponse_Data) ProtoMessage() {}

func (x *SayHiResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHiResponse_Data.ProtoReflect.Descriptor instead.
func (*SayHiResponse_Data) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SayHiResponse_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayHiResponse_Data) GetAgeNumber() uint64 {
	if x != nil {
		return x.AgeNumber
	}
	return 0
}

// Data 响应数据
type SayGoodByeResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age   uint64             `protobuf:"varint,1,opt,name=age,proto3" json:"age"`
	Name  string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Other *OtherHelloMessage `protobuf:"bytes,3,opt,name=other,proto3" json:"other"`
}

func (x *SayGoodByeResponse_Data) Reset() {
	*x = SayGoodByeResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodByeResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodByeResponse_Data) ProtoMessage() {}

func (x *SayGoodByeResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodByeResponse_Data.ProtoReflect.Descriptor instead.
func (*SayGoodByeResponse_Data) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{5, 0}
}

func (x *SayGoodByeResponse_Data) GetAge() uint64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *SayGoodByeResponse_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayGoodByeResponse_Data) GetOther() *OtherHelloMessage {
	if x != nil {
		return x.Other
	}
	return nil
}

var File_helloworld_v1_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_v1_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x28, 0x9a, 0x84, 0x9e, 0x03, 0x23, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20,
	0x75, 0x72, 0x69, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xf5, 0x02, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xfe, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10,
	0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x4e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x34, 0x0a, 0x0c, 0x53, 0x61, 0x79, 0x48,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa9,
	0x01, 0x0a, 0x0d, 0x53, 0x61, 0x79, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x53,
	0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x08, 0xc2, 0xd3, 0x04, 0x04, 0x08, 0x01, 0x10, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xde, 0x01, 0x0a, 0x12, 0x53, 0x61, 0x79, 0x47, 0x6f,
	0x6f, 0x64, 0x42, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x64, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x36, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x11, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x3d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x72,
	0x75, 0x6e, 0x65, 0x10, 0x02, 0x2a, 0x38, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x45, 0x58, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x58, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x58, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x32,
	0xfb, 0x02, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xcf, 0x01, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x81, 0x01, 0x92, 0x41, 0x2a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20,
	0x30, 0x30, 0x31, 0x72, 0x1b, 0x0a, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x28, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4e, 0x3a, 0x01, 0x2a, 0x5a, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x44, 0x0a, 0x05, 0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x1b, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x79, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x53, 0x61,
	0x79, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64,
	0x42, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x47, 0x6f,
	0x6f, 0x64, 0x42, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbb, 0x02,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x75,
	0x2f, 0x6a, 0x75, 0x70, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x76, 0x31, 0x92, 0x41, 0xce, 0x01, 0x12, 0x50, 0x0a,
	0x1e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x20, 0x64, 0x65, 0x6d, 0x6f,
	0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x29, 0x0a, 0x05, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x12, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75,
	0x79, 0x75, 0x2f, 0x6a, 0x75, 0x70, 0x69, 0x74, 0x65, 0x72, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x32,
	0x21, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x77,
	0x77, 0x77, 0x2d, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x75, 0x72, 0x6c, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x6a, 0x57, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x30, 0x30,
	0x31, 0x12, 0x15, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x68, 0x6f, 0x77, 0x20, 0x77,
	0x65, 0x20, 0x64, 0x6f, 0x20, 0x69, 0x74, 0x2e, 0x1a, 0x31, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64,
	0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f,
	0x75, 0x79, 0x75, 0x2f, 0x6a, 0x75, 0x70, 0x69, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_helloworld_v1_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_v1_helloworld_proto_rawDescData = file_helloworld_v1_helloworld_proto_rawDesc
)

func file_helloworld_v1_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_v1_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_helloworld_proto_rawDescData)
	})
	return file_helloworld_v1_helloworld_proto_rawDescData
}

var file_helloworld_v1_helloworld_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_helloworld_v1_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_helloworld_v1_helloworld_proto_goTypes = []interface{}{
	(Type)(0),                       // 0: helloworld.v1.Type
	(Sex)(0),                        // 1: helloworld.v1.Sex
	(*SayHelloRequest)(nil),         // 2: helloworld.v1.SayHelloRequest
	(*SayHelloResponse)(nil),        // 3: helloworld.v1.SayHelloResponse
	(*SayHiRequest)(nil),            // 4: helloworld.v1.SayHiRequest
	(*SayHiResponse)(nil),           // 5: helloworld.v1.SayHiResponse
	(*SayGoodByeRequest)(nil),       // 6: helloworld.v1.SayGoodByeRequest
	(*SayGoodByeResponse)(nil),      // 7: helloworld.v1.SayGoodByeResponse
	(*OtherHelloMessage)(nil),       // 8: helloworld.v1.OtherHelloMessage
	(*SayHelloResponse_Data)(nil),   // 9: helloworld.v1.SayHelloResponse.Data
	nil,                             // 10: helloworld.v1.SayHelloResponse.Data.MetadataEntry
	(*SayHiResponse_Data)(nil),      // 11: helloworld.v1.SayHiResponse.Data
	(*SayGoodByeResponse_Data)(nil), // 12: helloworld.v1.SayGoodByeResponse.Data
	(*fieldmaskpb.FieldMask)(nil),   // 13: google.protobuf.FieldMask
}
var file_helloworld_v1_helloworld_proto_depIdxs = []int32{
	9,  // 0: helloworld.v1.SayHelloResponse.data:type_name -> helloworld.v1.SayHelloResponse.Data
	11, // 1: helloworld.v1.SayHiResponse.data:type_name -> helloworld.v1.SayHiResponse.Data
	0,  // 2: helloworld.v1.SayGoodByeRequest.type:type_name -> helloworld.v1.Type
	13, // 3: helloworld.v1.SayGoodByeRequest.update_mask:type_name -> google.protobuf.FieldMask
	12, // 4: helloworld.v1.SayGoodByeResponse.data:type_name -> helloworld.v1.SayGoodByeResponse.Data
	1,  // 5: helloworld.v1.SayHelloResponse.Data.sex:type_name -> helloworld.v1.Sex
	10, // 6: helloworld.v1.SayHelloResponse.Data.metadata:type_name -> helloworld.v1.SayHelloResponse.Data.MetadataEntry
	8,  // 7: helloworld.v1.SayGoodByeResponse.Data.other:type_name -> helloworld.v1.OtherHelloMessage
	2,  // 8: helloworld.v1.GreeterService.SayHello:input_type -> helloworld.v1.SayHelloRequest
	4,  // 9: helloworld.v1.GreeterService.SayHi:input_type -> helloworld.v1.SayHiRequest
	6,  // 10: helloworld.v1.GreeterService.SayGoodBye:input_type -> helloworld.v1.SayGoodByeRequest
	3,  // 11: helloworld.v1.GreeterService.SayHello:output_type -> helloworld.v1.SayHelloResponse
	5,  // 12: helloworld.v1.GreeterService.SayHi:output_type -> helloworld.v1.SayHiResponse
	7,  // 13: helloworld.v1.GreeterService.SayGoodBye:output_type -> helloworld.v1.SayGoodByeResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_helloworld_v1_helloworld_proto_init() }
func file_helloworld_v1_helloworld_proto_init() {
	if File_helloworld_v1_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_v1_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHiRequest); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHiResponse); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodByeRequest); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodByeResponse); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherHelloMessage); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse_Data); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHiResponse_Data); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodByeResponse_Data); i {
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
			RawDescriptor: file_helloworld_v1_helloworld_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_v1_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_helloworld_proto_depIdxs,
		EnumInfos:         file_helloworld_v1_helloworld_proto_enumTypes,
		MessageInfos:      file_helloworld_v1_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_v1_helloworld_proto = out.File
	file_helloworld_v1_helloworld_proto_rawDesc = nil
	file_helloworld_v1_helloworld_proto_goTypes = nil
	file_helloworld_v1_helloworld_proto_depIdxs = nil
}
