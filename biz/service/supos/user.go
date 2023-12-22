package supos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/cache"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func (a *SuposService) GetSuposUser(ctx context.Context, req *freezonex_openiiot_api.GetSuposUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetSuposUserResponse, error) {
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

	resp := new(freezonex_openiiot_api.GetSuposUserResponse)
	data := make([]*freezonex_openiiot_api.SuposUser, 0)
	for _, v := range result.List {
		data = append(data, &freezonex_openiiot_api.SuposUser{
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
	loginUserRoleList, _ := cache.Get("CurrentUserRole")
	//loginUsername = "hongzhi"

	resp := new(freezonex_openiiot_api.GetCurrentUserResponse)
	data := new(freezonex_openiiot_api.User)
	if loginUsername == "" {
		logs.Error(ctx, "event=GetTenant, current user is empty")
		return nil, errors.New("current user is empty")
	}

	data.Username = loginUsername
	data.Role = loginUserRoleList
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Data = data

	return resp, nil
}
