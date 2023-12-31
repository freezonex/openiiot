package edge

import (
	"context"
	"freezonex/openiiot/biz/service/utils/common"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *EdgeService) AddEdge(ctx context.Context, req *freezonex_openiiot_api.AddEdgeRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddEdgeResponse, error) {
	edgeID, err := a.AddEdgeDB(ctx, req.Name, req.Description, req.TenantId, req.Url, req.Username, req.Password, req.Type)
	if err != nil {
		logs.Error(ctx, "event=AddEdge error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddEdgeResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = edgeID

	return resp, nil
}

// GetEdge will get edge record in condition
func (a *EdgeService) GetEdge(ctx context.Context, req *freezonex_openiiot_api.GetEdgeRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetEdgeResponse, error) {
	edges, err := a.GetEdgeDB(ctx, req.Id, req.Name, req.TenantId, req.Url, req.Type)

	if err != nil {
		logs.Error(ctx, "event=GetEdge error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetEdgeResponse)
	data := make([]*freezonex_openiiot_api.Edge, 0)
	for _, v := range edges {
		data = append(data, &freezonex_openiiot_api.Edge{
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

// UpdateEdge will update edge record
func (a *EdgeService) UpdateEdge(ctx context.Context, req *freezonex_openiiot_api.UpdateEdgeRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateEdgeResponse, error) {
	err := a.UpdateEdgeDB(ctx, req.Id, req.Name, req.Description, req.TenantId, req.Url, req.Username, req.Password, req.Type)
	if err != nil {
		logs.Error(ctx, "event=UpdateEdge error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateEdgeResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteEdge will delete edge record
func (a *EdgeService) DeleteEdge(ctx context.Context, req *freezonex_openiiot_api.DeleteEdgeRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteEdgeResponse, error) {

	err := a.DeleteEdgeDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteEdge error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteEdgeResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
