// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: level_service.proto

package rpc

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

type CreateLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level []*Rows `protobuf:"bytes,1,rep,name=Level,proto3" json:"Level,omitempty"`
}

func (x *CreateLevelRequest) Reset() {
	*x = CreateLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_level_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLevelRequest) ProtoMessage() {}

func (x *CreateLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_level_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLevelRequest.ProtoReflect.Descriptor instead.
func (*CreateLevelRequest) Descriptor() ([]byte, []int) {
	return file_level_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLevelRequest) GetLevel() []*Rows {
	if x != nil {
		return x.Level
	}
	return nil
}

type Rows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []int32 `protobuf:"varint,1,rep,packed,name=Cells,proto3" json:"Cells,omitempty"`
}

func (x *Rows) Reset() {
	*x = Rows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_level_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rows) ProtoMessage() {}

func (x *Rows) ProtoReflect() protoreflect.Message {
	mi := &file_level_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rows.ProtoReflect.Descriptor instead.
func (*Rows) Descriptor() ([]byte, []int) {
	return file_level_service_proto_rawDescGZIP(), []int{1}
}

func (x *Rows) GetCells() []int32 {
	if x != nil {
		return x.Cells
	}
	return nil
}

type CreateLevelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *CreateLevelResponse) Reset() {
	*x = CreateLevelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_level_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLevelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLevelResponse) ProtoMessage() {}

func (x *CreateLevelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_level_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLevelResponse.ProtoReflect.Descriptor instead.
func (*CreateLevelResponse) Descriptor() ([]byte, []int) {
	return file_level_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLevelResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_level_service_proto protoreflect.FileDescriptor

var file_level_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x77,
	0x73, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x1c, 0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x43, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x05, 0x43, 0x65, 0x6c, 0x6c, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x32, 0x4a, 0x0a,
	0x0c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x6b, 0x6f, 0x62, 0x44, 0x46, 0x72,
	0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x65, 0x6e, 0x6e, 0x2d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69,
	0x6b, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_level_service_proto_rawDescOnce sync.Once
	file_level_service_proto_rawDescData = file_level_service_proto_rawDesc
)

func file_level_service_proto_rawDescGZIP() []byte {
	file_level_service_proto_rawDescOnce.Do(func() {
		file_level_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_level_service_proto_rawDescData)
	})
	return file_level_service_proto_rawDescData
}

var file_level_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_level_service_proto_goTypes = []interface{}{
	(*CreateLevelRequest)(nil),  // 0: CreateLevelRequest
	(*Rows)(nil),                // 1: Rows
	(*CreateLevelResponse)(nil), // 2: CreateLevelResponse
}
var file_level_service_proto_depIdxs = []int32{
	1, // 0: CreateLevelRequest.Level:type_name -> Rows
	0, // 1: LevelService.CreateLevel:input_type -> CreateLevelRequest
	2, // 2: LevelService.CreateLevel:output_type -> CreateLevelResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_level_service_proto_init() }
func file_level_service_proto_init() {
	if File_level_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_level_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLevelRequest); i {
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
		file_level_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rows); i {
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
		file_level_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLevelResponse); i {
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
			RawDescriptor: file_level_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_level_service_proto_goTypes,
		DependencyIndexes: file_level_service_proto_depIdxs,
		MessageInfos:      file_level_service_proto_msgTypes,
	}.Build()
	File_level_service_proto = out.File
	file_level_service_proto_rawDesc = nil
	file_level_service_proto_goTypes = nil
	file_level_service_proto_depIdxs = nil
}
