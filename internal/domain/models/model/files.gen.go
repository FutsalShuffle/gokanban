// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFile = "files"

// File mapped from table <files>
type File struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OriginalName string    `gorm:"column:original_name;not null" json:"original_name"`
	Path         string    `gorm:"column:path;not null" json:"path"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName File's table name
func (*File) TableName() string {
	return TableNameFile
}
