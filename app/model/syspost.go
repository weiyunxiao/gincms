package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysPost struct {
	ID         uint            `gorm:"primaryKey;column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	PostCode   string          `gorm:"column:post_code;not null;default:'';comment: '岗位编码';type:varchar(100)" json:"postCode"`
	PostName   string          `gorm:"column:post_name;not null;default:'';comment: '岗位名称';type:varchar(100)" json:"postName"`
	Sort       int             `gorm:"column:sort;not null;default:0;comment: '排序';type:int(11)" json:"sort"`
	Status     int8            `gorm:"column:status;not null;default:1;comment: '状态  0：停用   1：正常';type:tinyint(4)" json:"status"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
}

func (s *SysPost) TableName() string {
	return "sys_post"
}
