package emqx

import (
	"context"
	"encoding/base64"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/http_utils"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/url"
)

// Client is a Grafana API client.
type EmqxClient struct {
	dsn freezonex_openiiot_api.EmqxDSN
}

func generateBasicAuth(username string, password string) (string, error) {
	auth := username + ":" + password
	authBase64 := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + authBase64, nil
}

func (c *EmqxClient) GetStatus(ctx context.Context, req *freezonex_openiiot_api.EmqxGetStatusRequest) (*freezonex_openiiot_api.EmqxStatusStruct, error) {
	basicAuth, err := generateBasicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return nil, err
	}

	urlPath := req.Dsn.Host + "/api/v5/status"

	query := url.Values{}
	query.Set("format", "json")

	result := &freezonex_openiiot_api.EmqxStatusStruct{}
	err = http_utils.GetWithUnmarshal(ctx, result, urlPath, []http_utils.Query{{Key: "format", Value: "json"}}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})

	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *EmqxClient) CreateBridge(ctx context.Context, req *freezonex_openiiot_api.EmqxCreateBridgeRequest) (*freezonex_openiiot_api.BridgeStatusResponse, error) {
	basicAuth, err := generateBasicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return nil, err
	}

	urlPath := req.Dsn.Host + "/api/v5/bridges"

	result := &freezonex_openiiot_api.BridgeStatusResponse{}
	err = http_utils.PostWithUnmarshal(ctx, result, urlPath, req.Bridge, []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	logs.Debug("client level response", result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *EmqxClient) CreateRule(ctx context.Context, req *freezonex_openiiot_api.EmqxCreateRuleRequest) (*freezonex_openiiot_api.RuleResponseStruct, error) {
	basicAuth, err := generateBasicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return nil, err
	}

	urlPath := req.Dsn.Host + "/api/v5/rules"

	result := &freezonex_openiiot_api.RuleResponseStruct{}
	err = http_utils.PostWithUnmarshal(ctx, result, urlPath, req.Rule, []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	logs.Debug("client level response", result)
	if err != nil {
		return nil, err
	}
	return result, err
}
