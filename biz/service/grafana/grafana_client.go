package grafana

import (
	"context"
	"encoding/base64"
	"fmt"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/http_utils"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"log"
)

// Client is a Grafana API client.
type GrafanaClient struct {
	dsn freezonex_openiiot_api.GrafanaDSN
}

func generateBaicAuth(username string, password string) (string, error) {
	auth := username + ":" + password
	authBase64 := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + authBase64, nil
}

func (c *GrafanaClient) GetDataSource(ctx context.Context, req *freezonex_openiiot_api.GrafanaGetDataSourceRequest) ([]*freezonex_openiiot_api.DataSource, error) {

	basicAuth, err := generateBaicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return nil, err
	}
	result := make([]*DataSource, 0)
	err = http_utils.GetWithUnmarshal(ctx, &result, req.Dsn.Host+"/api/datasources", []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	if err != nil {
		return nil, err
	}
	formattedResult, err := convertDataSources(result)
	if err != nil {
		log.Fatalf("Error converting data sources: %v", err)
		return nil, err
	}

	return formattedResult, nil
}

func (c *GrafanaClient) CreateDashBoard(ctx context.Context, req *freezonex_openiiot_api.GrafanaCreateDashboardRequest) (*freezonex_openiiot_api.DashboardSaveResponse, error) {
	data, err := convertDashboardModel(req.Dashboard)
	if err != nil {
		logs.Error("Cannot unmarshal model in dashboard", req.Dashboard)
		return nil, err
	}
	basicAuth, err := generateBaicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return nil, err
	}
	result := freezonex_openiiot_api.DashboardSaveResponse{}
	err = http_utils.PostWithUnmarshal(ctx, &result, req.Dsn.Host+"/api/dashboards/db", data, []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	if err != nil {
		return nil, err
	}
	logs.Debug("Client level debug", &result)
	return &result, err
}

//func (c *GrafanaClient) dashboard(path string) (*Dashboard, error) {
//	result := &Dashboard{}
//	err := c.request("GET", path, nil, nil, &result)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, err
//}

func (c *GrafanaClient) DashboardByUID(ctx context.Context, req *freezonex_openiiot_api.GrafanaSaveDashboardByUidRequest) (*Dashboard, error) {
	basicAuth, err := generateBaicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return nil, err
	}
	result := &Dashboard{}
	urlPath := req.Dsn.Host + fmt.Sprintf("/api/dashboards/uid/%s", req.Uid)
	err = http_utils.GetWithUnmarshal(ctx, result, urlPath, []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	if err != nil {
		return nil, err
	}
	result.FolderID = result.Meta.Folder
	return result, err
}

func (c *GrafanaClient) NewDataSource(ctx context.Context, req *freezonex_openiiot_api.GrafanaCreateDataSourceRequest) (int64, error) {
	data, err := convertPbDataSource(req.DataSource)
	if err != nil {
		return 0, err
	}
	basicAuth, err := generateBaicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return 0, err
	}
	result := struct {
		ID int64 `json:"id"`
	}{}
	err = http_utils.PostWithUnmarshal(ctx, &result, req.Dsn.Host+"/api/datasources", data, []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	if err != nil {
		return 0, err
	}
	logs.Debug("new data source", result)

	return result.ID, err
}

func (c *GrafanaClient) DeleteDataSourceByName(ctx context.Context, req *freezonex_openiiot_api.GrafanaDeleteDataSourceRequest) error {
	urlPath := req.Dsn.Host + fmt.Sprintf("/api/datasources/name/%s", req.Name)
	basicAuth, err := generateBaicAuth(req.GetDsn().Username, req.GetDsn().Password)
	if err != nil {
		return err
	}
	_, err = http_utils.Delete(ctx, urlPath, []http_utils.Query{}, []http_utils.Path{}, []http_utils.Header{{Key: "Authorization", Value: basicAuth}})
	if err != nil {
		return err
	}
	return err
}
