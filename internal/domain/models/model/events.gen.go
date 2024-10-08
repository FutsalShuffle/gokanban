// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEvent = "events"

// Event mapped from table <events>
type Event struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaskID    int64     `gorm:"column:task_id;not null" json:"task_id"`
	DateStart time.Time `gorm:"column:date_start" json:"date_start"`
	DateEnd   time.Time `gorm:"column:date_end" json:"date_end"`
	UserID    int64     `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Event's table name
func (*Event) TableName() string {
	return TableNameEvent
}
