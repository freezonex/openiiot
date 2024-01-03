package grafana

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"freezonex/openiiot/biz/service/utils/cache"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/google/uuid"
)

// Authorize composes the URL for OAuth2 authorization.
func (a *GrafanaService) Authorize(ctx context.Context, c *app.RequestContext) {
	redirectUri := c.Query("redirect_uri")
	state := c.Query("state")
	code := strings.Replace(uuid.New().String(), "-", "", -1)
	url := redirectUri + "?state=" + state + "&code=" + code
	logs.CtxInfof(ctx, url)
	cache.Set("grafana_code", code)
	c.Redirect(consts.StatusFound, []byte(url))
}

// Token function requests an access token from the Grafana API.
func (a *GrafanaService) GetAccessToken(ctx context.Context, req *freezonex_openiiot_api.GrafanaAccessTokenRequest, c *app.RequestContext) (*freezonex_openiiot_api.GrafanaAccessTokenResponse, error) {
	if req.Code == "" {
		return nil, errors.New("grafana authorization code is empty")
	}
	grafana_code, _ := cache.Get("grafana_code")
	if req.Code != grafana_code {
		return nil, errors.New("grafana authorization code not match")
	}

	// Generate a new UUID and remove hyphens
	accessToken := strings.Replace(uuid.New().String(), "-", "", -1)
	refreshToken := strings.Replace(uuid.New().String(), "-", "", -1)

	resp := new(freezonex_openiiot_api.GrafanaAccessTokenResponse)
	resp.AccessToken = accessToken
	resp.TokenType = "Bearer"
	resp.ExpiresIn = 1800
	resp.RefreshToken = refreshToken
	resp.BaseResp = middleware.SuccessResponseOK

	cache.Set("grafana_accessToken", accessToken)
	cache.Set("grafana_refreshToken", refreshToken)
	return resp, nil
}

func (a *GrafanaService) GetUser(ctx context.Context, req *freezonex_openiiot_api.GrafanaUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.GrafanaUserResponse, error) {
	headerAuth := req.Authorization
	if !strings.HasPrefix(headerAuth, "Bearer") {
		return nil, errors.New("grafana authorization header not valid, must start with Bearer")
	}
	token := strings.TrimPrefix(headerAuth, "Bearer ")
	grafana_accessToken, _ := cache.Get("grafana_accessToken")
	if token != grafana_accessToken {
		return nil, fmt.Errorf("grafana access token: %v, not match with %v", token, grafana_accessToken)
	}

	currentUserId, _ := cache.Get("CurrentUserId")
	currentUsername, _ := cache.Get("CurrentUsername")
	currentUserRole, _ := cache.Get("CurrentUserRole")
	resp := new(freezonex_openiiot_api.GrafanaUserResponse)
	resp.Sub = currentUserId
	resp.Name = currentUsername
	resp.Email = currentUsername + "@localhost"
	resp.Role = []string{currentUserRole}
	//resp.Role = []string{"Editor"}
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *GrafanaService) GetDatasource(ctx context.Context, req *freezonex_openiiot_api.GrafanaGetDataSourceRequest, c *app.RequestContext)(*freezonex_openiiot_api.GrafanaGetDataSourceResponse, error) {
	dataSources, err := a.client.GetDataSource(req.Dsn)
	if err != nil {
		return nil, err
	}
	
	resp := new(freezonex_openiiot_api.GrafanaGetDataSourceResponse)
	resp.DataSources = dataSources
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil

}

func (a *GrafanaService) NewDataSource(ctx context.Context, req *freezonex_openiiot_api.GrafanaCreateDataSourceRequest, c *app.RequestContext)(*freezonex_openiiot_api.GrafanaCreateDataSourceResponse, error) {
	dataSource, err := a.client.NewDataSource(req.Dsn, req.DataSource)
	if err != nil {
		return nil, err
	}
	
	resp := new(freezonex_openiiot_api.GrafanaCreateDataSourceResponse)
	resp.Message = fmt.Sprintf("New Data Scource %d created ", dataSource)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil

}
func (a *GrafanaService) DeleteDataSource(ctx context.Context, req *freezonex_openiiot_api.GrafanaDeleteDataSourceRequest, c *app.RequestContext)(*freezonex_openiiot_api.GrafanaDeleteDataSourceResponse, error) {
	err := a.client.DeleteDataSourceByName(req.Dsn, req.Name)
	if err != nil {
		return nil, err
	}
	
	resp := new(freezonex_openiiot_api.GrafanaDeleteDataSourceResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Message = fmt.Sprintf("Data Scource %s deleted ", req.Name)
	return resp, nil

}
func (a *GrafanaService) CreateDashBoard(ctx context.Context, req *freezonex_openiiot_api.GrafanaCreateDashboardRequest, c *app.RequestContext)(*freezonex_openiiot_api.GrafanaCreateDashboardResponse, error) {
	saveResponse, err := a.client.CreateDashBoard(req.Dsn, req.Dashboard)
	if err != nil {
		return nil, err
	}
	resp := new(freezonex_openiiot_api.GrafanaCreateDashboardResponse)
	resp.SaveResponse = saveResponse
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil

}

func (a *GrafanaService) SaveDashboardByUID(ctx context.Context, req *freezonex_openiiot_api.GrafanaSaveDashboardByUidRequest, c *app.RequestContext)(*freezonex_openiiot_api.GrafanaSaveDashboardByUidResponse, error) {
	returnedDashboard, err := a.client.DashboardByUID(req.Dsn, req.Uid, req.SavePath)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.MarshalIndent(returnedDashboard, "", "    ")
	if err != nil {
		logs.Error("error marshalling to JSON: %s", err)
	}
	err = os.WriteFile(req.SavePath, jsonData, 0644)
	if err != nil {
		logs.Error("error writing JSON to file: %s", err)
	}
	resp := new(freezonex_openiiot_api.GrafanaSaveDashboardByUidResponse)
	resp.Message = fmt.Sprintf("Dashboard JSON file saved in %s", req.SavePath)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil

}
