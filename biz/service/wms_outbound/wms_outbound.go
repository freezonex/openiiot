package wms_outbound

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_storage_location"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"strings"
)

func (a *WmsOutboundService) AddWmsOutbound(ctx context.Context, req *freezonex_openiiot_api.AddOutboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddOutboundResponse, error) {
	wmsID, err := a.AddWmsOutboundDB(ctx, req.Type, req.Source, req.ShelfRecords, req.Note, req.Status)
	if err != nil {
		logs.Error(ctx, "event=AddWmsOutbound error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddOutboundResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(wmsID)

	return resp, nil
}

// GetWmsOutbound will get wms record in condition
func (a *WmsOutboundService) GetWmsOutbound(ctx context.Context, req *freezonex_openiiot_api.GetOutboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetOutboundResponse, error) {
	wmss, err := a.GetWmsOutboundDB(ctx, common.StringToInt64(req.Id), req.RefId)

	if err != nil {
		logs.Error(ctx, "event=GetWmsOutbound error=%v", err.Error())
		return nil, err
	}

	var Names []string // 初始化一个空的字符串切片来存储id
	resp := new(freezonex_openiiot_api.GetOutboundResponse)
	data := make([]*freezonex_openiiot_api.Outbound, 0)
	for _, v := range wmss {

		for _, b := range strings.Split("", ",") {

			storageLocationService := wms_storage_location.DefaultStorageLocationService()
			storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, common.StringToInt64(b), "", nil, "")
			Names = append(Names, storageLocationData[0].Name) // 将每个Id添加到切片中
		}

		data = append(data, &freezonex_openiiot_api.Outbound{
			Id:         common.Int64ToString(v.ID), // Converts int64 ID to string using a custom common package function
			RefId:      v.RefID,
			Type:       v.Type,
			Operator:   v.Operator,
			Source:     v.Source,
			Status:     *v.Status,
			Note:       *v.Note,
			CreateTime: common.GetTimeStringFromTime(&v.CreateTime), // Converts time.Time to string using a custom common package function
			UpdateTime: common.GetTimeStringFromTime(&v.UpdateTime), // Converts time.Time to string using a custom common package function
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateWmsOutbound will update wms record
func (a *WmsOutboundService) UpdateWmsOutbound(ctx context.Context, req *freezonex_openiiot_api.UpdateOutboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateOutboundResponse, error) {
	err := a.UpdateWmsOutboundDB(ctx, common.StringToInt64(req.Id), req.RefId, req.Type)
	if err != nil {
		logs.Error(ctx, "event=UpdateWmsOutbound error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateOutboundResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteWmsOutbound will delete wms record
func (a *WmsOutboundService) DeleteWmsOutbound(ctx context.Context, req *freezonex_openiiot_api.DeleteOutboundRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteOutboundResponse, error) {
	//Delete wms also should delete wms user, edge pool, core pool, application pool, flow
	/*err := a.DeleteWmsOutboundUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteWmsOutbound user error=%v", err.Error())
		return nil, err
	}*/

	// Delete wms
	//delete all the users and flows?
	//_, err := a.DeleteWmsOutboundDB(ctx, common.StringToInt64(req.Id))

	err := a.DeleteWmsOutboundDB(ctx, common.StringToInt64(req.Id))

	if err != nil {
		logs.Error(ctx, "event=DeleteWmsOutbound error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteOutboundResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
