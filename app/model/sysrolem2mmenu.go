package model

type SysRoleM2mMenu struct {
	RoleID int64 `gorm:"column:role_id;not null;comment: '角色ID';type:bigint(20)" json:"role_id"`
	MenuID int64 `gorm:"column:menu_id;not null;comment: '菜单ID';type:bigint(20)" json:"menu_id"`
}

func (s *SysRoleM2mMenu) TableName() string {
	return "sys_role_m2m_menu"
}
