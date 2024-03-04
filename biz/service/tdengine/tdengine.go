package tdengine

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
)

// Token function requests an access token from the TDEngine API.
func (a *TDEngineService) TestConnection(ctx context.Context, req *freezonex_openiiot_api.TDEngineTestConnectionRequest, c *app.RequestContext) (*freezonex_openiiot_api.TDEngineTestConnectionResponse, error) {
	dsn := ConstructDSNString(req.Dsn.Username, req.Dsn.Password, req.Dsn.Host)
	isSuccessful := a.T.TestConnect(dsn)
	resp := new(freezonex_openiiot_api.TDEngineTestConnectionResponse)
	resp.Successful = isSuccessful
	resp.BaseResp = middleware.SuccessResponseOK
	return resp, nil
}

func (a *TDEngineService) Exec(ctx context.Context, req *freezonex_openiiot_api.TDEngineExecRequest, c *app.RequestContext) (*freezonex_openiiot_api.TDEngineExecResponse, error) {
	dsn := ConstructDSNString(req.Dsn.Username, req.Dsn.Password, req.Dsn.Host)
	rowsAffected, err := a.T.Exec(dsn, req.Sql)
	if err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.TDEngineExecResponse)
	resp.RowsAffected = rowsAffected
	resp.BaseResp = middleware.SuccessResponseOK
	return resp, nil
}

func (a *TDEngineService) BatchExec(ctx context.Context, req *freezonex_openiiot_api.TDEngineBatchExecRequest, c *app.RequestContext) (*freezonex_openiiot_api.TDEngineBatchExecResponse, error) {
	dsn := ConstructDSNString(req.Dsn.Username, req.Dsn.Password, req.Dsn.Host)
	rowsAffected, err := a.T.BatchExec(dsn, req.Sql)
	if err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.TDEngineBatchExecResponse)
	resp.RowsAffected = rowsAffected
	resp.BaseResp = middleware.SuccessResponseOK
	return resp, nil
}

func (a *TDEngineService) Query(ctx context.Context, req *freezonex_openiiot_api.TDEngineQueryRequest, c *app.RequestContext) (*freezonex_openiiot_api.TDEngineQueryResponse, error) {
	dsn := ConstructDSNString(req.Dsn.Username, req.Dsn.Password, req.Dsn.Host)
	result, err := a.T.Query(dsn, req.Sql)
	if err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.TDEngineQueryResponse)
	/*data := make([]*freezonex_openiiot_api.TDEngineRow, 0)
	for _, v := range result {
		data = append(data, &freezonex_openiiot_api.TDEngineRow{
			Time:  v.Time,
			Value: v.Value,
		})
	}*/

	resp.Data = result
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
