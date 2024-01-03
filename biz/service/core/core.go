package core

import (
	"context"
	"freezonex/openiiot/biz/service/utils/common"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *CoreService) AddCore(ctx context.Context, req *freezonex_openiiot_api.AddCoreRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddCoreResponse, error) {
	coreID, err := a.AddCoreDB(ctx, req.Name, req.Description, req.TenantId, req.Url, req.Username, req.Password, req.Type)
	if err != nil {
		logs.Error(ctx, "event=AddCore error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddCoreResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = coreID

	return resp, nil
}

// GetCore will get core record in condition
func (a *CoreService) GetCore(ctx context.Context, req *freezonex_openiiot_api.GetCoreRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetCoreResponse, error) {
	cores, err := a.GetCoreDB(ctx, req.Id, req.Name, req.TenantId, req.Url, req.Type)

	if err != nil {
		logs.Error(ctx, "event=GetCore error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetCoreResponse)
	data := make([]*freezonex_openiiot_api.Core, 0)
	for _, v := range cores {
		data = append(data, &freezonex_openiiot_api.Core{
			Id:         v.ID,
			Username:   v.Name,
			TenantId:   v.TenantID,
			Url:        v.URL,
			Type:       v.Type,
			CreateTime: common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime: common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateCore will update core record
func (a *CoreService) UpdateCore(ctx context.Context, req *freezonex_openiiot_api.UpdateCoreRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateCoreResponse, error) {
	err := a.UpdateCoreDB(ctx, req.Id, req.Name, req.Description, req.TenantId, req.Url, req.Username, req.Password, req.Type)
	if err != nil {
		logs.Error(ctx, "event=UpdateCore error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateCoreResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteCore will delete core record
func (a *CoreService) DeleteCore(ctx context.Context, req *freezonex_openiiot_api.DeleteCoreRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteCoreResponse, error) {

	err := a.DeleteCoreDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteCore error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteCoreResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
