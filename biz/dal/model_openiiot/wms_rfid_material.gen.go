// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_openiiot

import (
	"time"
)

const TableNameWmsRfidMaterial = "wms_rfid_material"

// WmsRfidMaterial mapped from table <wms_rfid_material>
type WmsRfidMaterial struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Rfid       string    `gorm:"column:rfid;not null" json:"rfid"`
	MaterialID int64     `gorm:"column:material_id;not null" json:"material_id"`
	Quantity   int32     `gorm:"column:quantity;not null" json:"quantity"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName WmsRfidMaterial's table name
func (*WmsRfidMaterial) TableName() string {
	return TableNameWmsRfidMaterial
}
