package wms_storage_location

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_storagelocationmaterial"
)

func (a *WmsStorageLocationService) AddStorageLocation(ctx context.Context, req *freezonex_openiiot_api.AddStorageLocationRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddStorageLocationResponse, error) {
	storagelocationID, err := a.AddStorageLocationDB(ctx, common.StringToInt64(req.WarehouseId), req.Name, req.Occupied, "")
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
	storagelocations, err := a.GetStorageLocationDB(ctx, common.StringToInt64(req.WarehouseId), req.Name, nil, "")

	if err != nil {
		logs.Error(ctx, "event=GetStorageLocation error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetStorageLocationResponse)
	data := make([]*freezonex_openiiot_api.StorageLocation, 0)

	for _, v := range storagelocations {
		storagelocationmaterialService := wms_storagelocationmaterial.DefaultStorageLocationMaterialService()
		slmaterials, err := storagelocationmaterialService.GetStorageLocationMaterialDB(ctx, 0, v.ID, 0, 0)

		if err != nil {
			logs.Error(ctx, "event=GetStorageLocationMaterialMaterialDB error=%v", err.Error())
			return nil, err
		}
		var convertedMaterials []*freezonex_openiiot_api.StorageLocationMaterial

		for _, slmaterial := range slmaterials {
			u := a.db.DBOpeniiotQuery.WmsMaterial
			e := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial

			err1 := u.WithContext(ctx).Select(u.Name, e.MaterialID).LeftJoin(e, e.MaterialID.EqCol(u.ID))
			existRecord, _ := err1.Where(e.MaterialID.Eq(slmaterial.MaterialID)).First()

			if err != nil {
				logs.Error(ctx, "event=GetWmsMaterialDB error=%v", err.Error())
				return nil, err
			}
			convertedmaterial := &freezonex_openiiot_api.StorageLocationMaterial{
				MaterialId:   common.Int64ToString(slmaterial.MaterialID),
				MaterialName: existRecord.Name,
				Quantity:     slmaterial.Quantity,
			}
			convertedMaterials = append(convertedMaterials, convertedmaterial)
		}

		data = append(data, &freezonex_openiiot_api.StorageLocation{
			Id:          common.Int64ToString(v.ID),
			Name:        v.Name,
			WarehouseId: common.Int64ToString(v.WarehouseID),
			Occupied:    *v.Occupied,
			Materials:   convertedMaterials,
			CreateTime:  common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:  common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateStorageLocation will update storagelocation record
func (a *WmsStorageLocationService) UpdateStorageLocation(ctx context.Context, req *freezonex_openiiot_api.UpdateStorageLocationRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateStorageLocationResponse, error) {
	err := a.UpdateStorageLocationDB(ctx, common.StringToInt64(req.Id), common.StringToInt64(req.WarehouseId), req.Name, &req.Occupied, "")
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
