package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysAttachment struct {
	ID         uint            `gorm:"column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Name       string          `gorm:"column:name;not null;comment: '附件名称';type:varchar(255)" json:"name"`
	URL        string          `gorm:"column:url;not null;comment: '附件地址';type:varchar(255)" json:"url"`
	Size       int64           `gorm:"column:size;not null;default:0;comment: '附件大小';type:bigint(20)" json:"size"`
	SizeTip    string          `gorm:"column:size_tip;not null;default:'';comment: '附件大小KB M显示';type:varchar(50)" json:"sizeTip"`
	Platform   string          `gorm:"column:platform;not null;default:'';comment: '存储平台';type:varchar(50)" json:"platform"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
}

func (s *SysAttachment) TableName() string {
	return "sys_attachment"
}
