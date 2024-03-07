package wms_rfid_material

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func (a *WmsRfidRfidMaterialService) AddWmsRfidMaterial(ctx context.Context, req *freezonex_openiiot_api.AddRfidMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddRfidMaterialResponse, error) {
	wmsID, err := a.AddWmsRfidMaterialDB(ctx, req.Rfid, common.StringToInt64(req.MaterialId), req.Quantity)
	if err != nil {
		logs.Error(ctx, "event=AddWmsRfidMaterial error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddRfidMaterialResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(wmsID)

	return resp, nil
}

// GetWmsRfidMaterial will get wms record in condition
func (a *WmsRfidRfidMaterialService) GetWmsRfidMaterial(ctx context.Context, req *freezonex_openiiot_api.GetRfidMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetRfidMaterialResponse, error) {
	wmss, err := a.GetWmsRfidMaterialDB(ctx, 0, req.Rfid, common.StringToInt64(req.MaterialId), req.Quantity)

	if err != nil {
		logs.Error(ctx, "event=GetWmsRfidMaterial error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetRfidMaterialResponse)
	data := make([]*freezonex_openiiot_api.RfidMaterial, 0)
	for _, v := range wmss {
		table := a.db.DBOpeniiotQuery.WmsMaterial
		existRecord, _ := table.WithContext(ctx).Where(table.ID.Eq(v.MaterialID)).First()

		data = append(data, &freezonex_openiiot_api.RfidMaterial{
			Id:           common.Int64ToString(v.ID), // Converts int64 ID to string using a custom common package function
			Rfid:         v.Rfid,
			MaterialId:   common.Int64ToString(v.MaterialID),
			Quantity:     v.Quantity,
			MaterialName: existRecord.Name,
			CreateTime:   common.GetTimeStringFromTime(&v.CreateTime), // Converts time.Time to string using a custom common package function
			UpdateTime:   common.GetTimeStringFromTime(&v.UpdateTime), // Converts time.Time to string using a custom common package function
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateWmsRfidMaterial will update wms record
func (a *WmsRfidRfidMaterialService) UpdateWmsRfidMaterial(ctx context.Context, req *freezonex_openiiot_api.UpdateRfidMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateRfidMaterialResponse, error) {
	err := a.UpdateWmsRfidMaterialDB(ctx, common.StringToInt64(req.Id), req.Rfid, common.StringToInt64(req.MaterialId), req.Quantity)
	if err != nil {
		logs.Error(ctx, "event=UpdateWmsRfidMaterial error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateRfidMaterialResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteWmsRfidMaterial will delete wms record
func (a *WmsRfidRfidMaterialService) DeleteWmsRfidMaterial(ctx context.Context, req *freezonex_openiiot_api.DeleteRfidMaterialRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteRfidMaterialResponse, error) {
	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
	/*err := a.DeleteWmsRfidMaterialUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteWmsRfidMaterial user error=%v", err.Error())
		return nil, err
	}*/

	// Delete wms
	//delete all the users and flows?
	//_, err := a.DeleteWmsRfidMaterialDB(ctx, common.StringToInt64(req.Id))

	err := a.DeleteWmsRfidMaterialDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=DeleteWmsRfidMaterial error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteRfidMaterialResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
