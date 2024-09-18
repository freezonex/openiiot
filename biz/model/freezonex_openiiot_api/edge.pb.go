// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.4
// source: freezonex/openiiot_api/edge.proto

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

type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    string `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Username    string `protobuf:"bytes,6,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password    string `protobuf:"bytes,7,opt,name=password,proto3" json:"password" form:"password" query:"password"`
	Type        string `protobuf:"bytes,8,opt,name=type,proto3" json:"type" form:"type" query:"type"`
	CreateTime  string `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time" form:"create_time" query:"create_time"`
	UpdateTime  string `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time" form:"update_time" query:"update_time"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{0}
}

func (x *Edge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Edge) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Edge) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Edge) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *Edge) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Edge) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Edge) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Edge) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Edge) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Edge) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AddEdgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Name        string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    string                `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string                `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Username    string                `protobuf:"bytes,6,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password    string                `protobuf:"bytes,7,opt,name=password,proto3" json:"password" form:"password" query:"password"`
	Type        string                `protobuf:"bytes,8,opt,name=type,proto3" json:"type" form:"type" query:"type"`
}

func (x *AddEdgeRequest) Reset() {
	*x = AddEdgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEdgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEdgeRequest) ProtoMessage() {}

func (x *AddEdgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEdgeRequest.ProtoReflect.Descriptor instead.
func (*AddEdgeRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{1}
}

func (x *AddEdgeRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *AddEdgeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddEdgeRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddEdgeRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *AddEdgeRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddEdgeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddEdgeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddEdgeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AddEdgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Id       string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *AddEdgeResponse) Reset() {
	*x = AddEdgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEdgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEdgeResponse) ProtoMessage() {}

func (x *AddEdgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEdgeResponse.ProtoReflect.Descriptor instead.
func (*AddEdgeResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{2}
}

func (x *AddEdgeResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *AddEdgeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEdgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	TenantId    string                `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string                `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Type        string                `protobuf:"bytes,6,opt,name=type,proto3" json:"type" form:"type" query:"type"`
}

func (x *GetEdgeRequest) Reset() {
	*x = GetEdgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEdgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEdgeRequest) ProtoMessage() {}

func (x *GetEdgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEdgeRequest.ProtoReflect.Descriptor instead.
func (*GetEdgeRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{3}
}

func (x *GetEdgeRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetEdgeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetEdgeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEdgeRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *GetEdgeRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetEdgeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetEdgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Data     []*Edge                 `protobuf:"bytes,2,rep,name=data,proto3" json:"data" form:"data" query:"data"`
}

func (x *GetEdgeResponse) Reset() {
	*x = GetEdgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEdgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEdgeResponse) ProtoMessage() {}

func (x *GetEdgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEdgeResponse.ProtoReflect.Descriptor instead.
func (*GetEdgeResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{4}
}

func (x *GetEdgeResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetEdgeResponse) GetData() []*Edge {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateEdgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    string                `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	Url         string                `protobuf:"bytes,6,opt,name=url,proto3" json:"url" form:"url" query:"url"`
	Username    string                `protobuf:"bytes,7,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password    string                `protobuf:"bytes,8,opt,name=password,proto3" json:"password" form:"password" query:"password"`
	Type        string                `protobuf:"bytes,9,opt,name=type,proto3" json:"type" form:"type" query:"type"`
}

func (x *UpdateEdgeRequest) Reset() {
	*x = UpdateEdgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEdgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEdgeRequest) ProtoMessage() {}

func (x *UpdateEdgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEdgeRequest.ProtoReflect.Descriptor instead.
func (*UpdateEdgeRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEdgeRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *UpdateEdgeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEdgeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEdgeRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEdgeRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *UpdateEdgeRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateEdgeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateEdgeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateEdgeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UpdateEdgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *UpdateEdgeResponse) Reset() {
	*x = UpdateEdgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEdgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEdgeResponse) ProtoMessage() {}

func (x *UpdateEdgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEdgeResponse.ProtoReflect.Descriptor instead.
func (*UpdateEdgeResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEdgeResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type DeleteEdgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *DeleteEdgeRequest) Reset() {
	*x = DeleteEdgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEdgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEdgeRequest) ProtoMessage() {}

func (x *DeleteEdgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEdgeRequest.ProtoReflect.Descriptor instead.
func (*DeleteEdgeRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEdgeRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *DeleteEdgeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEdgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *DeleteEdgeResponse) Reset() {
	*x = DeleteEdgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEdgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEdgeResponse) ProtoMessage() {}

func (x *DeleteEdgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_edge_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEdgeResponse.ProtoReflect.Descriptor instead.
func (*DeleteEdgeResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_edge_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEdgeResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_freezonex_openiiot_api_edge_proto protoreflect.FileDescriptor

var file_freezonex_openiiot_api_edge_proto_rawDesc = []byte{
	0x0a, 0x21, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x04, 0x45,
	0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
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
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x45, 0x64,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x64, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb1, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8e, 0x02, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5d, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
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
	file_freezonex_openiiot_api_edge_proto_rawDescOnce sync.Once
	file_freezonex_openiiot_api_edge_proto_rawDescData = file_freezonex_openiiot_api_edge_proto_rawDesc
)

func file_freezonex_openiiot_api_edge_proto_rawDescGZIP() []byte {
	file_freezonex_openiiot_api_edge_proto_rawDescOnce.Do(func() {
		file_freezonex_openiiot_api_edge_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_openiiot_api_edge_proto_rawDescData)
	})
	return file_freezonex_openiiot_api_edge_proto_rawDescData
}

var file_freezonex_openiiot_api_edge_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_freezonex_openiiot_api_edge_proto_goTypes = []interface{}{
	(*Edge)(nil),                   // 0: freezonex.openiiot_api.Edge
	(*AddEdgeRequest)(nil),         // 1: freezonex.openiiot_api.AddEdgeRequest
	(*AddEdgeResponse)(nil),        // 2: freezonex.openiiot_api.AddEdgeResponse
	(*GetEdgeRequest)(nil),         // 3: freezonex.openiiot_api.GetEdgeRequest
	(*GetEdgeResponse)(nil),        // 4: freezonex.openiiot_api.GetEdgeResponse
	(*UpdateEdgeRequest)(nil),      // 5: freezonex.openiiot_api.UpdateEdgeRequest
	(*UpdateEdgeResponse)(nil),     // 6: freezonex.openiiot_api.UpdateEdgeResponse
	(*DeleteEdgeRequest)(nil),      // 7: freezonex.openiiot_api.DeleteEdgeRequest
	(*DeleteEdgeResponse)(nil),     // 8: freezonex.openiiot_api.DeleteEdgeResponse
	(*base_req.BaseRequest)(nil),   // 9: base_req.BaseRequest
	(*base_resp.BaseResponse)(nil), // 10: base_resp.BaseResponse
}
var file_freezonex_openiiot_api_edge_proto_depIdxs = []int32{
	9,  // 0: freezonex.openiiot_api.AddEdgeRequest.base_request:type_name -> base_req.BaseRequest
	10, // 1: freezonex.openiiot_api.AddEdgeResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // 2: freezonex.openiiot_api.GetEdgeRequest.base_request:type_name -> base_req.BaseRequest
	10, // 3: freezonex.openiiot_api.GetEdgeResponse.base_resp:type_name -> base_resp.BaseResponse
	0,  // 4: freezonex.openiiot_api.GetEdgeResponse.data:type_name -> freezonex.openiiot_api.Edge
	9,  // 5: freezonex.openiiot_api.UpdateEdgeRequest.base_request:type_name -> base_req.BaseRequest
	10, // 6: freezonex.openiiot_api.UpdateEdgeResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // 7: freezonex.openiiot_api.DeleteEdgeRequest.base_request:type_name -> base_req.BaseRequest
	10, // 8: freezonex.openiiot_api.DeleteEdgeResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_freezonex_openiiot_api_edge_proto_init() }
func file_freezonex_openiiot_api_edge_proto_init() {
	if File_freezonex_openiiot_api_edge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_openiiot_api_edge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEdgeRequest); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEdgeResponse); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEdgeRequest); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEdgeResponse); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEdgeRequest); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEdgeResponse); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEdgeRequest); i {
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
		file_freezonex_openiiot_api_edge_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEdgeResponse); i {
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
			RawDescriptor: file_freezonex_openiiot_api_edge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_openiiot_api_edge_proto_goTypes,
		DependencyIndexes: file_freezonex_openiiot_api_edge_proto_depIdxs,
		MessageInfos:      file_freezonex_openiiot_api_edge_proto_msgTypes,
	}.Build()
	File_freezonex_openiiot_api_edge_proto = out.File
	file_freezonex_openiiot_api_edge_proto_rawDesc = nil
	file_freezonex_openiiot_api_edge_proto_goTypes = nil
	file_freezonex_openiiot_api_edge_proto_depIdxs = nil
}
