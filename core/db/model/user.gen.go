// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username string `gorm:"column:username;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
