package wms_stocktaking_record

import (
	"context"
	"gorm.io/gen/field"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
)

// GetStocktakingRecord will get storagelocation record in condition
func (a *WmsStocktakingRecordService) GetStocktakingRecord(ctx context.Context, req *freezonex_openiiot_api.GetStocktakingRecordRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetStocktakingRecordResponse, error) {
	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx).Select(field.ALL)
	if req.RefId != "" {
		tx = tx.Where(table.RefID.Eq(req.RefId))
	}
	if req.Id != "" {
		tx = tx.Where(table.ID.Eq(common.StringToInt64(req.Id)))
	}
	wmsStocktaking, err := tx.Find()
	storagelocations, err := a.GetStocktakingRecordDB(ctx, 0, wmsStocktaking[0].ID, 0, 0, 0, 0, 0)

	if err != nil {
		logs.Error(ctx, "event=GetStocktakingRecord error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetStocktakingRecordResponse)
	data := make([]*freezonex_openiiot_api.ShelfInventory, 0)
	data1 := make([]*freezonex_openiiot_api.Inventory, 0)

	for _, v := range storagelocations {
		table := a.db.DBOpeniiotQuery.WmsStorageLocation
		existRecord, _ := table.WithContext(ctx).Select(table.Name).Where(table.ID.Eq(v.StockLocationID)).First()

		table1 := a.db.DBOpeniiotQuery.WmsMaterial
		existRecord1, _ := table1.WithContext(ctx).Select(table1.Name).Where(table1.ID.Eq(v.MaterialID)).First()

		data1 = append(data1, &freezonex_openiiot_api.Inventory{
			Rfid:          "",
			MaterialId:    common.Int64ToString(v.MaterialID),
			Quantity:      *v.Quantity,
			MaterialName:  existRecord1.Name,
			StockQuantity: *v.StockQuantity,
			Discrepancy:   *v.Discrepancy,
		})

		data = append(data, &freezonex_openiiot_api.ShelfInventory{
			StorageLocationId: common.Int64ToString(v.StocktakingID),
			Inventory:         data1,
			StorageLocation:   existRecord.Name,
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
