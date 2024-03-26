// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArchivedViewHistory = "archived_view_history"

// ArchivedViewHistory 浏览历史记录
type ArchivedViewHistory struct {
	ID         int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime *time.Time     `gorm:"column:create_time;type:datetime(3);autoCreateTime" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;type:datetime(3);autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime(3)" json:"delete_time"`
}

// TableName ArchivedViewHistory's table name
func (*ArchivedViewHistory) TableName() string {
	return TableNameArchivedViewHistory
}
