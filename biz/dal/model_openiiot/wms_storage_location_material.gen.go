// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_openiiot

const TableNameWmsStorageLocationMaterial = "wms_storage_location_material"

// WmsStorageLocationMaterial mapped from table <wms_storage_location_material>
type WmsStorageLocationMaterial struct {
	ID              int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StoreLocationID int64 `gorm:"column:store_location_id;not null" json:"store_location_id"`
	MaterialID      int64 `gorm:"column:material_id;not null" json:"material_id"`
	Quantity        int32 `gorm:"column:quantity;not null" json:"quantity"`
}

// TableName WmsStorageLocationMaterial's table name
func (*WmsStorageLocationMaterial) TableName() string {
	return TableNameWmsStorageLocationMaterial
}
