package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/wms_outbound_record"
)

type WmsOutboundRecordHandler struct {
	wmsOutboundRecordService *wms_outbound_record.WmsOutboundRecordService
}

func NewWmsOutboundRecordHandler(s *wms_outbound_record.WmsOutboundRecordService) *WmsOutboundRecordHandler {
	return &WmsOutboundRecordHandler{wmsOutboundRecordService: s}
}

//func (a *WmsOutboundRecordHandler) AddWmsOutboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddOutboundRecordRequest)
//	resp, err := a.wmsOutboundRecordService.AddWmsOutboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=AddWmsOutboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}

func (a *WmsOutboundRecordHandler) GetWmsOutboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetOutboundRecordRequest)
	resp, err := a.wmsOutboundRecordService.GetWmsOutboundRecord(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsOutboundRecord error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

//func (a *WmsOutboundRecordHandler) UpdateWmsOutboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateOutboundRecordRequest)
//	resp, err := a.wmsOutboundRecordService.UpdateWmsOutboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=UpdateWmsOutboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}

//func (a *WmsOutboundRecordHandler) DeleteWmsOutboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteOutboundRecordRequest)
//	resp, err := a.wmsOutboundRecordService.DeleteWmsOutboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=DeleteWmsOutboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}
