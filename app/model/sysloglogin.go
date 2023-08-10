package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysLogLogin struct {
	ID         uint            `gorm:"primaryKey;column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Username   string          `gorm:"column:username;not null;default:'';comment: '用户名';type:varchar(50)" json:"username"`
	Ip         string          `gorm:"column:ip;not null;default:'';comment: '登录IP';type:varchar(32)" json:"ip"`
	Address    string          `gorm:"column:address;not null;default:'';comment: '登录地点';type:varchar(32)" json:"address"`
	UserAgent  string          `gorm:"column:user_agent;not null;default:'';comment: 'USER AGENT';type:varchar(500)" json:"userAgent"`
	Status     int8            `gorm:"column:status;not null;default:0;comment: '登录状态  0：失败   1：成功';type:tinyint(4)" json:"status"`
	Operation  int8            `gorm:"column:operation;not null;default:0;comment: '操作信息   0：登录成功   1：退出成功  2：验证码错误  3：账号密码错误';type:tinyint(3)" json:"operation"`
	TenantID   int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
}

func (s *SysLogLogin) TableName() string {
	return "sys_log_login"
}
