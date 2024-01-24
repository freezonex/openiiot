// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: freezonex/openiiot_api/tenant.proto

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

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	IsDefault   bool   `protobuf:"varint,4,opt,name=is_default,json=isDefault,proto3" json:"is_default" form:"is_default" query:"is_default"`
	CreateTime  string `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time" form:"create_time" query:"create_time"`
	UpdateTime  string `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time" form:"update_time" query:"update_time"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tenant) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *Tenant) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Tenant) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AddTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Name        string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	IsDefault   bool                  `protobuf:"varint,4,opt,name=is_default,json=isDefault,proto3" json:"is_default" form:"is_default" query:"is_default"`
}

func (x *AddTenantRequest) Reset() {
	*x = AddTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTenantRequest) ProtoMessage() {}

func (x *AddTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTenantRequest.ProtoReflect.Descriptor instead.
func (*AddTenantRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *AddTenantRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *AddTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTenantRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddTenantRequest) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

type AddTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Id       string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *AddTenantResponse) Reset() {
	*x = AddTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTenantResponse) ProtoMessage() {}

func (x *AddTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTenantResponse.ProtoReflect.Descriptor instead.
func (*AddTenantResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *AddTenantResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *AddTenantResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *GetTenantRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Data     []*Tenant               `protobuf:"bytes,2,rep,name=data,proto3" json:"data" form:"data" query:"data"`
}

func (x *GetTenantResponse) Reset() {
	*x = GetTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantResponse) ProtoMessage() {}

func (x *GetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantResponse.ProtoReflect.Descriptor instead.
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *GetTenantResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetTenantResponse) GetData() []*Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description" query:"description"`
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTenantRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *UpdateTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTenantRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *UpdateTenantResponse) Reset() {
	*x = UpdateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantResponse) ProtoMessage() {}

func (x *UpdateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantResponse.ProtoReflect.Descriptor instead.
func (*UpdateTenantResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTenantResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type DeleteTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTenantRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *DeleteTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *DeleteTenantResponse) Reset() {
	*x = DeleteTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantResponse) ProtoMessage() {}

func (x *DeleteTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantResponse.ProtoReflect.Descriptor instead.
func (*DeleteTenantResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTenantResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type GetAllTenantNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
}

func (x *GetAllTenantNameRequest) Reset() {
	*x = GetAllTenantNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTenantNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTenantNameRequest) ProtoMessage() {}

func (x *GetAllTenantNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTenantNameRequest.ProtoReflect.Descriptor instead.
func (*GetAllTenantNameRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllTenantNameRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

type GetAllTenantNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp    *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	TenantNames []string                `protobuf:"bytes,2,rep,name=tenant_names,json=tenantNames,proto3" json:"tenant_names" form:"tenant_names" query:"tenant_names"`
}

func (x *GetAllTenantNameResponse) Reset() {
	*x = GetAllTenantNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTenantNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTenantNameResponse) ProtoMessage() {}

func (x *GetAllTenantNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tenant_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTenantNameResponse.ProtoReflect.Descriptor instead.
func (*GetAllTenantNameResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tenant_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllTenantNameResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetAllTenantNameResponse) GetTenantNames() []string {
	if x != nil {
		return x.TenantNames
	}
	return nil
}

var File_freezonex_openiiot_api_tenant_proto protoreflect.FileDescriptor

var file_freezonex_openiiot_api_tenant_proto_rawDesc = []byte{
	0x0a, 0x23, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x66,
	0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a,
	0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa1,
	0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x70, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x7d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95,
	0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x5f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x53, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x73, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x35, 0x5a,
	0x33, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x69, 0x6f, 0x74, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x72,
	0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freezonex_openiiot_api_tenant_proto_rawDescOnce sync.Once
	file_freezonex_openiiot_api_tenant_proto_rawDescData = file_freezonex_openiiot_api_tenant_proto_rawDesc
)

func file_freezonex_openiiot_api_tenant_proto_rawDescGZIP() []byte {
	file_freezonex_openiiot_api_tenant_proto_rawDescOnce.Do(func() {
		file_freezonex_openiiot_api_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_openiiot_api_tenant_proto_rawDescData)
	})
	return file_freezonex_openiiot_api_tenant_proto_rawDescData
}

var file_freezonex_openiiot_api_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_freezonex_openiiot_api_tenant_proto_goTypes = []interface{}{
	(*Tenant)(nil),                   // 0: freezonex.openiiot_api.Tenant
	(*AddTenantRequest)(nil),         // 1: freezonex.openiiot_api.AddTenantRequest
	(*AddTenantResponse)(nil),        // 2: freezonex.openiiot_api.AddTenantResponse
	(*GetTenantRequest)(nil),         // 3: freezonex.openiiot_api.GetTenantRequest
	(*GetTenantResponse)(nil),        // 4: freezonex.openiiot_api.GetTenantResponse
	(*UpdateTenantRequest)(nil),      // 5: freezonex.openiiot_api.UpdateTenantRequest
	(*UpdateTenantResponse)(nil),     // 6: freezonex.openiiot_api.UpdateTenantResponse
	(*DeleteTenantRequest)(nil),      // 7: freezonex.openiiot_api.DeleteTenantRequest
	(*DeleteTenantResponse)(nil),     // 8: freezonex.openiiot_api.DeleteTenantResponse
	(*GetAllTenantNameRequest)(nil),  // 9: freezonex.openiiot_api.GetAllTenantNameRequest
	(*GetAllTenantNameResponse)(nil), // 10: freezonex.openiiot_api.GetAllTenantNameResponse
	(*base_req.BaseRequest)(nil),     // 11: base_req.BaseRequest
	(*base_resp.BaseResponse)(nil),   // 12: base_resp.BaseResponse
}
var file_freezonex_openiiot_api_tenant_proto_depIdxs = []int32{
	11, // 0: freezonex.openiiot_api.AddTenantRequest.base_request:type_name -> base_req.BaseRequest
	12, // 1: freezonex.openiiot_api.AddTenantResponse.base_resp:type_name -> base_resp.BaseResponse
	11, // 2: freezonex.openiiot_api.GetTenantRequest.base_request:type_name -> base_req.BaseRequest
	12, // 3: freezonex.openiiot_api.GetTenantResponse.base_resp:type_name -> base_resp.BaseResponse
	0,  // 4: freezonex.openiiot_api.GetTenantResponse.data:type_name -> freezonex.openiiot_api.Tenant
	11, // 5: freezonex.openiiot_api.UpdateTenantRequest.base_request:type_name -> base_req.BaseRequest
	12, // 6: freezonex.openiiot_api.UpdateTenantResponse.base_resp:type_name -> base_resp.BaseResponse
	11, // 7: freezonex.openiiot_api.DeleteTenantRequest.base_request:type_name -> base_req.BaseRequest
	12, // 8: freezonex.openiiot_api.DeleteTenantResponse.base_resp:type_name -> base_resp.BaseResponse
	11, // 9: freezonex.openiiot_api.GetAllTenantNameRequest.base_request:type_name -> base_req.BaseRequest
	12, // 10: freezonex.openiiot_api.GetAllTenantNameResponse.base_resp:type_name -> base_resp.BaseResponse
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_freezonex_openiiot_api_tenant_proto_init() }
func file_freezonex_openiiot_api_tenant_proto_init() {
	if File_freezonex_openiiot_api_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_openiiot_api_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTenantRequest); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTenantResponse); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantRequest); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantResponse); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantRequest); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantResponse); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTenantRequest); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTenantResponse); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTenantNameRequest); i {
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
		file_freezonex_openiiot_api_tenant_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTenantNameResponse); i {
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
			RawDescriptor: file_freezonex_openiiot_api_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_openiiot_api_tenant_proto_goTypes,
		DependencyIndexes: file_freezonex_openiiot_api_tenant_proto_depIdxs,
		MessageInfos:      file_freezonex_openiiot_api_tenant_proto_msgTypes,
	}.Build()
	File_freezonex_openiiot_api_tenant_proto = out.File
	file_freezonex_openiiot_api_tenant_proto_rawDesc = nil
	file_freezonex_openiiot_api_tenant_proto_goTypes = nil
	file_freezonex_openiiot_api_tenant_proto_depIdxs = nil
}
