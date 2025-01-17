// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameToggle = "toggle"

// Toggle mapped from table <toggle>
type Toggle struct {
	ID                     int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProjectID              int64     `gorm:"column:project_id;not null" json:"project_id"`
	Enable                 bool      `gorm:"column:enable;not null" json:"enable"`
	Title                  string    `gorm:"column:title;not null" json:"title"`
	Key                    string    `gorm:"column:key;not null" json:"key"`
	Description            string    `gorm:"column:description;not null" json:"description"`
	ValueType              int8      `gorm:"column:value_type;not null" json:"value_type"`
	DefaultValue           int64     `gorm:"column:default_value;not null" json:"default_value"`
	ReturnValueWhenDisable int64     `gorm:"column:return_value_when_disable;not null" json:"return_value_when_disable"`
	CreateTime             time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime             time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName Toggle's table name
func (*Toggle) TableName() string {
	return TableNameToggle
}
