// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameToggleValue = "toggle_value"

// ToggleValue mapped from table <toggle_value>
type ToggleValue struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ToggleID    int64     `gorm:"column:toggle_id;not null" json:"toggle_id"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Value       string    `gorm:"column:value;not null" json:"value"`
	Description string    `gorm:"column:description;not null" json:"description"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName ToggleValue's table name
func (*ToggleValue) TableName() string {
	return TableNameToggleValue
}
