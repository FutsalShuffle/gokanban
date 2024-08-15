// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGrade = "grades"

// Grade mapped from table <grades>
type Grade struct {
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GradeDirectionID int64     `gorm:"column:grade_direction_id;not null" json:"grade_direction_id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Slug             string    `gorm:"column:slug;not null" json:"slug"`
	Order_           int32     `gorm:"column:order;not null;default:1" json:"order"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Grade's table name
func (*Grade) TableName() string {
	return TableNameGrade
}
