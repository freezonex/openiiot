package supos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/cache"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *SuposService) GetAllUser(ctx context.Context, req *freezonex_openiiot_api.GetUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetUserResponse, error) {
	param := req.Param
	if strings.TrimSpace(param.CompanyCode) == "" {
		param.CompanyCode = "default_org_company"
	}

	searchParam := map[string]interface{}{
		"keyword":     param.Keyword,
		"pageIndex":   param.PageIndex,
		"pageSize":    param.PageSize,
		"companyCode": param.CompanyCode,
	}
	jsonString, _ := json.Marshal(searchParam)

	var result PageResult
	apiResponse, _ := a.DoRequest(USER_LISTS, string(jsonString), nil)
	if apiResponse.HttpCode != 200 {
		return nil, errors.New("cannot get response from Supos server when get user list")
	}

	err := json.Unmarshal([]byte(apiResponse.ResponseBody), &result)
	if err != nil {
		fmt.Println("error unmarshaling JSON:", err)
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetUserResponse)
	data := make([]*freezonex_openiiot_api.User, 0)
	for _, v := range result.List {
		data = append(data, &freezonex_openiiot_api.User{
			Username:    v.Username,
			UserDesc:    v.UserDesc,
			AccountType: int32(v.AccountType),
			LockStatus:  int32(v.LockStatus),
			PersonCode:  v.PersonCode,
			PersonName:  v.PersonName,
			CreateTime:  v.CreateTime,
			ModifyTime:  v.ModifyTime,
		})
	}
	resp.Pagination = &freezonex_openiiot_api.Pagination{}
	resp.Pagination.Total = int32(result.Pagination.Total)
	resp.Pagination.PageSize = int32(result.Pagination.PageSize)
	resp.Pagination.PageIndex = int32(result.Pagination.PageIndex)
	resp.List = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *SuposService) GetCurrentUser(ctx context.Context, req *freezonex_openiiot_api.GetCurrentUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetCurrentUserResponse, error) {
	loginUsername, _ := cache.Get("CurrentUsername")

	resp := new(freezonex_openiiot_api.GetCurrentUserResponse)
	if loginUsername != "" {
		loginUserRoleList, _ := cache.Get("CurrentUserType")
		resp.Code = 0
		resp.Message = "Success"
		resp.Data = loginUserRoleList
	} else {
		resp.Code = 0
		resp.Message = "No login user information"
	}

	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
