package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysDictType struct {
	ID         uint            `gorm:"primaryKey;column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	DictType   string          `gorm:"column:dict_type;not null;comment: '字典类型';type:varchar(100)" json:"dictType"`
	DictName   string          `gorm:"column:dict_name;not null;comment: '字典名称';type:varchar(255)" json:"dict_name"`
	DictSource int8            `gorm:"column:dict_source;not null;default:0;comment: '来源  0：字典数据  1：动态SQL';type:tinyint(4)" json:"dict_source"`
	DictSql    string          `gorm:"column:dict_sql;not null;default:'';comment: '动态SQL';type:varchar(500)" json:"dict_sql"`
	Remark     string          `gorm:"column:remark;not null;default:'';comment: '备注';type:varchar(255)" json:"remark"`
	Sort       int32           `gorm:"column:sort;not null;default:0;comment: '排序';type:int(11)" json:"sort"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"create_time"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
	DataList   []SysDictData   `gorm:"foreignKey:DictTypeID;references:ID" json:"dataList"`
}

func (s *SysDictType) TableName() string {
	return "sys_dict_type"
}
