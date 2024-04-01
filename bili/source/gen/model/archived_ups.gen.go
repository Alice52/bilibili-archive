// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArchivedUp = "archived_ups"

// ArchivedUp 关注的UP
type ArchivedUp struct {
	CreateTime     *time.Time      `gorm:"column:create_time;type:datetime(3);autoCreateTime" json:"create_time"`
	UpdateTime     *time.Time      `gorm:"column:update_time;type:datetime(3);autoUpdateTime" json:"update_time"`
	DeleteTime     gorm.DeletedAt  `gorm:"column:delete_time;type:datetime(3)" json:"delete_time"`
	TagID          int64           `gorm:"column:tag_id;type:bigint;primaryKey;comment:my group" json:"tag_id"`     // my group
	Sign           *string         `gorm:"column:sign;type:varchar(2048);comment:up desc" json:"sign"`              // up desc
	Uname          *string         `gorm:"column:uname;type:varchar(128);comment:up name" json:"uname"`             // up name
	Mid            int64           `gorm:"column:mid;type:bigint;primaryKey;comment:up uid" json:"mid"`             // up uid
	Level          *string         `gorm:"column:level;type:varchar(3);comment:up level" json:"level"`              // up level
	Rank           *string         `gorm:"column:rank;type:varchar(30);comment:up rank" json:"rank"`                // up rank
	Following      *string         `gorm:"column:following;type:varchar(30);comment:up following" json:"following"` // up following
	Follower       *string         `gorm:"column:follower;type:varchar(30);comment:up follower" json:"follower"`    // up follower
	View           *string         `gorm:"column:view;type:varchar(30);comment:up view" json:"view"`                // up view
	Likes          *string         `gorm:"column:likes;type:varchar(30);comment:up likes" json:"likes"`             // up likes
	Resp           *string         `gorm:"column:resp;type:json" json:"resp"`
	Video          *string         `gorm:"column:video;type:varchar(30);comment:up video count" json:"video"` // up video count
	ArchivedUpsTag *ArchivedUpsTag `gorm:"foreignKey:tag_id" json:"archived_ups_tag"`
}

// TableName ArchivedUp's table name
func (*ArchivedUp) TableName() string {
	return TableNameArchivedUp
}
