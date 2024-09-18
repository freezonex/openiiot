// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.4
// source: freezonex/openiiot_api/tdengine.proto

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

type TDEngineDSN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host" form:"host" query:"host"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username" form:"username" query:"username"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password" form:"password" query:"password"`
}

func (x *TDEngineDSN) Reset() {
	*x = TDEngineDSN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineDSN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineDSN) ProtoMessage() {}

func (x *TDEngineDSN) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineDSN.ProtoReflect.Descriptor instead.
func (*TDEngineDSN) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{0}
}

func (x *TDEngineDSN) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TDEngineDSN) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *TDEngineDSN) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TDEngineDSN) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type TDEngineRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  string  `protobuf:"bytes,1,opt,name=time,proto3" json:"time" form:"time" query:"time"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value" form:"value" query:"value"`
}

func (x *TDEngineRow) Reset() {
	*x = TDEngineRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineRow) ProtoMessage() {}

func (x *TDEngineRow) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineRow.ProtoReflect.Descriptor instead.
func (*TDEngineRow) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{1}
}

func (x *TDEngineRow) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *TDEngineRow) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TDEngineTestConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Dsn         *TDEngineDSN          `protobuf:"bytes,2,opt,name=dsn,proto3" json:"dsn" form:"dsn" query:"dsn"`
}

func (x *TDEngineTestConnectionRequest) Reset() {
	*x = TDEngineTestConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineTestConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineTestConnectionRequest) ProtoMessage() {}

func (x *TDEngineTestConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineTestConnectionRequest.ProtoReflect.Descriptor instead.
func (*TDEngineTestConnectionRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{2}
}

func (x *TDEngineTestConnectionRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *TDEngineTestConnectionRequest) GetDsn() *TDEngineDSN {
	if x != nil {
		return x.Dsn
	}
	return nil
}

type TDEngineTestConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp   *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	Successful bool                    `protobuf:"varint,2,opt,name=successful,proto3" json:"successful" form:"successful" query:"successful"`
}

func (x *TDEngineTestConnectionResponse) Reset() {
	*x = TDEngineTestConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineTestConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineTestConnectionResponse) ProtoMessage() {}

func (x *TDEngineTestConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineTestConnectionResponse.ProtoReflect.Descriptor instead.
func (*TDEngineTestConnectionResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{3}
}

func (x *TDEngineTestConnectionResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *TDEngineTestConnectionResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

type TDEngineExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Dsn         *TDEngineDSN          `protobuf:"bytes,2,opt,name=dsn,proto3" json:"dsn" form:"dsn" query:"dsn"`
	Sql         string                `protobuf:"bytes,3,opt,name=sql,proto3" json:"sql" form:"sql" query:"sql"`
}

func (x *TDEngineExecRequest) Reset() {
	*x = TDEngineExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineExecRequest) ProtoMessage() {}

func (x *TDEngineExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineExecRequest.ProtoReflect.Descriptor instead.
func (*TDEngineExecRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{4}
}

func (x *TDEngineExecRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *TDEngineExecRequest) GetDsn() *TDEngineDSN {
	if x != nil {
		return x.Dsn
	}
	return nil
}

func (x *TDEngineExecRequest) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

type TDEngineExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp     *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	RowsAffected int64                   `protobuf:"varint,2,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected" form:"rows_affected" query:"rows_affected"`
}

func (x *TDEngineExecResponse) Reset() {
	*x = TDEngineExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineExecResponse) ProtoMessage() {}

func (x *TDEngineExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineExecResponse.ProtoReflect.Descriptor instead.
func (*TDEngineExecResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{5}
}

func (x *TDEngineExecResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *TDEngineExecResponse) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type TDEngineBatchExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Dsn         *TDEngineDSN          `protobuf:"bytes,2,opt,name=dsn,proto3" json:"dsn" form:"dsn" query:"dsn"`
	Sql         string                `protobuf:"bytes,3,opt,name=sql,proto3" json:"sql" form:"sql" query:"sql"` // sql statement split by ";"
}

func (x *TDEngineBatchExecRequest) Reset() {
	*x = TDEngineBatchExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineBatchExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineBatchExecRequest) ProtoMessage() {}

func (x *TDEngineBatchExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineBatchExecRequest.ProtoReflect.Descriptor instead.
func (*TDEngineBatchExecRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{6}
}

func (x *TDEngineBatchExecRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *TDEngineBatchExecRequest) GetDsn() *TDEngineDSN {
	if x != nil {
		return x.Dsn
	}
	return nil
}

func (x *TDEngineBatchExecRequest) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

type TDEngineBatchExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp     *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	RowsAffected int64                   `protobuf:"varint,2,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected" form:"rows_affected" query:"rows_affected"`
}

func (x *TDEngineBatchExecResponse) Reset() {
	*x = TDEngineBatchExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineBatchExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineBatchExecResponse) ProtoMessage() {}

func (x *TDEngineBatchExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineBatchExecResponse.ProtoReflect.Descriptor instead.
func (*TDEngineBatchExecResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{7}
}

func (x *TDEngineBatchExecResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *TDEngineBatchExecResponse) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type TDEngineQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *base_req.BaseRequest `protobuf:"bytes,1,opt,name=base_request,json=baseRequest,proto3" json:"base_request" form:"base_request" query:"base_request"`
	Dsn         *TDEngineDSN          `protobuf:"bytes,2,opt,name=dsn,proto3" json:"dsn" form:"dsn" query:"dsn"`
	Sql         string                `protobuf:"bytes,3,opt,name=sql,proto3" json:"sql" form:"sql" query:"sql"`
}

func (x *TDEngineQueryRequest) Reset() {
	*x = TDEngineQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineQueryRequest) ProtoMessage() {}

func (x *TDEngineQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineQueryRequest.ProtoReflect.Descriptor instead.
func (*TDEngineQueryRequest) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{8}
}

func (x *TDEngineQueryRequest) GetBaseRequest() *base_req.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *TDEngineQueryRequest) GetDsn() *TDEngineDSN {
	if x != nil {
		return x.Dsn
	}
	return nil
}

func (x *TDEngineQueryRequest) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

type TDEngineQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base_resp.BaseResponse `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp" form:"base_resp" query:"base_resp"`
	// repeated TDEngineRow data = 2;
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data" form:"data" query:"data"`
}

func (x *TDEngineQueryResponse) Reset() {
	*x = TDEngineQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDEngineQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDEngineQueryResponse) ProtoMessage() {}

func (x *TDEngineQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_freezonex_openiiot_api_tdengine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDEngineQueryResponse.ProtoReflect.Descriptor instead.
func (*TDEngineQueryResponse) Descriptor() ([]byte, []int) {
	return file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP(), []int{9}
}

func (x *TDEngineQueryResponse) GetBaseResp() *base_resp.BaseResponse {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *TDEngineQueryResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_freezonex_openiiot_api_tdengine_proto protoreflect.FileDescriptor

var file_freezonex_openiiot_api_tdengine_proto_rawDesc = []byte{
	0x0a, 0x25, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x64, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a,
	0x1a, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66,
	0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d,
	0x0a, 0x0b, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x44, 0x53, 0x4e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x37, 0x0a,
	0x0b, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x1d, 0x54, 0x44, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x44, 0x53, 0x4e, 0x52, 0x03, 0x64, 0x73, 0x6e, 0x22, 0x76, 0x0a, 0x1e, 0x54, 0x44, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75,
	0x6c, 0x22, 0x98, 0x01, 0x0a, 0x13, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45, 0x78,
	0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x44, 0x53, 0x4e, 0x52, 0x03, 0x64, 0x73, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x22, 0x71, 0x0a, 0x14,
	0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f,
	0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22,
	0x9d, 0x01, 0x0a, 0x18, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x44, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x44, 0x53, 0x4e, 0x52, 0x03, 0x64, 0x73, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x71, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x22,
	0x76, 0x0a, 0x19, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x54, 0x44, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x64, 0x73,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x78, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x44, 0x53, 0x4e, 0x52, 0x03, 0x64, 0x73,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x71, 0x6c, 0x22, 0x61, 0x0a, 0x15, 0x54, 0x44, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x35, 0x5a, 0x33, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x78, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x2f, 0x62, 0x69, 0x7a,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x78,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freezonex_openiiot_api_tdengine_proto_rawDescOnce sync.Once
	file_freezonex_openiiot_api_tdengine_proto_rawDescData = file_freezonex_openiiot_api_tdengine_proto_rawDesc
)

func file_freezonex_openiiot_api_tdengine_proto_rawDescGZIP() []byte {
	file_freezonex_openiiot_api_tdengine_proto_rawDescOnce.Do(func() {
		file_freezonex_openiiot_api_tdengine_proto_rawDescData = protoimpl.X.CompressGZIP(file_freezonex_openiiot_api_tdengine_proto_rawDescData)
	})
	return file_freezonex_openiiot_api_tdengine_proto_rawDescData
}

var file_freezonex_openiiot_api_tdengine_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_freezonex_openiiot_api_tdengine_proto_goTypes = []interface{}{
	(*TDEngineDSN)(nil),                    // 0: freezonex.openiiot_api.TDEngineDSN
	(*TDEngineRow)(nil),                    // 1: freezonex.openiiot_api.TDEngineRow
	(*TDEngineTestConnectionRequest)(nil),  // 2: freezonex.openiiot_api.TDEngineTestConnectionRequest
	(*TDEngineTestConnectionResponse)(nil), // 3: freezonex.openiiot_api.TDEngineTestConnectionResponse
	(*TDEngineExecRequest)(nil),            // 4: freezonex.openiiot_api.TDEngineExecRequest
	(*TDEngineExecResponse)(nil),           // 5: freezonex.openiiot_api.TDEngineExecResponse
	(*TDEngineBatchExecRequest)(nil),       // 6: freezonex.openiiot_api.TDEngineBatchExecRequest
	(*TDEngineBatchExecResponse)(nil),      // 7: freezonex.openiiot_api.TDEngineBatchExecResponse
	(*TDEngineQueryRequest)(nil),           // 8: freezonex.openiiot_api.TDEngineQueryRequest
	(*TDEngineQueryResponse)(nil),          // 9: freezonex.openiiot_api.TDEngineQueryResponse
	(*base_req.BaseRequest)(nil),           // 10: base_req.BaseRequest
	(*base_resp.BaseResponse)(nil),         // 11: base_resp.BaseResponse
}
var file_freezonex_openiiot_api_tdengine_proto_depIdxs = []int32{
	10, // 0: freezonex.openiiot_api.TDEngineTestConnectionRequest.base_request:type_name -> base_req.BaseRequest
	0,  // 1: freezonex.openiiot_api.TDEngineTestConnectionRequest.dsn:type_name -> freezonex.openiiot_api.TDEngineDSN
	11, // 2: freezonex.openiiot_api.TDEngineTestConnectionResponse.base_resp:type_name -> base_resp.BaseResponse
	10, // 3: freezonex.openiiot_api.TDEngineExecRequest.base_request:type_name -> base_req.BaseRequest
	0,  // 4: freezonex.openiiot_api.TDEngineExecRequest.dsn:type_name -> freezonex.openiiot_api.TDEngineDSN
	11, // 5: freezonex.openiiot_api.TDEngineExecResponse.base_resp:type_name -> base_resp.BaseResponse
	10, // 6: freezonex.openiiot_api.TDEngineBatchExecRequest.base_request:type_name -> base_req.BaseRequest
	0,  // 7: freezonex.openiiot_api.TDEngineBatchExecRequest.dsn:type_name -> freezonex.openiiot_api.TDEngineDSN
	11, // 8: freezonex.openiiot_api.TDEngineBatchExecResponse.base_resp:type_name -> base_resp.BaseResponse
	10, // 9: freezonex.openiiot_api.TDEngineQueryRequest.base_request:type_name -> base_req.BaseRequest
	0,  // 10: freezonex.openiiot_api.TDEngineQueryRequest.dsn:type_name -> freezonex.openiiot_api.TDEngineDSN
	11, // 11: freezonex.openiiot_api.TDEngineQueryResponse.base_resp:type_name -> base_resp.BaseResponse
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_freezonex_openiiot_api_tdengine_proto_init() }
func file_freezonex_openiiot_api_tdengine_proto_init() {
	if File_freezonex_openiiot_api_tdengine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineDSN); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineRow); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineTestConnectionRequest); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineTestConnectionResponse); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineExecRequest); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineExecResponse); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineBatchExecRequest); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineBatchExecResponse); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineQueryRequest); i {
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
		file_freezonex_openiiot_api_tdengine_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDEngineQueryResponse); i {
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
			RawDescriptor: file_freezonex_openiiot_api_tdengine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freezonex_openiiot_api_tdengine_proto_goTypes,
		DependencyIndexes: file_freezonex_openiiot_api_tdengine_proto_depIdxs,
		MessageInfos:      file_freezonex_openiiot_api_tdengine_proto_msgTypes,
	}.Build()
	File_freezonex_openiiot_api_tdengine_proto = out.File
	file_freezonex_openiiot_api_tdengine_proto_rawDesc = nil
	file_freezonex_openiiot_api_tdengine_proto_goTypes = nil
	file_freezonex_openiiot_api_tdengine_proto_depIdxs = nil
}
