package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysDictData struct {
	ID         uint            `gorm:"column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	DictTypeID int64           `gorm:"column:dict_type_id;not null;comment: '字典类型ID';type:bigint(20)" json:"dict_type_id"`
	DictLabel  string          `gorm:"column:dict_label;not null;comment: '字典标签';type:varchar(255)" json:"dictLabel"`
	DictValue  string          `gorm:"column:dict_value;not null;default:'';comment: '字典值';type:varchar(255)" json:"dictValue"`
	LabelClass string          `gorm:"column:label_class;not null;default:'';comment: '标签样式';type:varchar(100)" json:"labelClass"`
	Remark     string          `gorm:"column:remark;not null;default:'';comment: '备注';type:varchar(255)" json:"remark"`
	Sort       int32           `gorm:"column:sort;not null;default:0;comment: '排序';type:int(11)" json:"sort"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"create_time"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
}

func (s *SysDictData) TableName() string {
	return "sys_dict_data"
}
