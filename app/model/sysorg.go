package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysOrg struct {
	ID         uint            `gorm:"primaryKey:column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Pid        uint            `gorm:"column:pid;not null;default:0;comment: '上级ID';type:bigint(20)" json:"pid"`
	Name       string          `gorm:"column:name;not null;default:'';comment: '机构名称';type:varchar(50)" json:"name"`
	Sort       int             `gorm:"column:sort;not null;default:0;comment: '排序';type:int(11)" json:"sort"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"create_time"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
	ParentName string          `gorm:"-" json:"parentName"`
	ParenOrg   *SysOrg         `gorm:"foreignKey:Pid;references:ID"`
	Children   []SysOrg        `gorm:"foreignKey:Pid;references:ID" json:"children"`
}

func (s *SysOrg) TableName() string {
	return "sys_org"
}
