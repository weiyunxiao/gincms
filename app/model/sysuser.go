package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysUser struct {
	ID           uint64          `gorm:"primaryKey;column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Username     string          `gorm:"column:username;not null;comment: '用户名';type:varchar(50)" json:"username"`
	Password     string          `gorm:"column:password;not null;default:'';comment: '密码';type:varchar(100)" json:"password"`
	RealName     string          `gorm:"column:real_name;not null;default:'';comment: '姓名';type:varchar(50)" json:"realName"`
	Avatar       string          `gorm:"column:avatar;not null;default:'';comment: '头像';type:varchar(200)" json:"avatar"`
	Gender       int8            `gorm:"column:gender;not null;default:0;comment: '性别   0：男   1：女   2：未知';type:tinyint(4)" json:"gender"`
	Email        string          `gorm:"column:email;not null;default:'';comment: '邮箱';type:varchar(100)" json:"email"`
	Mobile       string          `gorm:"column:mobile;not null;default:'';comment: '手机号';type:varchar(20)" json:"mobile"`
	OrgID        int             `gorm:"column:org_id;not null;comment: '机构ID';type:bigint(20)" json:"orgId"`
	SuperAdmin   int8            `gorm:"column:super_admin;not null;default:0;comment: '超级管理员   0：否   1：是';type:tinyint(4)" json:"super_admin"`
	Status       int8            `gorm:"column:status;not null;default:1;comment: '状态  0：停用   1：正常';type:tinyint(4)" json:"status"`
	TenantID     int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version      int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted      int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator      int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime   carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	Updater      int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime   carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
	OrgData      *SysOrg         `gorm:"foreignKey:OrgID;references:ID" json:"orgData"`
	OrgName      string          `gorm:"-" json:"orgName"`
	PostIdList   []uint          `gorm:"-" json:"postIdList"`
	RoleIdList   []uint          `gorm:"-" json:"roleIdList"`
	UserPostList []SysPost       `gorm:"many2many:sys_user_m2m_post;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:post_id"`
	UserRoleList []SysRole       `gorm:"many2many:sys_user_m2m_role;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:role_id"`
}

func (s *SysUser) TableName() string {
	return "sys_user"
}
