package emqx

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func (a *EmqxService) GetStatus(ctx context.Context, req *freezonex_openiiot_api.EmqxGetStatusRequest, c *app.RequestContext)(*freezonex_openiiot_api.EmqxGetStatusResponse, error) {
	logs.Debug("Service level")
	status, err := a.client.GetStatus(req.Dsn)
	if err != nil {
		return nil, err
	}
	resp := new(freezonex_openiiot_api.EmqxGetStatusResponse)
	resp.Status = status
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}