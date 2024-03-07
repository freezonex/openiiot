package wms_material

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
	"freezonex/openiiot/biz/service/wms_storagelocationmaterial"
	"freezonex/openiiot/biz/service/wms_warehouse"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
)

func (a *WmsMaterialService) AddWmsMaterial(ctx context.Context, req *freezonex_openiiot_api.AddMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddMaterialResponse, error) {
	wmsID, err := a.AddWmsMaterialDB(ctx, req.ProductCode, req.Name, req.ProductType, req.Unit, req.Note)
	if err != nil {
		logs.Error(ctx, "event=AddWmsMaterial error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddMaterialResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(wmsID)

	return resp, nil
}

// GetWmsMaterial will get wms record in condition
func (a *WmsMaterialService) GetWmsMaterial(ctx context.Context, req *freezonex_openiiot_api.GetMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetMaterialResponse, error) {
	wmss, err := a.GetWmsMaterialDB(ctx, common.StringToInt64(req.Id), req.ProductCode, req.Name, req.ProductType, req.Unit, "")

	if err != nil {
		logs.Error(ctx, "event=GetWmsMaterial error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetMaterialResponse)
	data := make([]*freezonex_openiiot_api.Material, 0)
	for _, v := range wmss {
		storageLocationMaterialService := wms_storagelocationmaterial.DefaultStorageLocationMaterialService()
		storageLocationData, _ := storageLocationMaterialService.GetStorageLocationMaterialMaterialDB(ctx, 0, v.ID)
		var StorageLocationId []string
		var StorageLocation []string

		for _, data := range storageLocationData {
			idStr := strconv.FormatInt(data.StoreLocationID, 10)
			StorageLocationId = append(StorageLocationId, idStr)

			storageLocationService := storagelocation.DefaultStorageLocationService()
			storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, data.StoreLocationID, "", nil, "")
			warehouseService := wms_warehouse.DefaultWmsWarehouseService()
			WarehouseID := storageLocationData[0].WarehouseID
			warehouseData, _ := warehouseService.GetWmsWarehouseDB(ctx, 0, "", common.Int64ToString(WarehouseID), "", "", "", "", "")
			locationame := storageLocationData[0].Name + "-" + warehouseData[0].Name

			StorageLocation = append(StorageLocation, locationame)
		}

		data = append(data, &freezonex_openiiot_api.Material{
			Id:                common.Int64ToString(v.ID), // Converts int64 ID to string using a custom common package function
			ProductCode:       v.ProductCode,
			Name:              v.Name,
			ProductType:       *v.ProductType,
			Unit:              *v.Unit,
			Note:              *v.Note,
			StorageLocationId: StorageLocationId,
			StorageLocation:   StorageLocation,
			CreateTime:        common.GetTimeStringFromTime(&v.CreateTime), // Converts time.Time to string using a custom common package function
			UpdateTime:        common.GetTimeStringFromTime(&v.UpdateTime), // Converts time.Time to string using a custom common package function
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateWmsMaterial will update wms record
func (a *WmsMaterialService) UpdateWmsMaterial(ctx context.Context, req *freezonex_openiiot_api.UpdateMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateMaterialResponse, error) {
	err := a.UpdateWmsMaterialDB(ctx, common.StringToInt64(req.Id), req.ProductCode, req.Name, req.ProductType, req.Unit, req.Note)
	if err != nil {
		logs.Error(ctx, "event=UpdateWmsMaterial error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateMaterialResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteWmsMaterial will delete wms record
func (a *WmsMaterialService) DeleteWmsMaterial(ctx context.Context, req *freezonex_openiiot_api.DeleteMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteMaterialResponse, error) {
	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
	/*err := a.DeleteWmsMaterialUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteWmsMaterial user error=%v", err.Error())
		return nil, err
	}*/

	// Delete wms
	//delete all the users and flows?
	//_, err := a.DeleteWmsMaterialDB(ctx, common.StringToInt64(req.Id))

	err := a.DeleteWmsMaterialDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=DeleteWmsMaterial error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteMaterialResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
