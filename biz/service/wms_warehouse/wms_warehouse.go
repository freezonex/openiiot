package wms_warehouse

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_material"
	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
	"freezonex/openiiot/biz/service/wms_storagelocationmaterial"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func (a *WmsWarehouseService) AddWmsWarehouse(ctx context.Context, req *freezonex_openiiot_api.AddWarehouseRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddWarehouseResponse, error) {
	wmsID, err := a.AddWmsWarehouseDB(ctx, req.Name, req.WarehouseId, req.Type, req.Manager, req.Department, req.Email, req.ProjectGroup, req.Note)
	if err != nil {
		logs.Error(ctx, "event=AddWmsWarehouse error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddWarehouseResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(wmsID)

	return resp, nil
}

// GetWmsWarehouse will get wms record in condition
func (a *WmsWarehouseService) GetWmsWarehouse(ctx context.Context, req *freezonex_openiiot_api.GetWarehouseRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetWarehouseResponse, error) {
	wmss, err := a.GetWmsWarehouseDB(ctx, common.StringToInt64(req.Id), req.Name, req.WarehouseId, req.Type, req.Manager, req.Department, req.Email, req.ProjectGroup)

	if err != nil {
		logs.Error(ctx, "event=GetWmsWarehouse error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetWarehouseResponse)
	data := make([]*freezonex_openiiot_api.Warehouse, 0)

	for _, v := range wmss {
		storagelocationService := storagelocation.DefaultStorageLocationService()
		storages, err := storagelocationService.GetStorageLocationDB(ctx, common.StringToInt64(v.WarehouseID), "", nil, "")
		var convertedStorages []*freezonex_openiiot_api.StorageLocation

		for _, storage := range storages {
			var convertedMaterials []*freezonex_openiiot_api.StorageLocationMaterial
			storagelocationmaterialService := wms_storagelocationmaterial.DefaultStorageLocationMaterialService()
			slmaterials, err := storagelocationmaterialService.GetStorageLocationMaterialMaterialDB(ctx, 0, v.ID)
			if err != nil {
				logs.Error(ctx, "event=storagelocationmaterialService error=%v", err.Error())
				return nil, err
			}
			for _, slmaterial := range slmaterials {
				materialService := wms_material.DefaultWmsMaterialService()
				materials, err := materialService.GetWmsMaterialDB(ctx, slmaterial.MaterialID, "", "", "", "", "")
				if err != nil {
					logs.Error(ctx, "event=GetWmsMaterialDB error=%v", err.Error())
					return nil, err
				}
				convertedmaterial := &freezonex_openiiot_api.StorageLocationMaterial{
					MaterialId:   common.Int64ToString(slmaterial.MaterialID),
					MaterialName: materials[0].Name,
					Quantity:     slmaterial.Quantity,
				}
				convertedMaterials = append(convertedMaterials, convertedmaterial)
			}

			convertedStorage := &freezonex_openiiot_api.StorageLocation{
				Id:          common.Int64ToString(storage.ID),
				WarehouseId: common.Int64ToString(storage.WarehouseID),
				Name:        storage.Name,
				Occupied:    *storage.Occupied,
				Materials:   convertedMaterials,
				CreateTime:  common.GetTimeStringFromTime(&storage.CreateTime),
				UpdateTime:  common.GetTimeStringFromTime(&storage.UpdateTime),
			}
			convertedStorages = append(convertedStorages, convertedStorage)
		}

		if err != nil {
			return nil, err
		}

		data = append(data, &freezonex_openiiot_api.Warehouse{
			Id:               common.Int64ToString(v.ID), // Converts int64 ID to string using a custom common package function
			Name:             v.Name,
			WarehouseId:      v.WarehouseID,
			Type:             v.Type,
			Manager:          *v.Manager,
			Department:       *v.Department,
			Email:            *v.Email,
			ProjectGroup:     *v.ProjectGroup,
			Note:             *v.Note,
			StorageLocations: convertedStorages,
			CreateTime:       common.GetTimeStringFromTime(&v.CreateTime), // Converts time.Time to string using a custom common package function
			UpdateTime:       common.GetTimeStringFromTime(&v.UpdateTime), // Converts time.Time to string using a custom common package function
		})

	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateWmsWarehouse will update wms record
func (a *WmsWarehouseService) UpdateWmsWarehouse(ctx context.Context, req *freezonex_openiiot_api.UpdateWarehouseRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateWarehouseResponse, error) {
	err := a.UpdateWmsWarehouseDB(ctx, common.StringToInt64(req.Id), req.Name, req.WarehouseId, req.Type, req.Manager, req.Department, req.Email, req.ProjectGroup, req.Note)
	if err != nil {
		logs.Error(ctx, "event=UpdateWmsWarehouse error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateWarehouseResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteWmsWarehouse will delete wms record
func (a *WmsWarehouseService) DeleteWmsWarehouse(ctx context.Context, req *freezonex_openiiot_api.DeleteWarehouseRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteWarehouseResponse, error) {
	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
	/*err := a.DeleteWmsWarehouseUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteWmsWarehouse user error=%v", err.Error())
		return nil, err
	}*/

	// Delete wms
	//delete all the users and flows?
	//_, err := a.DeleteWmsWarehouseDB(ctx, common.StringToInt64(req.Id))

	err := a.DeleteWmsWarehouseDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=DeleteWmsWarehouse error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteWarehouseResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
