package tenant

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
)

func (a *TenantService) AddTenant(ctx context.Context, req *freezonex_openiiot_api.AddTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddTenantResponse, error) {
	tenantID, err := a.AddTenantDB(ctx, req.Name, req.Description, req.IsDefault)
	if err != nil {
		logs.Error(ctx, "event=AddTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddTenantResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(tenantID)

	return resp, nil
}

// GetTenant will get tenant record in condition
func (a *TenantService) GetTenant(ctx context.Context, req *freezonex_openiiot_api.GetTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetTenantResponse, error) {
	tenants, err := a.GetTenantDB(ctx, common.StringToInt64(req.Id), req.Name)

	if err != nil {
		logs.Error(ctx, "event=GetTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetTenantResponse)
	data := make([]*freezonex_openiiot_api.Tenant, 0)
	for _, v := range tenants {
		data = append(data, &freezonex_openiiot_api.Tenant{
			Id:          common.Int64ToString(v.ID),
			Name:        v.Name,
			Description: *v.Description,
			IsDefault:   common.ProtectNullBoolPointer(v.IsDefault),
			CreateTime:  common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:  common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateTenant will update tenant record
func (a *TenantService) UpdateTenant(ctx context.Context, req *freezonex_openiiot_api.UpdateTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateTenantResponse, error) {
	err := a.UpdateTenantDB(ctx, common.StringToInt64(req.Id), req.Name, req.Description)
	if err != nil {
		logs.Error(ctx, "event=UpdateTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateTenantResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteTenant will delete tenant record
func (a *TenantService) DeleteTenant(ctx context.Context, req *freezonex_openiiot_api.DeleteTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteTenantResponse, error) {
	//Delete tenant also should delete tenant user, edge pool, core pool, application pool, flow
	/*err := a.DeleteTenantUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteTenant user error=%v", err.Error())
		return nil, err
	}*/

	// Delete tenant
	err := a.DeleteTenantDB(ctx, common.StringToInt64(req.Id))
	if err != nil {
		logs.Error(ctx, "event=DeleteTenant error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteTenantResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *TenantService) GetAllTenantName(ctx context.Context, req *freezonex_openiiot_api.GetAllTenantNameRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetAllTenantNameResponse, error) {
	tenants, err := a.GetTenantDB(ctx, 0, "")

	if err != nil {
		logs.Error(ctx, "event=GetTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetAllTenantNameResponse)
	tenantNames := make([]string, 0) // Change to slice of strings
	for _, v := range tenants {
		tenantNames = append(tenantNames, v.Name) // Append only the name
	}

	resp.TenantNames = tenantNames
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
