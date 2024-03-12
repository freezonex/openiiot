package wms_stocktaking

import (
	"context"
	"gorm.io/gen/field"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_stocktaking_record"
)

func (a *WmsStocktakingService) AddStocktaking(ctx context.Context, req *freezonex_openiiot_api.AddStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddStocktakingResponse, error) {
	var stocktakingID int64
	var err error // 预先声明错误变量，以便在循环外部使用
	stocktakingID, err = a.AddStocktakingDB(ctx, common.Int64ToString(common.GetUUID()), req.Type, req.Source, "", "")

	for _, b := range req.ShelfRecords {
		for _, d := range b.Inventory {
			stocktakingrecordservie := wms_stocktaking_record.DefaultStocktakingRecordService()
			table := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial
			tx := table.WithContext(ctx).Select(field.ALL).Where(table.StoreLocationID.Eq(common.StringToInt64(b.StorageLocationId))).Where(table.MaterialID.Eq(common.StringToInt64(d.MaterialId)))
			data, err := tx.Find()
			_, err = stocktakingrecordservie.AddStocktakingRecordDB(ctx, common.GetUUID(), stocktakingID, common.StringToInt64(b.StorageLocationId), common.StringToInt64(d.MaterialId), d.Quantity, data[0].Quantity, data[0].Quantity-d.Quantity)
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			logs.Error(ctx, "event=AddStocktaking error=%v", err.Error())
			return nil, err
		}
	}

	resp := new(freezonex_openiiot_api.AddStocktakingResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(stocktakingID)

	return resp, nil
}

// GetStocktaking will get stocktaking record in condition
func (a *WmsStocktakingService) GetStocktaking(ctx context.Context, req *freezonex_openiiot_api.GetStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetStocktakingResponse, error) {
	stocktakings, err := a.GetStocktakingDB(ctx, common.StringToInt64(req.Id), req.RefId, req.Type)

	if err != nil {
		logs.Error(ctx, "event=GetStocktaking error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetStocktakingResponse)
	data := make([]*freezonex_openiiot_api.Stocktaking, 0)

	for _, v := range stocktakings {
		data = append(data, &freezonex_openiiot_api.Stocktaking{
			Id:         common.Int64ToString(v.ID),
			RefId:      v.RefID,
			Type:       v.Type,
			Source:     v.Source,
			Note:       *v.Note,
			Operator:   v.Operator,
			Status:     *v.Status,
			CreateTime: common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime: common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateStocktaking will update stocktaking record
func (a *WmsStocktakingService) UpdateStocktaking(ctx context.Context, req *freezonex_openiiot_api.UpdateStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateStocktakingResponse, error) {
	err := a.UpdateStocktakingDB(ctx, common.StringToInt64(req.Id), req.RefId, req.Type)
	if err != nil {
		logs.Error(ctx, "event=UpdateStocktaking error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateStocktakingResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteStocktaking will delete stocktaking record
func (a *WmsStocktakingService) DeleteStocktaking(ctx context.Context, req *freezonex_openiiot_api.DeleteStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteStocktakingResponse, error) {
	//Delete stocktaking also should delete stocktaking , edge pool, core pool, application pool, flow
	/*err := a.DeleteStocktakingStocktakingDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteStocktaking stocktaking error=%v", err.Error())
		return nil, err
	}*/

	// Delete stocktaking
	err := a.DeleteStocktakingDB(ctx, common.StringToInt64(req.Id))
	if err != nil {
		logs.Error(ctx, "event=DeleteStocktaking error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteStocktakingResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
