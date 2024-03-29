// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: player_service.proto

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

type Direction int32

const (
	Direction_LEFT  Direction = 0
	Direction_UP    Direction = 1
	Direction_RIGHT Direction = 2
	Direction_DOWN  Direction = 3
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "LEFT",
		1: "UP",
		2: "RIGHT",
		3: "DOWN",
	}
	Direction_value = map[string]int32{
		"LEFT":  0,
		"UP":    1,
		"RIGHT": 2,
		"DOWN":  3,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_player_service_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_player_service_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_player_service_proto_rawDescGZIP(), []int{0}
}

type MovePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Direction Direction `protobuf:"varint,2,opt,name=Direction,proto3,enum=Direction" json:"Direction,omitempty"`
}

func (x *MovePlayerRequest) Reset() {
	*x = MovePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePlayerRequest) ProtoMessage() {}

func (x *MovePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePlayerRequest.ProtoReflect.Descriptor instead.
func (*MovePlayerRequest) Descriptor() ([]byte, []int) {
	return file_player_service_proto_rawDescGZIP(), []int{0}
}

func (x *MovePlayerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MovePlayerRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_LEFT
}

type MovePlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map string `protobuf:"bytes,1,opt,name=Map,proto3" json:"Map,omitempty"`
}

func (x *MovePlayerResponse) Reset() {
	*x = MovePlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePlayerResponse) ProtoMessage() {}

func (x *MovePlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_player_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePlayerResponse.ProtoReflect.Descriptor instead.
func (*MovePlayerResponse) Descriptor() ([]byte, []int) {
	return file_player_service_proto_rawDescGZIP(), []int{1}
}

func (x *MovePlayerResponse) GetMap() string {
	if x != nil {
		return x.Map
	}
	return ""
}

var File_player_service_proto protoreflect.FileDescriptor

var file_player_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x61, 0x70, 0x2a, 0x32, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45,
	0x46, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10,
	0x03, 0x32, 0x48, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x12, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x6b, 0x6f, 0x62, 0x44,
	0x46, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x65, 0x6e, 0x6e, 0x2d, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_service_proto_rawDescOnce sync.Once
	file_player_service_proto_rawDescData = file_player_service_proto_rawDesc
)

func file_player_service_proto_rawDescGZIP() []byte {
	file_player_service_proto_rawDescOnce.Do(func() {
		file_player_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_service_proto_rawDescData)
	})
	return file_player_service_proto_rawDescData
}

var file_player_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_player_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_player_service_proto_goTypes = []interface{}{
	(Direction)(0),             // 0: Direction
	(*MovePlayerRequest)(nil),  // 1: MovePlayerRequest
	(*MovePlayerResponse)(nil), // 2: MovePlayerResponse
}
var file_player_service_proto_depIdxs = []int32{
	0, // 0: MovePlayerRequest.Direction:type_name -> Direction
	1, // 1: PlayerService.MovePlayer:input_type -> MovePlayerRequest
	2, // 2: PlayerService.MovePlayer:output_type -> MovePlayerResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_player_service_proto_init() }
func file_player_service_proto_init() {
	if File_player_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePlayerRequest); i {
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
		file_player_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePlayerResponse); i {
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
			RawDescriptor: file_player_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_service_proto_goTypes,
		DependencyIndexes: file_player_service_proto_depIdxs,
		EnumInfos:         file_player_service_proto_enumTypes,
		MessageInfos:      file_player_service_proto_msgTypes,
	}.Build()
	File_player_service_proto = out.File
	file_player_service_proto_rawDesc = nil
	file_player_service_proto_goTypes = nil
	file_player_service_proto_depIdxs = nil
}
