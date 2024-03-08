package wms_inbound_record

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_material"
	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
	"freezonex/openiiot/biz/service/wms_storagelocationmaterial"
	"freezonex/openiiot/biz/service/wms_warehouse"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

//func (a *WmsInboundRecordService) AddWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.AddInboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddInboundRecordResponse, error) {
//	wmsID, err := a.AddWmsInboundRecordDB(ctx, req.InboundId, req.Rfid, req.MaterialId)
//	if err != nil {
//		logs.Error(ctx, "event=AddWmsInboundRecord error=%v", err.Error())
//		return nil, err
//	}
//
//	resp := new(freezonex_openiiot_api.AddInboundRecordResponse)
//	resp.BaseResp = middleware.SuccessResponseOK
//	resp.Id = common.Int64ToString(wmsID)
//
//	return resp, nil
//}

// GetWmsInboundRecord will get wms record in condition
func (a *WmsInboundRecordService) GetWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.GetInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetInboundRecordResponse, error) {
	wmss, err := a.GetWmsInboundRecordDB(ctx, common.StringToInt64(req.Id), common.StringToInt64(req.RefId))

	if err != nil {
		logs.Error(ctx, "event=GetWmsInboundRecord error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetInboundRecordResponse)
	data := make([]*freezonex_openiiot_api.ShelfInventory, 0)
	for _, v := range wmss {

		materialService := wms_material.DefaultWmsMaterialService()
		materials, err := materialService.GetWmsMaterialDB(ctx, v.MaterialID, "", "", "", "", "")
		if err != nil {
			logs.Error(ctx, "event=GetWmsInboundRecord error=%v", err.Error())
			return nil, err
		}

		var convertedimaterials []*freezonex_openiiot_api.Inventory

		storageLocationService := storagelocation.DefaultStorageLocationService()
		storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, v.StockLocationID, "", nil, "")
		warehouseService := wms_warehouse.DefaultWmsWarehouseService()
		WarehouseID := storageLocationData[0].WarehouseID
		warehouseData, _ := warehouseService.GetWmsWarehouseDB(ctx, 0, "", common.Int64ToString(WarehouseID), "", "", "", "", "")
		storageLocationMaterialService := wms_storagelocationmaterial.DefaultStorageLocationMaterialService()
		storageLocationMaterialData, _ := storageLocationMaterialService.GetStorageLocationMaterialDB(ctx, 0, v.StockLocationID, v.MaterialID, 0)
		locationame := warehouseData[0].Name + "-" + storageLocationData[0].Name

		//u := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial
		//e := a.db.DBOpeniiotQuery.WmsStorageLocation
		//w:=a.db.DBOpeniiotQuery.WmsWarehouse
		//getid :=  v.MaterialID
		//err1 := w.WithContext(ctx).Select(e.Name.As("WmsStorageLocationName"), w.Name.As("WmsWarehouse"), u.Quantity).LeftJoin(e, e.WarehouseID.EqCol(w.ID))
		//existRecord1,err:= err1.Where(u.MaterialID.Eq(getid)).First()
		//
		//err2 := e.WithContext(ctx).Select(e.Name.As("WmsStorageLocationName"), w.Name.As("WmsWarehouse"), u.Quantity).LeftJoin(u, u.StoreLocationID.EqCol(e.ID))
		//existRecord2,err:= err2.Where(u.MaterialID.Eq(getid)).First()
		//
		//err3 := u.WithContext(ctx)
		//existRecord3,err:= err3.Where(u.MaterialID.Eq(getid)).First()
		//locationame := existRecord1.Name+"-"+existRecord2.Name

		convertedimaterial := &freezonex_openiiot_api.Inventory{
			MaterialName: materials[0].Name,
			MaterialId:   common.Int64ToString(v.MaterialID),
			Quantity:     storageLocationMaterialData[0].Quantity,
		}
		convertedimaterials = append(convertedimaterials, convertedimaterial)

		data = append(data, &freezonex_openiiot_api.ShelfInventory{
			StorageLocationId: common.Int64ToString(v.StockLocationID),
			Inventory:         convertedimaterials,
			StorageLocation:   locationame,
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

//// UpdateWmsInboundRecord will update wms record
//func (a *WmsInboundRecordService) UpdateWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.UpdateInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateInboundRecordResponse, error) {
//	err := a.UpdateWmsInboundRecordDB(ctx, common.StringToInt64(req.Id), req.InboundId, req.Rfid, common.StringToInt64(req.MaterialId))
//	if err != nil {
//		logs.Error(ctx, "event=UpdateWmsInboundRecord error=%v", err.Error())
//		return nil, err
//	}
//
//	resp := new(freezonex_openiiot_api.UpdateInboundRecordResponse)
//	resp.BaseResp = middleware.SuccessResponseOK
//
//	return resp, nil
//}
//
//// DeleteWmsInboundRecord will delete wms record
//func (a *WmsInboundRecordService) DeleteWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.DeleteInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteInboundRecordResponse, error) {
//	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
//	/*err := a.DeleteWmsInboundRecordUserDB(ctx, req.Id)
//	if err != nil {
//		logs.Error(ctx, "event=DeleteWmsInboundRecord user error=%v", err.Error())
//		return nil, err
//	}*/
//
//	// Delete wms
//	//delete all the users and flows?
//	//_, err := a.DeleteWmsInboundRecordDB(ctx, common.StringToInt64(req.Id))
//
//	err := a.DeleteWmsInboundRecordDB(ctx, common.StringToInt64(req.Id))
//
//	if err != nil {
//		logs.Error(ctx, "event=DeleteWmsInboundRecord error=%v", err.Error())
//		return nil, err
//	}
//	resp := new(freezonex_openiiot_api.DeleteInboundRecordResponse)
//	resp.BaseResp = middleware.SuccessResponseOK
//
//	return resp, nil
//}
