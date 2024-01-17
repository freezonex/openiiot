package application

import (
	"context"
	"freezonex/openiiot/biz/service/utils/common"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *AppService) AddApp(ctx context.Context, req *freezonex_openiiot_api.AddAppRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddAppResponse, error) {
	appID, err := a.AddAppDB(ctx, req.Name, req.Description, req.TenantId, req.Url, req.Username, req.Password, req.Type)
	if err != nil {
		logs.Error(ctx, "event=AddApp error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddAppResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = appID

	return resp, nil
}

// GetApp will get application record in condition
func (a *AppService) GetApp(ctx context.Context, req *freezonex_openiiot_api.GetAppRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetAppResponse, error) {
	apps, err := a.GetAppDB(ctx, req.Id, req.Name, req.TenantId, req.Url, req.Type)

	if err != nil {
		logs.Error(ctx, "event=GetApp error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetAppResponse)
	data := make([]*freezonex_openiiot_api.App, 0)
	for _, v := range apps {
		data = append(data, &freezonex_openiiot_api.App{
			Id:          v.ID,
			Name:        v.Name,
			Description: *v.Description,
			Username:    *v.Username,
			Password:    *v.Password,
			TenantId:    v.TenantID,
			Url:         v.URL,
			Type:        v.Type,
			CreateTime:  common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:  common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateApp will update application record
func (a *AppService) UpdateApp(ctx context.Context, req *freezonex_openiiot_api.UpdateAppRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateAppResponse, error) {
	err := a.UpdateAppDB(ctx, req.Id, req.Name, req.Description, req.TenantId, req.Url, req.Username, req.Password, req.Type)
	if err != nil {
		logs.Error(ctx, "event=UpdateApp error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateAppResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteApp will delete application record
func (a *AppService) DeleteApp(ctx context.Context, req *freezonex_openiiot_api.DeleteAppRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteAppResponse, error) {

	err := a.DeleteAppDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteApp error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteAppResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
