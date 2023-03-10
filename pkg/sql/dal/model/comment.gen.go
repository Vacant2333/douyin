// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID         uint      `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	UserID     uint      `gorm:"column:user_id;type:int;not null" json:"user_id"`
	VideoID    uint      `gorm:"column:video_id;type:int;not null" json:"video_id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	Removed    uint      `gorm:"column:removed;type:tinyint;not null" json:"removed"`
	Deleted    uint      `gorm:"column:deleted;type:tinyint;not null" json:"deleted"`
	Content    string    `gorm:"column:content;type:text;not null" json:"content"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
