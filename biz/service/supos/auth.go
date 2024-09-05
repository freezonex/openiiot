package supos

import (
	"context"
	"encoding/json"
	"fmt"
	"freezonex/openiiot/biz/service/user"
	"strconv"
	"strings"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/cache"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

// Authorize composes the URL for OAuth2 authorization.
func (a *SuposService) Authorize(redirectUri, state string) string {
	return fmt.Sprintf("%s%s?responseType=code&state=%s&redirectUri=%s",
		a.c.URL,
		OAUTH2_GET_AUTHORIZE_CODE.GetURL(),
		state,
		redirectUri)
}

// AccessToken function requests an access token from the Supos API.
func (a *SuposService) GetAccessToken(ctx context.Context, req *freezonex_openiiot_api.GetAccessTokenRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetAccessTokenResponse, error) {
	// Generate a new UUID and remove hyphens
	token := strings.Replace(uuid.New().String(), "-", "", -1)
	logoutUri := fmt.Sprintf(a.c.LOGOUT, token)

	//req.Code = "376bdcc84304e5118bd6fb976abda5f4"
	//logoutUri = "https://566ow1752.goho.co/auth/logout/8bd64a78862845e1b62d573680fc973b"
	requestBody := map[string]string{
		"grantType": "authorization_code",
		"code":      req.Code,
		"logoutUri": logoutUri,
	}

	jsonValue, _ := json.Marshal(requestBody)
	apiResponse, _ := a.DoRequest(OAUTH2_GET_ACCESS_TOKEN, string(jsonValue), nil)
	if apiResponse.HttpCode == 200 {
		var result map[string]interface{}
		err := json.Unmarshal([]byte(apiResponse.ResponseBody), &result)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling JSON when get access token: %w", err)
		}
		//accessToken1 := result["accessToken"].(string)
		username, _ := result["username"].(string)
		/*personCode := result["personCode"]
		companyCode := result["companyCode"]
		accessToken := result["accessToken"]
		refreshToken := result["refreshToken"]*/
		cache.Set(token, apiResponse.ResponseBody)
		if username != "" {
			apiResponse1, err := a.DoRequest(USER_DETAIL, "", nil, username)
			if err != nil {
				return nil, fmt.Errorf("error send request to Supos to ge detail for user: %s, %w", username, err)
			}
			if apiResponse1.HttpCode != 200 {
				return nil, fmt.Errorf("supos server response error, %w", err)
			}

			if apiResponse1.HttpCode == 200 {
				var user1 UserDetail
				err := json.Unmarshal([]byte(apiResponse1.ResponseBody), &user1)
				if err != nil {
					return nil, fmt.Errorf("error unmarshaling JSON when get userdetail: %w", err)
				}

				username1 := user1.PersonCode
				authid := user1.PersonCode
				userRole := classifyUserRole(user1)
				userService := user.DefaultUserService()
				id, tenantId, err := userService.AddSuposUserID(ctx, username1, authid, userRole)

				cache.Set("CurrentUserId", strconv.FormatInt(id, 10)) //authid,tenantid?
				cache.Set("CurrentUsername", username1)
				cache.Set("CurrentUserRole", userRole)
				cache.Set("CurrentTenantId", strconv.FormatInt(tenantId, 10))

				ctx = context.WithValue(ctx, "currentuserid", id)
				ctx = context.WithValue(ctx, "currentusername", username1)
				ctx = context.WithValue(ctx, "currentuserrole", userRole)
				ctx = context.WithValue(ctx, "currenttenantid", tenantId)
				if err != nil {
					return nil, err
				}

				err = userService.UpdateSuposToken(ctx, id, token)
				if err != nil {
					return nil, err
				}

			}
		}
	}

	resp := new(freezonex_openiiot_api.GetAccessTokenResponse)
	resp.Code = 0
	resp.Msg = "Success"
	resp.Data = token
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func classifyUserRole(user UserDetail) string {
	if user.AccountType == 1 {
		return "Admin"
	}
	for _, role := range user.UserRoleList {
		switch role.Name {
		case "systemRole", "companySystemRole":
			return "Admin"
		}
	}

	if user.Username == "developer" {
		return "Editor"
	}

	return "Viewer"
}

func (a *SuposService) Logout(ctx context.Context, req *freezonex_openiiot_api.LogoutRequest, c *app.RequestContext) (*freezonex_openiiot_api.LogoutResponse, error) {
	tokenKey := req.TokenKey
	cache.Delete(tokenKey)

	userservice := user.DefaultUserService()
	userservice.DeleteUserToken(ctx, tokenKey)
	resp := new(freezonex_openiiot_api.LogoutResponse)
	resp.Code = 0
	resp.Msg = "Success"
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *SuposService) Login(ctx context.Context, req *freezonex_openiiot_api.LoginRequest, c *app.RequestContext) (*freezonex_openiiot_api.LoginResponse, error) {
	userService := user.DefaultUserService()

	id, Accesstoken, userRole, tenantId, err := userService.UserLoginDB(ctx, req.Username, req.Password, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=GetUser error=%v", err.Error())
		return nil, err
	}
	cache.Set("CurrentUserId", strconv.FormatInt(id, 10)) //authid,tenantid?
	cache.Set("CurrentUsername", req.Username)
	cache.Set("CurrentUserRole", userRole)
	cache.Set("CurrentTenantId", strconv.FormatInt(tenantId, 10))

	context.WithValue(ctx, "currentid", id)
	context.WithValue(ctx, "currentusername", req.Username)
	context.WithValue(ctx, "currentuserrole", userRole)
	context.WithValue(ctx, "currenttenantid", tenantId)
	err = userService.UpdateSuposToken(ctx, id, Accesstoken)

	if err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.LoginResponse)
	resp.Accesstoken = Accesstoken
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
