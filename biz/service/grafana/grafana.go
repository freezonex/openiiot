package grafana

import (
	"context"
	"errors"
	"fmt"
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
