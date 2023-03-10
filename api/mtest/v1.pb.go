// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.6
// source: api/mtest/v1.proto

package mtest

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateV1Request) Reset() {
	*x = CreateV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateV1Request) ProtoMessage() {}

func (x *CreateV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateV1Request.ProtoReflect.Descriptor instead.
func (*CreateV1Request) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{0}
}

type CreateV1Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateV1Reply) Reset() {
	*x = CreateV1Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateV1Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateV1Reply) ProtoMessage() {}

func (x *CreateV1Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateV1Reply.ProtoReflect.Descriptor instead.
func (*CreateV1Reply) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{1}
}

type UpdateV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateV1Request) Reset() {
	*x = UpdateV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateV1Request) ProtoMessage() {}

func (x *UpdateV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateV1Request.ProtoReflect.Descriptor instead.
func (*UpdateV1Request) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{2}
}

type UpdateV1Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateV1Reply) Reset() {
	*x = UpdateV1Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateV1Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateV1Reply) ProtoMessage() {}

func (x *UpdateV1Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateV1Reply.ProtoReflect.Descriptor instead.
func (*UpdateV1Reply) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{3}
}

type DeleteV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteV1Request) Reset() {
	*x = DeleteV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteV1Request) ProtoMessage() {}

func (x *DeleteV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteV1Request.ProtoReflect.Descriptor instead.
func (*DeleteV1Request) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{4}
}

type DeleteV1Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteV1Reply) Reset() {
	*x = DeleteV1Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteV1Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteV1Reply) ProtoMessage() {}

func (x *DeleteV1Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteV1Reply.ProtoReflect.Descriptor instead.
func (*DeleteV1Reply) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{5}
}

type GetV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetV1Request) Reset() {
	*x = GetV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetV1Request) ProtoMessage() {}

func (x *GetV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetV1Request.ProtoReflect.Descriptor instead.
func (*GetV1Request) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{6}
}

type GetV1Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetV1Reply) Reset() {
	*x = GetV1Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetV1Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetV1Reply) ProtoMessage() {}

func (x *GetV1Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetV1Reply.ProtoReflect.Descriptor instead.
func (*GetV1Reply) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{7}
}

type ListV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListV1Request) Reset() {
	*x = ListV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListV1Request) ProtoMessage() {}

func (x *ListV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListV1Request.ProtoReflect.Descriptor instead.
func (*ListV1Request) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{8}
}

type ListV1Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListV1Reply) Reset() {
	*x = ListV1Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mtest_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListV1Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListV1Reply) ProtoMessage() {}

func (x *ListV1Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_mtest_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListV1Reply.ProtoReflect.Descriptor instead.
func (*ListV1Reply) Descriptor() ([]byte, []int) {
	return file_api_mtest_v1_proto_rawDescGZIP(), []int{9}
}

var File_api_mtest_v1_proto protoreflect.FileDescriptor

var file_api_mtest_v1_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x31,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xcf, 0x02, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x50,
	0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x31, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x6d, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x40, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x31, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x31, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x56, 0x31, 0x12, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a,
	0x06, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x31, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x31, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x0a, 0x09, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x50, 0x01, 0x5a, 0x1a, 0x6d, 0x63, 0x70, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x6d,
	0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_mtest_v1_proto_rawDescOnce sync.Once
	file_api_mtest_v1_proto_rawDescData = file_api_mtest_v1_proto_rawDesc
)

func file_api_mtest_v1_proto_rawDescGZIP() []byte {
	file_api_mtest_v1_proto_rawDescOnce.Do(func() {
		file_api_mtest_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mtest_v1_proto_rawDescData)
	})
	return file_api_mtest_v1_proto_rawDescData
}

var file_api_mtest_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_mtest_v1_proto_goTypes = []interface{}{
	(*CreateV1Request)(nil), // 0: api.mtest.CreateV1Request
	(*CreateV1Reply)(nil),   // 1: api.mtest.CreateV1Reply
	(*UpdateV1Request)(nil), // 2: api.mtest.UpdateV1Request
	(*UpdateV1Reply)(nil),   // 3: api.mtest.UpdateV1Reply
	(*DeleteV1Request)(nil), // 4: api.mtest.DeleteV1Request
	(*DeleteV1Reply)(nil),   // 5: api.mtest.DeleteV1Reply
	(*GetV1Request)(nil),    // 6: api.mtest.GetV1Request
	(*GetV1Reply)(nil),      // 7: api.mtest.GetV1Reply
	(*ListV1Request)(nil),   // 8: api.mtest.ListV1Request
	(*ListV1Reply)(nil),     // 9: api.mtest.ListV1Reply
}
var file_api_mtest_v1_proto_depIdxs = []int32{
	0, // 0: api.mtest.V1.CreateV1:input_type -> api.mtest.CreateV1Request
	2, // 1: api.mtest.V1.UpdateV1:input_type -> api.mtest.UpdateV1Request
	4, // 2: api.mtest.V1.DeleteV1:input_type -> api.mtest.DeleteV1Request
	6, // 3: api.mtest.V1.GetV1:input_type -> api.mtest.GetV1Request
	8, // 4: api.mtest.V1.ListV1:input_type -> api.mtest.ListV1Request
	1, // 5: api.mtest.V1.CreateV1:output_type -> api.mtest.CreateV1Reply
	3, // 6: api.mtest.V1.UpdateV1:output_type -> api.mtest.UpdateV1Reply
	5, // 7: api.mtest.V1.DeleteV1:output_type -> api.mtest.DeleteV1Reply
	7, // 8: api.mtest.V1.GetV1:output_type -> api.mtest.GetV1Reply
	9, // 9: api.mtest.V1.ListV1:output_type -> api.mtest.ListV1Reply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_mtest_v1_proto_init() }
func file_api_mtest_v1_proto_init() {
	if File_api_mtest_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_mtest_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateV1Request); i {
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
		file_api_mtest_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateV1Reply); i {
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
		file_api_mtest_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateV1Request); i {
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
		file_api_mtest_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateV1Reply); i {
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
		file_api_mtest_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteV1Request); i {
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
		file_api_mtest_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteV1Reply); i {
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
		file_api_mtest_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetV1Request); i {
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
		file_api_mtest_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetV1Reply); i {
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
		file_api_mtest_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListV1Request); i {
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
		file_api_mtest_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListV1Reply); i {
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
			RawDescriptor: file_api_mtest_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mtest_v1_proto_goTypes,
		DependencyIndexes: file_api_mtest_v1_proto_depIdxs,
		MessageInfos:      file_api_mtest_v1_proto_msgTypes,
	}.Build()
	File_api_mtest_v1_proto = out.File
	file_api_mtest_v1_proto_rawDesc = nil
	file_api_mtest_v1_proto_goTypes = nil
	file_api_mtest_v1_proto_depIdxs = nil
}
