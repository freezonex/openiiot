package user

import (
	"context"
	"freezonex/openiiot/biz/service/utils/common"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *UserService) AddUser(ctx context.Context, req *freezonex_openiiot_api.AddUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddUserResponse, error) {
	userID, err := a.AddUserDB(ctx, req.Username, req.Password, req.Description, req.TenantId, req.Role, req.AuthId, req.Source)
	if err != nil {
		logs.Error(ctx, "event=AddUser error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddUserResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = userID

	return resp, nil
}

// GetUser will get user record in condition
func (a *UserService) GetUser(ctx context.Context, req *freezonex_openiiot_api.GetUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetUserResponse, error) {
	users, err := a.GetUserDB(ctx, req.Username, req.TenantId, req.Role, req.AuthId, req.Source)

	if err != nil {
		logs.Error(ctx, "event=GetUser error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetUserResponse)
	data := make([]*freezonex_openiiot_api.User, 0)
	for _, v := range users {
		data = append(data, &freezonex_openiiot_api.User{
			Id:         v.ID,
			Username:   v.Username,
			TenantId:   v.TenantID,
			Role:       v.Role,
			AuthId:     *v.AuthID,
			Source:     *v.Source,
			CreateTime: common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime: common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateUser will update user record
func (a *UserService) UpdateUser(ctx context.Context, req *freezonex_openiiot_api.UpdateUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateUserResponse, error) {
	err := a.UpdateUserDB(ctx, req.Id, req.Username, req.Password, req.Description, req.TenantId, req.Role, req.AuthId, req.Source)
	if err != nil {
		logs.Error(ctx, "event=UpdateUser error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateUserResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteUser will delete user record
func (a *UserService) DeleteUser(ctx context.Context, req *freezonex_openiiot_api.DeleteUserRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteUserResponse, error) {
	//Delete user also should delete user , edge pool, core pool, application pool, flow
	/*err := a.DeleteUserUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteUser user error=%v", err.Error())
		return nil, err
	}*/

	// Delete user
	err := a.DeleteUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteUser error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteUserResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
