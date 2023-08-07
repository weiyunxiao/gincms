package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysRole struct {
	ID           uint            `gorm:"column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Name         string          `gorm:"column:name;not null;default:'';comment: '角色名称';type:varchar(50)" json:"name"`
	Remark       string          `gorm:"column:remark;not null;default:'';comment: '备注';type:varchar(100)" json:"remark"`
	DataScope    int8            `gorm:"column:data_scope;not null;default:0;comment: '数据范围  0：全部数据  1：本机构及子机构数据  2：本机构数据  3：本人数据  4：自定义数据';type:tinyint(4)" json:"data_scope"`
	OrgID        int64           `gorm:"column:org_id;not null;default:0;comment: '机构ID';type:bigint(20)" json:"org_id"`
	TenantID     int64           `gorm:"column:tenant_id;not null;default:0;comment: '租户ID';type:bigint(20)" json:"tenant_id"`
	Version      int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted      int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator      int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime   carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	Updater      int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime   carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
	MenuIdList   []int           `gorm:"-" json:"menuIdList"`
	RoleMenuList []SysMenu       `gorm:"many2many:sys_role_m2m_menu;foreignKey:id;joinForeignKey:role_id;References:id;joinReferences:menu_id"`
	RoleUserList []SysUser       `gorm:"many2many:sys_user_m2m_role;foreignKey:id;joinForeignKey:role_id;References:id;joinReferences:user_id"`
}

func (s *SysRole) TableName() string {
	return "sys_role"
}
