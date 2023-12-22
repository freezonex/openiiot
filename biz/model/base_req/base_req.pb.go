// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: freezonex/common/base_req.proto

package base_req

import (
	_ "freezonex/openiiot/biz/model/api"
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

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count          int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count" form:"count" query:"count"`     // number of elements requested
	Cursor         int32  `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor" form:"cursor" query:"cursor"` // offset to start from
	SortKey        string `protobuf:"bytes,3,opt,name=sort_key,json=sortKey,proto3" json:"sort_key" form:"sort_key" query:"sort_key"`
	SortDescending bool   `protobuf:"varint,4,opt,name=sort_descending,json=sortDescending,proto3" json:"sort_descending" form:"sort_descending" query:"sort_descending"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_common_base_req_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_common_base_req_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_common_base_req_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PaginationRequest) GetCursor() int32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *PaginationRequest) GetSortKey() string {
	if x != nil {
		return x.SortKey
	}
	return ""
}

func (x *PaginationRequest) GetSortDescending() bool {
	if x != nil {
		return x.SortDescending
	}
	return false
}

// All request will contain BaseRequest
type BaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token" header:"X-Custom-Token"`
}

func (x *BaseRequest) Reset() {
	*x = BaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_common_base_req_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseRequest) ProtoMessage() {}

func (x *BaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_common_base_req_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseRequest.ProtoReflect.Descriptor instead.
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_common_base_req_proto_rawDescGZIP(), []int{1}
}

func (x *BaseRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

var File_freezonex_common_base_req_proto protoreflect.FileDescriptor

var file_freezonex_common_base_req_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x1a, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0x40, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x12, 0xba, 0xbb, 0x18, 0x0e, 0x58, 0x2d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x2d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x27, 0x5a, 0x25, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_freezonex_common_base_req_proto_rawDescOnce sync.Once
	file_freezonex_common_base_req_proto_rawDescData = file_freezonex_common_base_req_proto_rawDesc
)

func file_freezonex_common_base_req_proto_rawDescGZIP() []byte {
	file_freezonex_common_base_req_proto_rawDescOnce.Do(func() {
		file_freezonex_common_base_req_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_common_base_req_proto_rawDescData)
	})
	return file_freezonex_common_base_req_proto_rawDescData
}

var file_freezonex_common_base_req_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_freezonex_common_base_req_proto_goTypes = []interface{}{
	(*PaginationRequest)(nil), // 0: base_req.PaginationRequest
	(*BaseRequest)(nil),       // 1: base_req.BaseRequest
}
var file_freezonex_common_base_req_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_freezonex_common_base_req_proto_init() }
func file_freezonex_common_base_req_proto_init() {
	if File_freezonex_common_base_req_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_common_base_req_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
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
		file_freezonex_common_base_req_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseRequest); i {
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
			RawDescriptor: file_freezonex_common_base_req_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_common_base_req_proto_goTypes,
		DependencyIndexes: file_freezonex_common_base_req_proto_depIdxs,
		MessageInfos:      file_freezonex_common_base_req_proto_msgTypes,
	}.Build()
	File_freezonex_common_base_req_proto = out.File
	file_freezonex_common_base_req_proto_rawDesc = nil
	file_freezonex_common_base_req_proto_goTypes = nil
	file_freezonex_common_base_req_proto_depIdxs = nil
}
