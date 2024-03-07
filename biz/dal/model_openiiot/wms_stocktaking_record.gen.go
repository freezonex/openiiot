// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_openiiot

const TableNameWmsStocktakingRecord = "wms_stocktaking_record"

// WmsStocktakingRecord mapped from table <wms_stocktaking_record>
type WmsStocktakingRecord struct {
	ID              int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StocktakingID   int64  `gorm:"column:stocktaking_id;not null" json:"stocktaking_id"`
	StockLocationID int64  `gorm:"column:stock_location_id;not null" json:"stock_location_id"`
	MaterialID      int64  `gorm:"column:material_id;not null" json:"material_id"`
	Quantity        *int32 `gorm:"column:quantity" json:"quantity"`
	StockQuantity   *int32 `gorm:"column:stock_quantity" json:"stock_quantity"`
	Discrepancy     *int32 `gorm:"column:discrepancy" json:"discrepancy"`
}

// TableName WmsStocktakingRecord's table name
func (*WmsStocktakingRecord) TableName() string {
	return TableNameWmsStocktakingRecord
}