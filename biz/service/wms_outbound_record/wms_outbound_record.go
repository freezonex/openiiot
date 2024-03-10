package wms_outbound_record

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_material"
	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
	"freezonex/openiiot/biz/service/wms_warehouse"
)

//func (a *WmsOutboundRecordService) AddWmsOutboundRecord(ctx context.Context, req *freezonex_openiiot_api.AddInboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddOutboundRecordResponse, error) {
//	wmsID, err := a.AddWmsOutboundRecordDB(ctx, req.InboundId, req.Rfid, req.MaterialId)
//	if err != nil {
//		logs.Error(ctx, "event=AddWmsOutboundRecord error=%v", err.Error())
//		return nil, err
//	}
//
//	resp := new(freezonex_openiiot_api.AddOutboundRecordResponse)
//	resp.BaseResp = middleware.SuccessResponseOK
//	resp.Id = common.Int64ToString(wmsID)
//
//	return resp, nil
//}

// GetWmsOutboundRecord will get wms record in condition
func (a *WmsOutboundRecordService) GetWmsOutboundRecord(ctx context.Context, req *freezonex_openiiot_api.GetOutboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetOutboundRecordResponse, error) {
	wmss, err := a.GetWmsOutboundRecordDB(ctx, common.StringToInt64(req.Id), common.StringToInt64(req.RefId))

	if err != nil {
		logs.Error(ctx, "event=GetWmsOutboundRecord error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetOutboundRecordResponse)
	data := make([]*freezonex_openiiot_api.ShelfInventory, 0)
	for _, v := range wmss {

		materialService := wms_material.DefaultWmsMaterialService()
		materials, err := materialService.GetWmsMaterialDB(ctx, v.MaterialID, "", "", "", "", "")
		if err != nil {
			logs.Error(ctx, "event=GetWmsOutboundRecord error=%v", err.Error())
			return nil, err
		}

		var convertedimaterials []*freezonex_openiiot_api.Inventory

		storageLocationService := storagelocation.DefaultStorageLocationService()
		storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, 0, "", nil, "", v.StockLocationID)
		warehouseService := wms_warehouse.DefaultWmsWarehouseService()
		WarehouseID := storageLocationData[0].WarehouseID
		warehouseData, _ := warehouseService.GetWmsWarehouseDB(ctx, 0, "", common.Int64ToString(WarehouseID), "", "", "", "", "")

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
			Quantity:     v.Quantity,
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

//// UpdateWmsOutboundRecord will update wms record
//func (a *WmsOutboundRecordService) UpdateWmsOutboundRecord(ctx context.Context, req *freezonex_openiiot_api.UpdateOutboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateOutboundRecordResponse, error) {
//	err := a.UpdateWmsOutboundRecordDB(ctx, common.StringToInt64(req.Id), req.InboundId, req.Rfid, common.StringToInt64(req.MaterialId))
//	if err != nil {
//		logs.Error(ctx, "event=UpdateWmsOutboundRecord error=%v", err.Error())
//		return nil, err
//	}
//
//	resp := new(freezonex_openiiot_api.UpdateOutboundRecordResponse)
//	resp.BaseResp = middleware.SuccessResponseOK
//
//	return resp, nil
//}
//
//// DeleteWmsOutboundRecord will delete wms record
//func (a *WmsOutboundRecordService) DeleteWmsOutboundRecord(ctx context.Context, req *freezonex_openiiot_api.DeleteOutboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteOutboundRecordResponse, error) {
//	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
//	/*err := a.DeleteWmsOutboundRecordUserDB(ctx, req.Id)
//	if err != nil {
//		logs.Error(ctx, "event=DeleteWmsOutboundRecord user error=%v", err.Error())
//		return nil, err
//	}*/
//
//	// Delete wms
//	//delete all the users and flows?
//	//_, err := a.DeleteWmsOutboundRecordDB(ctx, common.StringToInt64(req.Id))
//
//	err := a.DeleteWmsOutboundRecordDB(ctx, common.StringToInt64(req.Id))
//
//	if err != nil {
//		logs.Error(ctx, "event=DeleteWmsOutboundRecord error=%v", err.Error())
//		return nil, err
//	}
//	resp := new(freezonex_openiiot_api.DeleteOutboundRecordResponse)
//	resp.BaseResp = middleware.SuccessResponseOK
//
//	return resp, nil
//}
