// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProjectStream = "project_streams"

// ProjectStream mapped from table <project_streams>
type ProjectStream struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Color     string    `gorm:"column:color;not null" json:"color"`
	Order_    int32     `gorm:"column:order;not null" json:"order"`
	ProjectID int64     `gorm:"column:project_id;not null" json:"project_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ProjectStream's table name
func (*ProjectStream) TableName() string {
	return TableNameProjectStream
}
