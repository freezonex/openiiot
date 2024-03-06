// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_openiiot

import (
	"time"
)

const TableNameWmsInbound = "wms_inbound"

// WmsInbound mapped from table <wms_inbound>
type WmsInbound struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RefID             string    `gorm:"column:ref_id;not null;uniqueIndex:ref_id,priority:1" json:"ref_id"`
	Type              string    `gorm:"column:type;not null" json:"type"`
	StorageLocationID int64     `gorm:"column:storage_location_id;not null" json:"storage_location_id"`
	MaterialName      string    `gorm:"column:material_name;not null" json:"material_name"`
	Source            string    `gorm:"column:source;not null" json:"source"`
	Operator          string    `gorm:"column:operator;not null" json:"operator"`
	UpdateTime        time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName WmsInbound's table name
func (*WmsInbound) TableName() string {
	return TableNameWmsInbound
}
