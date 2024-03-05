package wms_storage_location

import (
	"context"
	"freezonex/openiiot/biz/service/utils/common"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *WmsStorageLocationService) AddStorageLocation(ctx context.Context, req *freezonex_openiiot_api.AddStorageLocationRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddStorageLocationResponse, error) {
	storagelocationID, err := a.AddStorageLocationDB(ctx, common.StringToInt64(req.WarehouseId), req.Name, req.Occupied, req.MaterialName)
	if err != nil {
		logs.Error(ctx, "event=AddStorageLocation error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddStorageLocationResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(storagelocationID)

	return resp, nil
}

// GetStorageLocation will get storagelocation record in condition
func (a *WmsStorageLocationService) GetStorageLocation(ctx context.Context, req *freezonex_openiiot_api.GetStorageLocationRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetStorageLocationResponse, error) {
	storagelocations, err := a.GetStorageLocationDB(ctx, common.StringToInt64(req.WarehouseId), req.Name, &req.Occupied, req.MaterialName)

	if err != nil {
		logs.Error(ctx, "event=GetStorageLocation error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetStorageLocationResponse)
	data := make([]*freezonex_openiiot_api.StorageLocation, 0)

	for _, v := range storagelocations {
		data = append(data, &freezonex_openiiot_api.StorageLocation{
			Id:           common.Int64ToString(v.ID),
			Name:         v.Name,
			WarehouseId:  common.Int64ToString(v.WarehouseID),
			Occupied:     *v.Occupied,
			MaterialName: *v.MaterialName,
			CreateTime:   common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:   common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateStorageLocation will update storagelocation record
func (a *WmsStorageLocationService) UpdateStorageLocation(ctx context.Context, req *freezonex_openiiot_api.UpdateStorageLocationRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateStorageLocationResponse, error) {
	err := a.UpdateStorageLocationDB(ctx, common.StringToInt64(req.Id), common.StringToInt64(req.WarehouseId), req.Name, &req.Occupied, req.MaterialName)
	if err != nil {
		logs.Error(ctx, "event=UpdateStorageLocation error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateStorageLocationResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteStorageLocation will delete storagelocation record
func (a *WmsStorageLocationService) DeleteStorageLocation(ctx context.Context, req *freezonex_openiiot_api.DeleteStorageLocationRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteStorageLocationResponse, error) {
	//Delete storagelocation also should delete storagelocation , edge pool, core pool, application pool, flow
	/*err := a.DeleteStorageLocationStorageLocationDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteStorageLocation storagelocation error=%v", err.Error())
		return nil, err
	}*/

	// Delete storagelocation
	err := a.DeleteStorageLocationDB(ctx, common.StringToInt64(req.Id))
	if err != nil {
		logs.Error(ctx, "event=DeleteStorageLocation error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteStorageLocationResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
