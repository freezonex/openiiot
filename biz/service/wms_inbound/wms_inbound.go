package wms_inbound

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
	"freezonex/openiiot/biz/service/wms_warehouse"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func (a *WmsInboundService) AddWmsInbound(ctx context.Context, req *freezonex_openiiot_api.AddInboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddInboundResponse, error) {
	wmsID, err := a.AddWmsInboundDB(ctx, common.StringToInt64(req.StorageLocationId), req.Type, req.MaterialName, req.Source, req.Rfids, common.ConvertStringSliceToInt64Slice(req.MaterialIds))
	if err != nil {
		logs.Error(ctx, "event=AddWmsInbound error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddInboundResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(wmsID)

	return resp, nil
}

// GetWmsInbound will get wms record in condition
func (a *WmsInboundService) GetWmsInbound(ctx context.Context, req *freezonex_openiiot_api.GetInboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetInboundResponse, error) {
	wmss, err := a.GetWmsInboundDB(ctx, common.StringToInt64(req.Id), req.RefId)

	if err != nil {
		logs.Error(ctx, "event=GetWmsInbound error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetInboundResponse)
	data := make([]*freezonex_openiiot_api.Inbound, 0)
	for _, v := range wmss {
		storageLocationService := storagelocation.DefaultStorageLocationService()
		storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, v.StorageLocationID, "")
		warehouseService := wms_warehouse.DefaultWmsWarehouseService()
		WarehouseID := storageLocationData[0].WarehouseID
		warehouseData, _ := warehouseService.GetWmsWarehouseDB(ctx, 0, "", common.Int64ToString(WarehouseID), "", "", "", "", "")
		locationame := storageLocationData[0].Name + "-" + warehouseData[0].Name

		data = append(data, &freezonex_openiiot_api.Inbound{
			Id:                common.Int64ToString(v.ID), // Converts int64 ID to string using a custom common package function
			RefId:             v.RefID,
			Type:              v.Type,
			StorageLocationId: common.Int64ToString(storageLocationData[0].ID),
			StorageLocation:   locationame,
			MaterialName:      v.MaterialName,
			Operator:          v.Operator,
			CreateTime:        common.GetTimeStringFromTime(&v.CreateTime), // Converts time.Time to string using a custom common package function
			UpdateTime:        common.GetTimeStringFromTime(&v.UpdateTime), // Converts time.Time to string using a custom common package function
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateWmsInbound will update wms record
func (a *WmsInboundService) UpdateWmsInbound(ctx context.Context, req *freezonex_openiiot_api.UpdateInboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateInboundResponse, error) {
	err := a.UpdateWmsInboundDB(ctx, common.StringToInt64(req.Id), req.RefId, req.Type, req.StorageLocationId, req.MaterialName)
	if err != nil {
		logs.Error(ctx, "event=UpdateWmsInbound error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateInboundResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteWmsInbound will delete wms record
func (a *WmsInboundService) DeleteWmsInbound(ctx context.Context, req *freezonex_openiiot_api.DeleteInboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteInboundResponse, error) {
	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
	/*err := a.DeleteWmsInboundUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteWmsInbound user error=%v", err.Error())
		return nil, err
	}*/

	// Delete wms
	//delete all the users and flows?
	//_, err := a.DeleteWmsInboundDB(ctx, common.StringToInt64(req.Id))

	err := a.DeleteWmsInboundDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=DeleteWmsInbound error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteInboundResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
