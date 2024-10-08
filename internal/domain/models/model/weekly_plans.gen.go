// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameWeeklyPlan = "weekly_plans"

// WeeklyPlan mapped from table <weekly_plans>
type WeeklyPlan struct {
	ID               int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProjectID        int64          `gorm:"column:project_id;not null" json:"project_id"`
	UserID           int64          `gorm:"column:user_id;not null" json:"user_id"`
	GradeDirectionID int64          `gorm:"column:grade_direction_id;not null" json:"grade_direction_id"`
	HoursNeeded      float64        `gorm:"column:hours_needed;not null" json:"hours_needed"`
	HoursPlanned     float64        `gorm:"column:hours_planned;not null" json:"hours_planned"`
	Begin            time.Time      `gorm:"column:begin;not null" json:"begin"`
	End              time.Time      `gorm:"column:end;not null" json:"end"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	ManagerID        int64          `gorm:"column:manager_id" json:"manager_id"`
}

// TableName WeeklyPlan's table name
func (*WeeklyPlan) TableName() string {
	return TableNameWeeklyPlan
}
