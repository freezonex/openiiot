// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_openiiot

import (
	"time"
)

const TableNameWmsStocktaking = "wms_stocktaking"

// WmsStocktaking mapped from table <wms_stocktaking>
type WmsStocktaking struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RefID           string    `gorm:"column:ref_id;not null;uniqueIndex:ref_id,priority:1" json:"ref_id"`
	Type            string    `gorm:"column:type;not null" json:"type"`
	StorageLocation int64     `gorm:"column:storage_location;not null" json:"storage_location"`
	Operator        string    `gorm:"column:operator;not null" json:"operator"`
	Result          string    `gorm:"column:result;not null" json:"result"`
	UpdateTime      time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName WmsStocktaking's table name
func (*WmsStocktaking) TableName() string {
	return TableNameWmsStocktaking
}
