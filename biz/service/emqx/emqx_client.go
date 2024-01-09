package emqx

import (
	"context"
	"encoding/base64"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/http_utils"
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
