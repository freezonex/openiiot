// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: freezonex/openiiot_api/flow.proto

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

type Flow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description    string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId       int64   `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	LastModifiedBy string  `protobuf:"bytes,5,opt,name=last_modified_by,json=lastModifiedBy,proto3" json:"last_modified_by" form:"last_modified_by" query:"last_modified_by"`
	FlowType       string  `protobuf:"bytes,6,opt,name=flow_type,json=flowType,proto3" json:"flow_type" form:"flow_type" query:"flow_type"`
	EdgeIds        []int64 `protobuf:"varint,7,rep,packed,name=edge_ids,json=edgeIds,proto3" json:"edge_ids" form:"edge_ids" query:"edge_ids"`
	CoreIds        []int64 `protobuf:"varint,8,rep,packed,name=core_ids,json=coreIds,proto3" json:"core_ids" form:"core_ids" query:"core_ids"`
	AppIds         []int64 `protobuf:"varint,9,rep,packed,name=app_ids,json=appIds,proto3" json:"app_ids" form:"app_ids" query:"app_ids"`
	CreateTime     string  `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time" form:"create_time" query:"create_time"`
	UpdateTime     string  `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time" form:"update_time" query:"update_time"`
}

func (x *Flow) Reset() {
	*x = Flow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flow) ProtoMessage() {}

func (x *Flow) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flow.ProtoReflect.Descriptor instead.
func (*Flow) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{0}
}

func (x *Flow) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Flow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flow) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Flow) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *Flow) GetLastModifiedBy() string {
	if x != nil {
		return x.LastModifiedBy
	}
	return ""
}

func (x *Flow) GetFlowType() string {
	if x != nil {
		return x.FlowType
	}
	return ""
}

func (x *Flow) GetEdgeIds() []int64 {
	if x != nil {
		return x.EdgeIds
	}
	return nil
}

func (x *Flow) GetCoreIds() []int64 {
	if x != nil {
		return x.CoreIds
	}
	return nil
}

func (x *Flow) GetAppIds() []int64 {
	if x != nil {
		return x.AppIds
	}
	return nil
}

func (x *Flow) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Flow) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AddFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Name        string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    int64                 `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	FlowType    string                `protobuf:"bytes,5,opt,name=flow_type,json=flowType,proto3" json:"flow_type" form:"flow_type" query:"flow_type"`
}

func (x *AddFlowRequest) Reset() {
	*x = AddFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFlowRequest) ProtoMessage() {}

func (x *AddFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFlowRequest.ProtoReflect.Descriptor instead.
func (*AddFlowRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{1}
}

func (x *AddFlowRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *AddFlowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddFlowRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddFlowRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *AddFlowRequest) GetFlowType() string {
	if x != nil {
		return x.FlowType
	}
	return ""
}

type AddFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Id       int64                   `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *AddFlowResponse) Reset() {
	*x = AddFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFlowResponse) ProtoMessage() {}

func (x *AddFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFlowResponse.ProtoReflect.Descriptor instead.
func (*AddFlowResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{2}
}

func (x *AddFlowResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *AddFlowResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest    *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id             int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name           string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	TenantId       int64                 `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	LastModifiedBy string                `protobuf:"bytes,5,opt,name=last_modified_by,json=lastModifiedBy,proto3" json:"last_modified_by" form:"last_modified_by" query:"last_modified_by"`
	FlowType       string                `protobuf:"bytes,6,opt,name=flow_type,json=flowType,proto3" json:"flow_type" form:"flow_type" query:"flow_type"`
}

func (x *GetFlowRequest) Reset() {
	*x = GetFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlowRequest) ProtoMessage() {}

func (x *GetFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlowRequest.ProtoReflect.Descriptor instead.
func (*GetFlowRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{3}
}

func (x *GetFlowRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetFlowRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetFlowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFlowRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *GetFlowRequest) GetLastModifiedBy() string {
	if x != nil {
		return x.LastModifiedBy
	}
	return ""
}

func (x *GetFlowRequest) GetFlowType() string {
	if x != nil {
		return x.FlowType
	}
	return ""
}

type GetFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Data     []*Flow                 `protobuf:"bytes,2,rep,name=data,proto3" json:"data" form:"data" query:"data"`
}

func (x *GetFlowResponse) Reset() {
	*x = GetFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlowResponse) ProtoMessage() {}

func (x *GetFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlowResponse.ProtoReflect.Descriptor instead.
func (*GetFlowResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{4}
}

func (x *GetFlowResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetFlowResponse) GetData() []*Flow {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string                `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	TenantId    int64                 `protobuf:"varint,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id" form:"tenant_id" query:"tenant_id"`
	FlowType    string                `protobuf:"bytes,6,opt,name=flow_type,json=flowType,proto3" json:"flow_type" form:"flow_type" query:"flow_type"`
	EdgeIds     []int64               `protobuf:"varint,7,rep,packed,name=edge_ids,json=edgeIds,proto3" json:"edge_ids" form:"edge_ids" query:"edge_ids"`
	CoreIds     []int64               `protobuf:"varint,8,rep,packed,name=core_ids,json=coreIds,proto3" json:"core_ids" form:"core_ids" query:"core_ids"`
	AppIds      []int64               `protobuf:"varint,9,rep,packed,name=app_ids,json=appIds,proto3" json:"app_ids" form:"app_ids" query:"app_ids"`
}

func (x *UpdateFlowRequest) Reset() {
	*x = UpdateFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlowRequest) ProtoMessage() {}

func (x *UpdateFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlowRequest.ProtoReflect.Descriptor instead.
func (*UpdateFlowRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFlowRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *UpdateFlowRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFlowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFlowRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateFlowRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *UpdateFlowRequest) GetFlowType() string {
	if x != nil {
		return x.FlowType
	}
	return ""
}

func (x *UpdateFlowRequest) GetEdgeIds() []int64 {
	if x != nil {
		return x.EdgeIds
	}
	return nil
}

func (x *UpdateFlowRequest) GetCoreIds() []int64 {
	if x != nil {
		return x.CoreIds
	}
	return nil
}

func (x *UpdateFlowRequest) GetAppIds() []int64 {
	if x != nil {
		return x.AppIds
	}
	return nil
}

type UpdateFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *UpdateFlowResponse) Reset() {
	*x = UpdateFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlowResponse) ProtoMessage() {}

func (x *UpdateFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlowResponse.ProtoReflect.Descriptor instead.
func (*UpdateFlowResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFlowResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type DeleteFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Id          int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id" form:"id" query:"id"`
}

func (x *DeleteFlowRequest) Reset() {
	*x = DeleteFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlowRequest) ProtoMessage() {}

func (x *DeleteFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlowRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlowRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFlowRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *DeleteFlowRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
}

func (x *DeleteFlowResponse) Reset() {
	*x = DeleteFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlowResponse) ProtoMessage() {}

func (x *DeleteFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_flow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlowResponse.ProtoReflect.Descriptor instead.
func (*DeleteFlowResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_flow_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFlowResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_freezonex_openiiot_api_flow_proto protoreflect.FileDescriptor

var file_freezonex_openiiot_api_flow_proto_rawDesc = []byte{
	0x0a, 0x21, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x04, 0x46,
	0x6c, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x64, 0x67, 0x65, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x61, 0x70, 0x70, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xba,
	0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x22, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x9c, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x65, 0x64, 0x67, 0x65, 0x49, 0x64, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x5d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x35, 0x5a, 0x33, 0x66, 0x72,
	0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74,
	0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freezonex_openiiot_api_flow_proto_rawDescOnce sync.Once
	file_freezonex_openiiot_api_flow_proto_rawDescData = file_freezonex_openiiot_api_flow_proto_rawDesc
)

func file_freezonex_openiiot_api_flow_proto_rawDescGZIP() []byte {
	file_freezonex_openiiot_api_flow_proto_rawDescOnce.Do(func() {
		file_freezonex_openiiot_api_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_openiiot_api_flow_proto_rawDescData)
	})
	return file_freezonex_openiiot_api_flow_proto_rawDescData
}

var file_freezonex_openiiot_api_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_freezonex_openiiot_api_flow_proto_goTypes = []interface{}{
	(*Flow)(nil),                   // 0: freezonex.openiiot_api.Flow
	(*AddFlowRequest)(nil),         // 1: freezonex.openiiot_api.AddFlowRequest
	(*AddFlowResponse)(nil),        // 2: freezonex.openiiot_api.AddFlowResponse
	(*GetFlowRequest)(nil),         // 3: freezonex.openiiot_api.GetFlowRequest
	(*GetFlowResponse)(nil),        // 4: freezonex.openiiot_api.GetFlowResponse
	(*UpdateFlowRequest)(nil),      // 5: freezonex.openiiot_api.UpdateFlowRequest
	(*UpdateFlowResponse)(nil),     // 6: freezonex.openiiot_api.UpdateFlowResponse
	(*DeleteFlowRequest)(nil),      // 7: freezonex.openiiot_api.DeleteFlowRequest
	(*DeleteFlowResponse)(nil),     // 8: freezonex.openiiot_api.DeleteFlowResponse
	(*base_req.BaseRequest)(nil),   // 9: base_req.BaseRequest
	(*base_resp.BaseResponse)(nil), // 10: base_resp.BaseResponse
}
var file_freezonex_openiiot_api_flow_proto_depIdxs = []int32{
	9,  // 0: freezonex.openiiot_api.AddFlowRequest.base_request:type_name -> base_req.BaseRequest
	10, // 1: freezonex.openiiot_api.AddFlowResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // 2: freezonex.openiiot_api.GetFlowRequest.base_request:type_name -> base_req.BaseRequest
	10, // 3: freezonex.openiiot_api.GetFlowResponse.base_resp:type_name -> base_resp.BaseResponse
	0,  // 4: freezonex.openiiot_api.GetFlowResponse.data:type_name -> freezonex.openiiot_api.Flow
	9,  // 5: freezonex.openiiot_api.UpdateFlowRequest.base_request:type_name -> base_req.BaseRequest
	10, // 6: freezonex.openiiot_api.UpdateFlowResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // 7: freezonex.openiiot_api.DeleteFlowRequest.base_request:type_name -> base_req.BaseRequest
	10, // 8: freezonex.openiiot_api.DeleteFlowResponse.base_resp:type_name -> base_resp.BaseResponse
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_freezonex_openiiot_api_flow_proto_init() }
func file_freezonex_openiiot_api_flow_proto_init() {
	if File_freezonex_openiiot_api_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_openiiot_api_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flow); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFlowRequest); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFlowResponse); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlowRequest); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlowResponse); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlowRequest); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlowResponse); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlowRequest); i {
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
		file_freezonex_openiiot_api_flow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlowResponse); i {
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
			RawDescriptor: file_freezonex_openiiot_api_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_openiiot_api_flow_proto_goTypes,
		DependencyIndexes: file_freezonex_openiiot_api_flow_proto_depIdxs,
		MessageInfos:      file_freezonex_openiiot_api_flow_proto_msgTypes,
	}.Build()
	File_freezonex_openiiot_api_flow_proto = out.File
	file_freezonex_openiiot_api_flow_proto_rawDesc = nil
	file_freezonex_openiiot_api_flow_proto_goTypes = nil
	file_freezonex_openiiot_api_flow_proto_depIdxs = nil
}
