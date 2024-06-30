// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameReqGroupOption = "req_group_option"

// ReqGroupOption mapped from table <req_group_option>
type ReqGroupOption struct {
	OptionID       int64     `gorm:"column:option_id;primaryKey;autoIncrement:true" json:"option_id"`
	GroupID        int64     `gorm:"column:group_id;not null" json:"group_id"`
	ListNum        int64     `gorm:"column:list_num;not null" json:"list_num"`
	AttributeType  int8      `gorm:"column:attribute_type;not null" json:"attribute_type"`
	AttributeName  string    `gorm:"column:attribute_name;not null" json:"attribute_name"`
	OperationType  int8      `gorm:"column:operation_type;not null" json:"operation_type"`
	AttributeValue string    `gorm:"column:attribute_value;not null" json:"attribute_value"`
	CreateTime     time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName ReqGroupOption's table name
func (*ReqGroupOption) TableName() string {
	return TableNameReqGroupOption
}