// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOutgoingWebhook = "outgoing_webhooks"

// OutgoingWebhook mapped from table <outgoing_webhooks>
type OutgoingWebhook struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsActive  bool      `gorm:"column:is_active;not null;default:1" json:"is_active"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	URL       string    `gorm:"column:url;not null" json:"url"`
	SecretKey string    `gorm:"column:secret_key;not null" json:"secret_key"`
	Event     string    `gorm:"column:event;not null" json:"event"`
	Timeout   int32     `gorm:"column:timeout;not null;default:5" json:"timeout"`
}

// TableName OutgoingWebhook's table name
func (*OutgoingWebhook) TableName() string {
	return TableNameOutgoingWebhook
}
