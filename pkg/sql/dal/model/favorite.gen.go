// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFavorite = "favorite"

// Favorite mapped from table <favorite>
type Favorite struct {
	ID      uint `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	VideoID uint `gorm:"column:video_id;type:int;not null" json:"video_id"`
	UserID  uint `gorm:"column:user_id;type:int;not null" json:"user_id"`
	Removed uint `gorm:"column:removed;type:tinyint;not null" json:"removed"`
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}
