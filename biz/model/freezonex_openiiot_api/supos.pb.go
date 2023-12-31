// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.4
// source: freezonex/openiiot_api/supos.proto

package freezonex_openiiot_api

import (
	_ "freezonex/openiiot/biz/model/api"
	base_req "freezonex/openiiot/biz/model/base_req"
	base_resp "freezonex/openiiot/biz/model/base_resp"
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

type GetSuposLoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
}

func (x *GetSuposLoginUserRequest) Reset() {
	*x = GetSuposLoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_supos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuposLoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuposLoginUserRequest) ProtoMessage() {}

func (x *GetSuposLoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_supos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuposLoginUserRequest.ProtoReflect.Descriptor instead.
func (*GetSuposLoginUserRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_supos_proto_rawDescGZIP(), []int{0}
}

func (x *GetSuposLoginUserRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

type GetSuposLoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	UserName string                  `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name" form:"user_name" query:"user_name"`
	UserRole string                  `protobuf:"bytes,3,opt,name=user_role,json=userRole,proto3" json:"user_role" form:"user_role" query:"user_role"`
}

func (x *GetSuposLoginUserResponse) Reset() {
	*x = GetSuposLoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_supos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuposLoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuposLoginUserResponse) ProtoMessage() {}

func (x *GetSuposLoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_supos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuposLoginUserResponse.ProtoReflect.Descriptor instead.
func (*GetSuposLoginUserResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_supos_proto_rawDescGZIP(), []int{1}
}

func (x *GetSuposLoginUserResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetSuposLoginUserResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetSuposLoginUserResponse) GetUserRole() string {
	if x != nil {
		return x.UserRole
	}
	return ""
}

var File_freezonex_openiiot_api_supos_proto protoreflect.FileDescriptor

var file_freezonex_openiiot_api_supos_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x70, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x66, 0x72,
	0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x72, 0x65, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x6f, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x6f, 0x73, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x35, 0x5a, 0x33, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x69, 0x6f, 0x74, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69,
	0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freezonex_openiiot_api_supos_proto_rawDescOnce sync.Once
	file_freezonex_openiiot_api_supos_proto_rawDescData = file_freezonex_openiiot_api_supos_proto_rawDesc
)

func file_freezonex_openiiot_api_supos_proto_rawDescGZIP() []byte {
	file_freezonex_openiiot_api_supos_proto_rawDescOnce.Do(func() {
		file_freezonex_openiiot_api_supos_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_openiiot_api_supos_proto_rawDescData)
	})
	return file_freezonex_openiiot_api_supos_proto_rawDescData
}

var file_freezonex_openiiot_api_supos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_freezonex_openiiot_api_supos_proto_goTypes = []interface{}{
	(*GetSuposLoginUserRequest)(nil),  // 0: freezonex.openiiot_api.GetSuposLoginUserRequest
	(*GetSuposLoginUserResponse)(nil), // 1: freezonex.openiiot_api.GetSuposLoginUserResponse
	(*base_req.BaseRequest)(nil),      // 2: base_req.BaseRequest
	(*base_resp.BaseResponse)(nil),    // 3: base_resp.BaseResponse
}
var file_freezonex_openiiot_api_supos_proto_depIdxs = []int32{
	2, // 0: freezonex.openiiot_api.GetSuposLoginUserRequest.base_request:type_name -> base_req.BaseRequest
	3, // 1: freezonex.openiiot_api.GetSuposLoginUserResponse.base_resp:type_name -> base_resp.BaseResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_freezonex_openiiot_api_supos_proto_init() }
func file_freezonex_openiiot_api_supos_proto_init() {
	if File_freezonex_openiiot_api_supos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_openiiot_api_supos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuposLoginUserRequest); i {
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
		file_freezonex_openiiot_api_supos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuposLoginUserResponse); i {
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
			RawDescriptor: file_freezonex_openiiot_api_supos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_openiiot_api_supos_proto_goTypes,
		DependencyIndexes: file_freezonex_openiiot_api_supos_proto_depIdxs,
		MessageInfos:      file_freezonex_openiiot_api_supos_proto_msgTypes,
	}.Build()
	File_freezonex_openiiot_api_supos_proto = out.File
	file_freezonex_openiiot_api_supos_proto_rawDesc = nil
	file_freezonex_openiiot_api_supos_proto_goTypes = nil
	file_freezonex_openiiot_api_supos_proto_depIdxs = nil
}
