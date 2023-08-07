package model

import (
	"github.com/golang-module/carbon/v2"
)

type SysMenu struct {
	ID         uint            `gorm:"primaryKey;column:id;not null;comment: 'ID';AUTO_INCREMENT;type:bigint(20)" json:"id"`
	Pid        uint            `gorm:"column:pid;not null;default:0;comment: '上级ID，一级菜单为0';type:bigint(20)" json:"pid"`
	Name       string          `gorm:"column:name;not null;default:'';comment: '菜单名称';type:varchar(200)" json:"name"`
	URL        string          `gorm:"column:url;not null;default:'';comment: '菜单URL';type:varchar(200)" json:"url"`
	Authority  string          `gorm:"column:authority;not null;default:'';comment: '授权标识(多个用逗号分隔，如：SYS:MENU:LIST,SYS:MENU:SAVE)';type:varchar(500)" json:"authority"`
	Type       int8            `gorm:"column:type;not null;default:0;comment: '类型   0：菜单   1：按钮   2：接口';type:tinyint(4)" json:"type"`
	OpenStyle  int8            `gorm:"column:open_style;not null;default:0;comment: '打开方式   0：内部   1：外部';type:tinyint(4)" json:"openStyle"`
	Icon       string          `gorm:"column:icon;not null;default:'';comment: '菜单图标';type:varchar(50)" json:"icon"`
	Sort       int32           `gorm:"column:sort;not null;default:0;comment: '排序';type:int(11)" json:"sort"`
	Version    int32           `gorm:"column:version;not null;default:0;comment: '版本号';type:int(11)" json:"version"`
	Deleted    int8            `gorm:"column:deleted;not null;default:0;comment: '删除标识  0：正常   1：已删除';type:tinyint(4)" json:"deleted"`
	Creator    int64           `gorm:"column:creator;not null;default:0;comment: '创建者';type:bigint(20)" json:"creator"`
	CreateTime carbon.DateTime `gorm:"column:create_time;not null;comment: '创建时间';type:datetime" json:"createTime"`
	Updater    int64           `gorm:"column:updater;not null;default:0;comment: '更新者';type:bigint(20)" json:"updater"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;not null;comment: '更新时间';type:datetime" json:"update_time"`
	Children   []SysMenu       `gorm:"foreignKey:Pid;references:ID" json:"children"`
	ParentName string          `gorm:"-"`
	ParentMenu *SysMenu        `gorm:"foreignKey:Pid;references:ID" json:"parentMenu"`
}

func (s *SysMenu) TableName() string {
	return "sys_menu"
}
