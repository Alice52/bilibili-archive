// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArchivedLike = "archived_like"

// ArchivedLike 点赞视频
type ArchivedLike struct {
	ID         int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime *time.Time     `gorm:"column:create_time;type:datetime(3);autoCreateTime" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;type:datetime(3);autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime(3)" json:"delete_time"`
	Fid        *string        `gorm:"column:fid;type:varchar(64);comment:bili folder" json:"fid"`                                                                                                              // bili folder
	Vid        *string        `gorm:"column:vid;type:varchar(64);comment:bili avid" json:"vid"`                                                                                                                // bili avid
	Cover      *string        `gorm:"column:cover;type:varchar(256);comment:video cover" json:"cover"`                                                                                                         // video cover
	Duration   *int64         `gorm:"column:duration;type:bigint;comment:video duration" json:"duration"`                                                                                                      // video duration
	CoinTime   int64          `gorm:"column:coin_time;type:bigint;not null;comment:video favor time" json:"coin_time"`                                                                                         // video favor time
	Intro      *string        `gorm:"column:intro;type:varchar(64);comment:video intro" json:"intro"`                                                                                                          // video intro
	Title      *string        `gorm:"column:title;type:varchar(64);comment:video title" json:"title"`                                                                                                          // video title
	Type       *string        `gorm:"column:type;type:varchar(64);comment:video type" json:"type"`                                                                                                             // video type
	Season     *string        `gorm:"column:season;type:varchar(64);comment:video season" json:"season"`                                                                                                       // video season
	UpperMid   *string        `gorm:"column:upper_mid;type:json;comment:{"mid": 173986740, "name": "这个月-"}" json:"upper_mid"`                                                                                  // {"mid": 173986740, "name": "这个月-"}
	CntInfo    *string        `gorm:"column:cnt_info;type:json;comment:{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }" json:"cnt_info"` // {"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }
}

// TableName ArchivedLike's table name
func (*ArchivedLike) TableName() string {
	return TableNameArchivedLike
}
