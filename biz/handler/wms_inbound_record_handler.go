package handler

import (
	"freezonex/openiiot/biz/service/wms_inbound_record"
)

type WmsInboundRecordHandler struct {
	wmsInboundRecordService *wms_inbound_record.WmsInboundRecordService
}

func NewWmsInboundRecordHandler(s *wms_inbound_record.WmsInboundRecordService) *WmsInboundRecordHandler {
	return &WmsInboundRecordHandler{wmsInboundRecordService: s}
}

//func (a *WmsInboundRecordHandler) AddWmsInboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddInboundRecordRequest)
//	resp, err := a.wmsInboundRecordService.AddWmsInboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=AddWmsInboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}

//func (a *WmsInboundRecordHandler) GetWmsInboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetInboundRecordRequest)
//	resp, err := a.wmsInboundRecordService.GetWmsInboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=GetWmsInboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}

//func (a *WmsInboundRecordHandler) UpdateWmsInboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateInboundRecordRequest)
//	resp, err := a.wmsInboundRecordService.UpdateWmsInboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=UpdateWmsInboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}

//func (a *WmsInboundRecordHandler) DeleteWmsInboundRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
//	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteInboundRecordRequest)
//	resp, err := a.wmsInboundRecordService.DeleteWmsInboundRecord(ctx, req, c)
//	if err != nil {
//		logs.CtxErrorf(ctx, "event=DeleteWmsInboundRecord error=%v", err)
//		return middleware.ErrorResp(http.StatusInternalServerError, err)
//	}
//	return resp
//}
