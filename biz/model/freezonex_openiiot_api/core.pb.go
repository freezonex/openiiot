// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: freezonex/openiiot_api/core.proto

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

type Core struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    int64  `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Username    string `protobuf:"bytes,6,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password    string `protobuf:"bytes,7,opt,name=password,proto3" json:"password" form:"password" query:"password"`
	Type        string `protobuf:"bytes,8,opt,name=type,proto3" json:"type" form:"type" query:"type"`
	CreateTime  string `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time" form:"create_time" query:"create_time"`
	UpdateTime  string `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time" form:"update_time" query:"update_time"`
}

func (x *Core) Reset() {
	*x = Core{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Core) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Core) ProtoMessage() {}

func (x *Core) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Core.ProtoReflect.Descriptor instead.
func (*Core) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{0}
}

func (x *Core) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Core) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Core) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Core) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *Core) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Core) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Core) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Core) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Core) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Core) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AddCoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Name        string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    int64                 `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string                `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Username    string                `protobuf:"bytes,6,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password    string                `protobuf:"bytes,7,opt,name=password,proto3" json:"password" form:"password" query:"password"`
	Type        string                `protobuf:"bytes,8,opt,name=type,proto3" json:"type" form:"type" query:"type"`
}

func (x *AddCoreRequest) Reset() {
	*x = AddCoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCoreRequest) ProtoMessage() {}

func (x *AddCoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCoreRequest.ProtoReflect.Descriptor instead.
func (*AddCoreRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{1}
}

func (x *AddCoreRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *AddCoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCoreRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddCoreRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *AddCoreRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddCoreRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddCoreRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddCoreRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AddCoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Id       int64                   `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *AddCoreResponse) Reset() {
	*x = AddCoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCoreResponse) ProtoMessage() {}

func (x *AddCoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCoreResponse.ProtoReflect.Descriptor instead.
func (*AddCoreResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{2}
}

func (x *AddCoreResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *AddCoreResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	TenantId    int64                 `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string                `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Type        string                `protobuf:"bytes,6,opt,name=type,proto3" json:"type" form:"type" query:"type"`
}

func (x *GetCoreRequest) Reset() {
	*x = GetCoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoreRequest) ProtoMessage() {}

func (x *GetCoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoreRequest.ProtoReflect.Descriptor instead.
func (*GetCoreRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoreRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetCoreRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCoreRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *GetCoreRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetCoreRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetCoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Data     []*Core                 `protobuf:"bytes,2,rep,name=data,proto3" json:"data" form:"data" query:"data"`
}

func (x *GetCoreResponse) Reset() {
	*x = GetCoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoreResponse) ProtoMessage() {}

func (x *GetCoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoreResponse.ProtoReflect.Descriptor instead.
func (*GetCoreResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{4}
}

func (x *GetCoreResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetCoreResponse) GetData() []*Core {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateCoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    int64                 `protobuf:"varint,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string                `protobuf:"bytes,6,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Username    string                `protobuf:"bytes,7,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password    string                `protobuf:"bytes,8,opt,name=password,proto3" json:"password" form:"password" query:"password"`
	Type        string                `protobuf:"bytes,9,opt,name=type,proto3" json:"type" form:"type" query:"type"`
}

func (x *UpdateCoreRequest) Reset() {
	*x = UpdateCoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoreRequest) ProtoMessage() {}

func (x *UpdateCoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoreRequest.ProtoReflect.Descriptor instead.
func (*UpdateCoreRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCoreRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *UpdateCoreRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCoreRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCoreRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *UpdateCoreRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateCoreRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateCoreRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateCoreRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UpdateCoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *UpdateCoreResponse) Reset() {
	*x = UpdateCoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoreResponse) ProtoMessage() {}

func (x *UpdateCoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoreResponse.ProtoReflect.Descriptor instead.
func (*UpdateCoreResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCoreResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type DeleteCoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *DeleteCoreRequest) Reset() {
	*x = DeleteCoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoreRequest) ProtoMessage() {}

func (x *DeleteCoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoreRequest.ProtoReflect.Descriptor instead.
func (*DeleteCoreRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCoreRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *DeleteCoreRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *DeleteCoreResponse) Reset() {
	*x = DeleteCoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoreResponse) ProtoMessage() {}

func (x *DeleteCoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoreResponse.ProtoReflect.Descriptor instead.
func (*DeleteCoreResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_core_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCoreResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_freezonex_openiiot_api_core_proto protoreflect.FileDescriptor

var file_freezonex_openiiot_api_core_proto_rawDesc = []byte{
	0x0a, 0x21, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x04, 0x43,
	0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb1, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8e, 0x02, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5d, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x35, 0x5a, 0x33, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_freezonex_openiiot_api_core_proto_rawDescOnce sync.Once
	file_freezonex_openiiot_api_core_proto_rawDescData = file_freezonex_openiiot_api_core_proto_rawDesc
)

func file_freezonex_openiiot_api_core_proto_rawDescGZIP() []byte {
	file_freezonex_openiiot_api_core_proto_rawDescOnce.Do(func() {
		file_freezonex_openiiot_api_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_openiiot_api_core_proto_rawDescData)
	})
	return file_freezonex_openiiot_api_core_proto_rawDescData
}

var file_freezonex_openiiot_api_core_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_freezonex_openiiot_api_core_proto_goTypes = []interface{}{
	(*Core)(nil),                   // 0: freezonex.openiiot_api.Core
	(*AddCoreRequest)(nil),         // 1: freezonex.openiiot_api.AddCoreRequest
	(*AddCoreResponse)(nil),        // 2: freezonex.openiiot_api.AddCoreResponse
	(*GetCoreRequest)(nil),         // 3: freezonex.openiiot_api.GetCoreRequest
	(*GetCoreResponse)(nil),        // 4: freezonex.openiiot_api.GetCoreResponse
	(*UpdateCoreRequest)(nil),      // 5: freezonex.openiiot_api.UpdateCoreRequest
	(*UpdateCoreResponse)(nil),     // 6: freezonex.openiiot_api.UpdateCoreResponse
	(*DeleteCoreRequest)(nil),      // 7: freezonex.openiiot_api.DeleteCoreRequest
	(*DeleteCoreResponse)(nil),     // 8: freezonex.openiiot_api.DeleteCoreResponse
	(*base_req.BaseRequest)(nil),   // 9: base_req.BaseRequest
	(*base_resp.BaseResponse)(nil), // 10: base_resp.BaseResponse
}
var file_freezonex_openiiot_api_core_proto_depIdxs = []int32{
	9,  // 0: freezonex.openiiot_api.AddCoreRequest.base_request:type_name -> base_req.BaseRequest
	10, // 1: freezonex.openiiot_api.AddCoreResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // 2: freezonex.openiiot_api.GetCoreRequest.base_request:type_name -> base_req.BaseRequest
	10, // 3: freezonex.openiiot_api.GetCoreResponse.base_resp:type_name -> base_resp.BaseResponse
	0,  // 4: freezonex.openiiot_api.GetCoreResponse.data:type_name -> freezonex.openiiot_api.Core
	9,  // 5: freezonex.openiiot_api.UpdateCoreRequest.base_request:type_name -> base_req.BaseRequest
	10, // 6: freezonex.openiiot_api.UpdateCoreResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // 7: freezonex.openiiot_api.DeleteCoreRequest.base_request:type_name -> base_req.BaseRequest
	10, // 8: freezonex.openiiot_api.DeleteCoreResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_freezonex_openiiot_api_core_proto_init() }
func file_freezonex_openiiot_api_core_proto_init() {
	if File_freezonex_openiiot_api_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_openiiot_api_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Core); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCoreRequest); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCoreResponse); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoreRequest); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoreResponse); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoreRequest); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoreResponse); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCoreRequest); i {
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
		file_freezonex_openiiot_api_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCoreResponse); i {
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
			RawDescriptor: file_freezonex_openiiot_api_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_openiiot_api_core_proto_goTypes,
		DependencyIndexes: file_freezonex_openiiot_api_core_proto_depIdxs,
		MessageInfos:      file_freezonex_openiiot_api_core_proto_msgTypes,
	}.Build()
	File_freezonex_openiiot_api_core_proto = out.File
	file_freezonex_openiiot_api_core_proto_rawDesc = nil
	file_freezonex_openiiot_api_core_proto_goTypes = nil
	file_freezonex_openiiot_api_core_proto_depIdxs = nil
}