package wms_stocktaking_record

import (
	"context"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/dal/model_openiiot"
)

// GetStocktakingRecordDB will get storagelocation record from the DB in condition
func (a *WmsStocktakingRecordService) GetStocktakingRecordDB(ctx context.Context, ID int64, StocktakingID int64, StockLocationID int64, MaterialID int64, Quantity int32, StockQuantity int32, Discrepancy int32) ([]*model_openiiot.WmsStocktakingRecord, error) {

	table := a.db.DBOpeniiotQuery.WmsStocktakingRecord
	tx := table.WithContext(ctx).Select(field.ALL)
	if ID != 0 {
		tx = tx.Where(table.ID.Eq(ID))
	}
	if StocktakingID != 0 {
		tx = tx.Where(table.StocktakingID.Eq(StocktakingID))
	}
	if StockLocationID != 0 {
		tx = tx.Where(table.StockLocationID.Eq(StockLocationID))
	}
	if MaterialID != 0 {
		tx = tx.Where(table.MaterialID.Eq(MaterialID))
	}
	if Quantity != 0 {
		tx = tx.Where(table.Quantity.Eq(Quantity))
	}
	if StockQuantity != 0 {
		tx = tx.Where(table.StockQuantity.Eq(StockQuantity))
	}
	if Discrepancy != 0 {
		tx = tx.Where(table.Discrepancy.Eq(Discrepancy))
	}

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}