// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameContainFavorite = "contain_favorites"

// ContainFavorite mapped from table <contain_favorites>
type ContainFavorite struct {
	ID                  int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID              int64  `gorm:"column:user_id;not null" json:"user_id"`
	ContainFavoriteID   int64  `gorm:"column:contain_favorite_id;not null" json:"contain_favorite_id"`
	ContainFavoriteType string `gorm:"column:contain_favorite_type;not null" json:"contain_favorite_type"`
}

// TableName ContainFavorite's table name
func (*ContainFavorite) TableName() string {
	return TableNameContainFavorite
}
