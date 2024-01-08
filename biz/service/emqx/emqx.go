package emqx

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *EmqxService) GetStatus(ctx context.Context, req *freezonex_openiiot_api.EmqxGetStatusRequest, c *app.RequestContext) (*freezonex_openiiot_api.EmqxGetStatusResponse, error) {
	status, err := a.client.GetStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := new(freezonex_openiiot_api.EmqxGetStatusResponse)
	resp.Status = status
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *EmqxService) CreateBridge(ctx context.Context, req *freezonex_openiiot_api.EmqxCreateBridgeRequest, c *app.RequestContext) (*freezonex_openiiot_api.EmqxCreateBridgeResponse, error) {
	status, err := a.client.CreateBridge(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := new(freezonex_openiiot_api.EmqxCreateBridgeResponse)
	resp.BridgeStatus = status
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *EmqxService) CreateRule(ctx context.Context, req *freezonex_openiiot_api.EmqxCreateRuleRequest, c *app.RequestContext) (*freezonex_openiiot_api.EmqxCreateRuleResponse, error) {
	rule, err := a.client.CreateRule(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := new(freezonex_openiiot_api.EmqxCreateRuleResponse)
	resp.RuleResp = rule
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
