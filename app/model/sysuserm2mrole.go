package model

type SysUserM2mRole struct {
	UserID int64 `gorm:"column:user_id;not null;default:0;comment: '用户ID';type:bigint(20)" json:"user_id"`
	RoleID int64 `gorm:"column:role_id;not null;default:0;comment: '角色ID';type:bigint(20)" json:"role_id"`
}

func (s *SysUserM2mRole) TableName() string {
	return "sys_user_m2m_role"
}
