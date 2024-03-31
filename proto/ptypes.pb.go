// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: proto/ptypes.proto

package types

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

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ptypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ptypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_proto_ptypes_proto_rawDescGZIP(), []int{0}
}

type InsertDriverLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32   `protobuf:"varint,1,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
	Lat      float64 `protobuf:"fixed64,2,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lon      float64 `protobuf:"fixed64,3,opt,name=Lon,proto3" json:"Lon,omitempty"`
}

func (x *InsertDriverLogRequest) Reset() {
	*x = InsertDriverLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ptypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertDriverLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertDriverLogRequest) ProtoMessage() {}

func (x *InsertDriverLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ptypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertDriverLogRequest.ProtoReflect.Descriptor instead.
func (*InsertDriverLogRequest) Descriptor() ([]byte, []int) {
	return file_proto_ptypes_proto_rawDescGZIP(), []int{1}
}

func (x *InsertDriverLogRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *InsertDriverLogRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *InsertDriverLogRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type InsertDriverLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *InsertDriverLogResponse) Reset() {
	*x = InsertDriverLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ptypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertDriverLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertDriverLogResponse) ProtoMessage() {}

func (x *InsertDriverLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ptypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertDriverLogResponse.ProtoReflect.Descriptor instead.
func (*InsertDriverLogResponse) Descriptor() ([]byte, []int) {
	return file_proto_ptypes_proto_rawDescGZIP(), []int{2}
}

func (x *InsertDriverLogResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_ptypes_proto protoreflect.FileDescriptor

var file_proto_ptypes_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x22, 0x58, 0x0a, 0x16,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x4c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x4c, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x17, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x55, 0x0a, 0x0d, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0f,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x12,
	0x17, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x79, 0x72, 0x6f, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x74,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ptypes_proto_rawDescOnce sync.Once
	file_proto_ptypes_proto_rawDescData = file_proto_ptypes_proto_rawDesc
)

func file_proto_ptypes_proto_rawDescGZIP() []byte {
	file_proto_ptypes_proto_rawDescOnce.Do(func() {
		file_proto_ptypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ptypes_proto_rawDescData)
	})
	return file_proto_ptypes_proto_rawDescData
}

var file_proto_ptypes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ptypes_proto_goTypes = []interface{}{
	(*None)(nil),                    // 0: None
	(*InsertDriverLogRequest)(nil),  // 1: InsertDriverLogRequest
	(*InsertDriverLogResponse)(nil), // 2: InsertDriverLogResponse
}
var file_proto_ptypes_proto_depIdxs = []int32{
	1, // 0: DriverService.InsertDriverLog:input_type -> InsertDriverLogRequest
	2, // 1: DriverService.InsertDriverLog:output_type -> InsertDriverLogResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ptypes_proto_init() }
func file_proto_ptypes_proto_init() {
	if File_proto_ptypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ptypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*None); i {
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
		file_proto_ptypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertDriverLogRequest); i {
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
		file_proto_ptypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertDriverLogResponse); i {
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
			RawDescriptor: file_proto_ptypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ptypes_proto_goTypes,
		DependencyIndexes: file_proto_ptypes_proto_depIdxs,
		MessageInfos:      file_proto_ptypes_proto_msgTypes,
	}.Build()
	File_proto_ptypes_proto = out.File
	file_proto_ptypes_proto_rawDesc = nil
	file_proto_ptypes_proto_goTypes = nil
	file_proto_ptypes_proto_depIdxs = nil
}
