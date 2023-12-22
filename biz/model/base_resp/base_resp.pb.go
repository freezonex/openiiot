// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: freezonex/common/base_resp.proto

package base_resp

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

// All response will contain BaseResponse
type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"zigzag32,1,opt,name=status_code,json=statusCode,proto3" json:"status_code" form:"status_code" query:"status_code"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message" form:"status_message" query:"status_message"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_common_base_resp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_common_base_resp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_common_base_resp_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

// GeneralResponse is used for return early with only the BaseResponse
type GeneralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *GeneralResponse) Reset() {
	*x = GeneralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_common_base_resp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResponse) ProtoMessage() {}

func (x *GeneralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_common_base_resp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResponse.ProtoReflect.Descriptor instead.
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_common_base_resp_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralResponse) GetBaseResp() *BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count          int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count" form:"count" query:"count"`                                         // number of elements found
	Cursor         int32  `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor" form:"cursor" query:"cursor"`                                     // offset+count
	TotalCount     int64  `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count" form:"total_count" query:"total_count"` // total number of elements
	SortKey        string `protobuf:"bytes,4,opt,name=sort_key,json=sortKey,proto3" json:"sort_key" form:"sort_key" query:"sort_key"`
	SortDescending bool   `protobuf:"varint,5,opt,name=sort_descending,json=sortDescending,proto3" json:"sort_descending" form:"sort_descending" query:"sort_descending"`
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_common_base_resp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_common_base_resp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_common_base_resp_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PaginationResponse) GetCursor() int32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *PaginationResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *PaginationResponse) GetSortKey() string {
	if x != nil {
		return x.SortKey
	}
	return ""
}

func (x *PaginationResponse) GetSortDescending() bool {
	if x != nil {
		return x.SortDescending
	}
	return false
}

var File_freezonex_common_base_resp_proto protoreflect.FileDescriptor

var file_freezonex_common_base_resp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x22, 0x56, 0x0a,
	0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0xa7,
	0x01, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x65,
	0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x28, 0x5a, 0x26, 0x66, 0x72, 0x65, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freezonex_common_base_resp_proto_rawDescOnce sync.Once
	file_freezonex_common_base_resp_proto_rawDescData = file_freezonex_common_base_resp_proto_rawDesc
)

func file_freezonex_common_base_resp_proto_rawDescGZIP() []byte {
	file_freezonex_common_base_resp_proto_rawDescOnce.Do(func() {
		file_freezonex_common_base_resp_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_common_base_resp_proto_rawDescData)
	})
	return file_freezonex_common_base_resp_proto_rawDescData
}

var file_freezonex_common_base_resp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_freezonex_common_base_resp_proto_goTypes = []interface{}{
	(*BaseResponse)(nil),       // 0: base_resp.BaseResponse
	(*GeneralResponse)(nil),    // 1: base_resp.GeneralResponse
	(*PaginationResponse)(nil), // 2: base_resp.PaginationResponse
}
var file_freezonex_common_base_resp_proto_depIdxs = []int32{
	0, // 0: base_resp.GeneralResponse.base_resp:type_name -> base_resp.BaseResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_freezonex_common_base_resp_proto_init() }
func file_freezonex_common_base_resp_proto_init() {
	if File_freezonex_common_base_resp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_common_base_resp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_freezonex_common_base_resp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResponse); i {
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
		file_freezonex_common_base_resp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationResponse); i {
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
			RawDescriptor: file_freezonex_common_base_resp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_common_base_resp_proto_goTypes,
		DependencyIndexes: file_freezonex_common_base_resp_proto_depIdxs,
		MessageInfos:      file_freezonex_common_base_resp_proto_msgTypes,
	}.Build()
	File_freezonex_common_base_resp_proto = out.File
	file_freezonex_common_base_resp_proto_rawDesc = nil
	file_freezonex_common_base_resp_proto_goTypes = nil
	file_freezonex_common_base_resp_proto_depIdxs = nil
}
