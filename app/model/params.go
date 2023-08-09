package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysParams struct {
	ID         uint            `gorm:"column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	ParamName  string          `gorm:"column:param_name;not null;default:'';comment: '参数名称';type:varchar(100)" json:"paramName"`
	ParamType  int8            `gorm:"column:param_type;not null;comment: '系统参数   0：否   1：是';type:tinyint(4)" json:"paramType"`
	ParamKey   string          `gorm:"column:param_key;not null;default:'';comment: '参数键';type:varchar(100)" json:"paramKey"`
	ParamValue string          `gorm:"column:param_value;not null;default:'';comment: '参数值';type:varchar(2000)" json:"paramValue"`
	Remark     string          `gorm:"column:remark;not null;default:'';comment: '备注';type:varchar(200)" json:"remark"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
}

func (s *SysParams) TableName() string {
	return "sys_params"
}
