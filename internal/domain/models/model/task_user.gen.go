// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTaskUser = "task_user"

// TaskUser mapped from table <task_user>
type TaskUser struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID int64 `gorm:"column:user_id;not null" json:"user_id"`
	TaskID int64 `gorm:"column:task_id;not null" json:"task_id"`
}

// TableName TaskUser's table name
func (*TaskUser) TableName() string {
	return TableNameTaskUser
}
