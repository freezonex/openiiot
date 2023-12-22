// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_openiiot

import (
	"time"
)

const TableNameCore = "core"

// Core mapped from table <core>
type Core struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name;not null;uniqueIndex:name,priority:1" json:"name"`
	Description *string   `gorm:"column:description" json:"description"`
	TenantID    int64     `gorm:"column:tenant_id;not null" json:"tenant_id"`
	URL         string    `gorm:"column:url;not null" json:"url"`
	Type        string    `gorm:"column:type;not null" json:"type"`
	Source      *string   `gorm:"column:source" json:"source"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName Core's table name
func (*Core) TableName() string {
	return TableNameCore
}
