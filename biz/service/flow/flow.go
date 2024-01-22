package flow

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/service/utils/common"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *FlowService) AddFlow(ctx context.Context, req *freezonex_openiiot_api.AddFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddFlowResponse, error) {
	flowID, err := a.AddFlowDB(ctx, req.Name, req.Description, req.TenantId, req.FlowType)
	if err != nil {
		logs.Error(ctx, "event=AddFlow error=%v", err.Error())
		return nil, err
	}

	if _, err := a.AddFlowEdge(ctx, flowID, req.EdgeIds); err != nil {
		return nil, err
	}
	if _, err := a.AddFlowCore(ctx, flowID, req.CoreIds); err != nil {
		return nil, err
	}
	if _, err := a.AddFlowApp(ctx, flowID, req.AppIds); err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = flowID

	return resp, nil
}

// GetFlow will get flow record in condition
func (a *FlowService) GetFlow(ctx context.Context, req *freezonex_openiiot_api.GetFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetFlowResponse, error) {
	// Fetch flows from the database
	flows, err := a.GetFlowDB(ctx, req.Id, req.Name, req.TenantId, req.LastModifiedBy, req.FlowType)
	if err != nil {
		logs.Error(ctx, "event=GetFlow error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.GetFlowResponse)
	// Prepare the response data
	data := make([]*freezonex_openiiot_api.Flow, 0)
	for _, v := range flows {
		// For each flow, fetch EdgeIds, CoreIds, and AppIds from their respective tables
		edgeIds, err := a.GetFlowEdgeIDs(ctx, v.ID)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		coreIds, err := a.GetFlowCoreIDs(ctx, v.ID)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		appIds, err := a.GetFlowAppIDs(ctx, v.ID)
		if err != nil {
			return nil, err // Handle error appropriately
		}

		data = append(data, &freezonex_openiiot_api.Flow{
			Id:             v.ID,
			Name:           v.Name,
			Description:    *v.Description,
			TenantId:       v.TenantID,
			LastModifiedBy: *v.LastModifiedBy,
			FlowType:       *v.FlowType,
			EdgeIds:        edgeIds,
			CoreIds:        coreIds,
			AppIds:         appIds,
			CreateTime:     common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:     common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}

	// Prepare and return the final response
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Data = data

	return resp, nil
}

// UpdateFlow will update flow record
func (a *FlowService) UpdateFlow(ctx context.Context, req *freezonex_openiiot_api.UpdateFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateFlowResponse, error) {
	err := a.UpdateFlowDB(ctx, req.Id, req.Name, req.Description, req.TenantId, req.FlowType)

	if err != nil {
		logs.Error(ctx, "event=UpdateFlow error=%v", err.Error())
		return nil, err
	}
	// Update the flow_edge table
	err = a.UpdateFlowEdgeDB(ctx, req.Id, req.EdgeIds)
	if err != nil {
		logs.Error(ctx, "event=UpdateFlow error=%v", err.Error())
		return nil, err
	}

	// Update the flow_core table
	err = a.UpdateFlowCoreDB(ctx, req.Id, req.CoreIds)
	if err != nil {
		logs.Error(ctx, "event=UpdateFlow error=%v", err.Error())
		return nil, err
	}

	// Update the flow_app table
	err = a.UpdateFlowAppDB(ctx, req.Id, req.AppIds)
	if err != nil {
		logs.Error(ctx, "event=UpdateFlow error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *FlowService) DeleteFlow(ctx context.Context, req *freezonex_openiiot_api.DeleteFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteFlowResponse, error) {
	// Delete related records in the flow_edge table
	if err := a.DeleteFlowEdgeRecords(ctx, req.Id); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Delete related records in the flow_core table
	if err := a.DeleteFlowCoreRecords(ctx, req.Id); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Delete related records in the flow_app table
	if err := a.DeleteFlowAppRecords(ctx, req.Id); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Finally, delete the main flow record
	if err := a.DeleteFlowDB(ctx, req.Id); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Prepare and return the response
	resp := new(freezonex_openiiot_api.DeleteFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
