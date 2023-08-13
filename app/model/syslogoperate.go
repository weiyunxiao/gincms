package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysLogOperate struct {
	ID          uint            `gorm:"primaryKey;column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Module      string          `gorm:"column:module;not null;default:'';comment: '模块名';type:varchar(100)" json:"module"`
	Name        string          `gorm:"column:name;not null;default:'';comment: '操作名';type:varchar(100)" json:"name"`
	ReqUri      string          `gorm:"column:req_uri;not null;default:'';comment: '请求URI';type:varchar(200)" json:"reqUri"`
	ReqMethod   string          `gorm:"column:req_method;not null;default:'';comment: '请求方法';type:varchar(20)" json:"reqMethod"`
	ReqParams   string          `gorm:"column:req_params;not null;comment: '请求参数';type:text" json:"req_params"`
	Ip          string          `gorm:"column:ip;not null;default:'';comment: '操作IP';type:varchar(32)" json:"ip"`
	Address     string          `gorm:"column:address;not null;default:'';comment: '登录地点';type:varchar(32)" json:"address"`
	UserAgent   string          `gorm:"column:user_agent;not null;default:'';comment: 'USER AGENT';type:varchar(500)" json:"userAgent"`
	OperateType int8            `gorm:"column:operate_type;not null;default:0;comment: '操作类型';type:tinyint(4)" json:"operate_type"`
	Duration    string          `gorm:"column:duration;not null;default:'';comment: '执行时长';type:varchar(50)" json:"duration"`
	Status      int8            `gorm:"column:status;not null;default:1;comment: '操作状态  0：失败   1：成功';type:tinyint(4)" json:"status"`
	UserID      int64           `gorm:"column:user_id;not null;default:0;comment: '用户ID';type:bigint(20)" json:"user_id"`
	RealName    string          `gorm:"column:real_name;not null;default:'';comment: '操作人';type:varchar(50)" json:"real_name"`
	ResultMsg   string          `gorm:"column:result_msg;not null;default:'';comment: '返回消息';type:varchar(500)" json:"result_msg"`
	TenantID    int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	CreateTime  carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	OperateUser *SysUser        `gorm:"foreignKey:UserID;references:ID" json:"operateUser"`
}

func (s *SysLogOperate) TableName() string {
	return "sys_log_operate"
}
