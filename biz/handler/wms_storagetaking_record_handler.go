package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/wms_stocktaking_record"
)

type WmsStocktakingRecordHandler struct {
	wmsStocktakingRecordService *wms_stocktaking_record.WmsStocktakingRecordService
}

func NewWmsStocktakingRecordHandler(s *wms_stocktaking_record.WmsStocktakingRecordService) *WmsStocktakingRecordHandler {
	return &WmsStocktakingRecordHandler{wmsStocktakingRecordService: s}
}

func (a *WmsStocktakingRecordHandler) GetWmsStocktakingRecord(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetStocktakingRecordRequest)
	resp, err := a.wmsStocktakingRecordService.GetStocktakingRecord(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsStocktakingRecord error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
