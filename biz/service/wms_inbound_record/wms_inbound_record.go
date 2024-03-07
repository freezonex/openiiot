package wms_inbound_record

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func (a *WmsInboundRecordService) AddWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.AddInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddInboundRecordResponse, error) {
	wmsID, err := a.AddWmsInboundRecordDB(ctx, req.InboundId, req.Rfid, req.MaterialId)
	if err != nil {
		logs.Error(ctx, "event=AddWmsInboundRecord error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddInboundRecordResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(wmsID)

	return resp, nil
}

// GetWmsInboundRecord will get wms record in condition
func (a *WmsInboundRecordService) GetWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.GetInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetInboundRecordResponse, error) {
	wmss, err := a.GetWmsInboundRecordDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=GetWmsInboundRecord error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetInboundRecordResponse)
	data := make([]*freezonex_openiiot_api.InboundRecord, 0)
	for _, v := range wmss {
		data = append(data, &freezonex_openiiot_api.InboundRecord{
			Id:         common.Int64ToString(v.ID), // Converts int64 ID to string using a custom common package function
			InboundId:  common.Int64ToString(v.InboundID),
			Rfid:       *v.Rfid,
			MaterialId: common.Int64ToString(v.MaterialID),
			CreateTime: common.GetTimeStringFromTime(&v.CreateTime), // Converts time.Time to string using a custom common package function
			UpdateTime: common.GetTimeStringFromTime(&v.UpdateTime), // Converts time.Time to string using a custom common package function
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateWmsInboundRecord will update wms record
func (a *WmsInboundRecordService) UpdateWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.UpdateInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateInboundRecordResponse, error) {
	err := a.UpdateWmsInboundRecordDB(ctx, common.StringToInt64(req.Id), req.InboundId, req.Rfid, common.StringToInt64(req.MaterialId))
	if err != nil {
		logs.Error(ctx, "event=UpdateWmsInboundRecord error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateInboundRecordResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteWmsInboundRecord will delete wms record
func (a *WmsInboundRecordService) DeleteWmsInboundRecord(ctx context.Context, req *freezonex_openiiot_api.DeleteInboundRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteInboundRecordResponse, error) {
	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
	/*err := a.DeleteWmsInboundRecordUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteWmsInboundRecord user error=%v", err.Error())
		return nil, err
	}*/

	// Delete wms
	//delete all the users and flows?
	//_, err := a.DeleteWmsInboundRecordDB(ctx, common.StringToInt64(req.Id))

	err := a.DeleteWmsInboundRecordDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=DeleteWmsInboundRecord error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteInboundRecordResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}