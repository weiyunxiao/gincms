package types

import (
	"gincms/app/common/typescom"
)

// LoginReq 登录请求体
type LoginReq struct {
	UserName   string `json:"username" form:"username" binding:"required,max=50"`
	Password   string `json:"password" form:"password" binding:"required,max=50"`
	CaptchaKey string `json:"key" form:"key"`
	Captcha    string `json:"captcha" form:"captcha"`
}

// PostPageReq 岗位管理分页列表请求
type PostPageReq struct {
	PostCode string `json:"postCode" form:"postCode"`
	PostName string `json:"postName" form:"postName"`
	Status   string `json:"status" form:"status"`
	typescom.PageOrderCommonReq
}

// PostAddSaveReq 添岗位加修改的请求
type PostAddSaveReq struct {
	Id       uint   `json:"id"`
	PostCode string `json:"postCode" binding:"required,max=50"`
	PostName string `json:"postName" binding:"required,max=50"`
	Sort     int    `json:"sort"`
	Status   int    `json:"status"`
}

// OrgAddSaveReq 机构添加修改请求
type OrgAddSaveReq struct {
	Id   uint   `json:"id"`
	Name string `json:"name" binding:"required,max=50"`
	Pid  uint   `json:"pid"`
	Sort int    `json:"sort"`
}

// UserAddSaveReq 用户添加修改请求
type UserAddSaveReq struct {
	Id         uint   `json:"id"`
	Username   string `json:"username" binding:"required,min=4,max=20"`
	RealName   string `json:"realName" binding:"required,min=2,max=20"`
	OrgId      int    `json:"orgId" binding:"required"`
	OrgName    string `json:"orgName"`
	Password   string `json:"password"`
	Gender     any    `json:"gender"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile" binding:"required,min=11"`
	RoleIdList []uint `json:"roleIdList"`
	PostIdList []uint `json:"postIdList"`
	Status     int8   `json:"status"  binding:"oneof=0 1"`
}

// UserPageReq 用户分页列表请求
type UserPageReq struct {
	UserName string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Gender   string `json:"gender" form:"gender"`
	typescom.PageOrderCommonReq
}

// RolePageReq 角色分页列表请求
type RolePageReq struct {
	Name string `json:"name" form:"name"`
	typescom.PageOrderCommonReq
}

// RoleAddSaveReq 角色添加更新请求
type RoleAddSaveReq struct {
	Id         uint   `json:"id"`
	Name       string `json:"name" form:"name" binding:"required,max=20"`
	MenuIdList []uint `json:"menuIdList" form:"menuIdList"`
	OrgIdList  []uint `json:"orgIdList" form:"orgIdList"`
	Remark     string `json:"remark" form:"remark"`
}

// RoleHaveUserPageReq 某个角色拥有的用户分页列表请求
type RoleHaveUserPageReq struct {
	RoleId   int    `json:"roleId" form:"roleId" binding:"required,min=1"`
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Gender   string `json:"gender" form:"gender"`
	typescom.PageOrderCommonReq
}

// DelRoleUserReq 移除某个角色下的用户
type DelRoleUserReq struct {
	RoleId int    `json:"roleId" form:"roleId" binding:"required"`
	IdArr  []uint `json:"id_arr" form:"id_arr" binding:"required,min=1"`
}

// MenuAddSaveReq 菜单的添加与修改请求
type MenuAddSaveReq struct {
	Id         uint   `json:"id" form:"id"`
	Type       int8   `json:"type" form:"type"`
	Name       string `json:"name" form:"name" binding:"required"`
	Pid        uint   `json:"pid" form:"pid"`
	ParentName string `json:"parentName" form:"parentName"`
	Url        string `json:"url" form:"url"`
	Authority  string `json:"authority" form:"authority"`
	Sort       int32  `json:"sort" form:"sort"`
	Icon       string `json:"icon" form:"icon"`
	OpenStyle  int8   `json:"openStyle" form:"openStyle"`
}

// DictPageReq 用户分页列表请求
type DictPageReq struct {
	DictName string `json:"dictName" form:"dictName"`
	DictType string `json:"dictType" form:"dictType"`
	typescom.PageOrderCommonReq
}

// DictTypeAddSaveReq 字典类型添加修改请求
type DictTypeAddSaveReq struct {
	Id         uint   `json:"id" form:"id"`
	DictType   string `json:"dictType" form:"dictType" binding:"max=50"`
	DictName   string `json:"dictName" form:"dictName" binding:"max=50"`
	Sort       int32  `json:"sort" form:"sort" binding:"max=50000"`
	DictSource int8   `json:"dictSource" form:"dictSource" binding:"max=10"`
	DictSql    string `json:"dictSql" form:"dictSql" binding:"max=450"`
	Remark     string `json:"remark" form:"remark" binding:"max=450"`
}

// DictDataPageReq 字典二级数据分页请求
type DictDataPageReq struct {
	DictTypeId int `json:"dictTypeId" form:"dictTypeId" binding:"required"`
	typescom.PageOrderCommonReq
}

// DictDataAddSaveReq 字典二级数据添加修改请求
type DictDataAddSaveReq struct {
	Id         uint   `json:"id" form:"id"`
	DictTypeId int    `json:"dictTypeId" form:"dictTypeId" binding:"required"`
	DictLabel  string `json:"dictLabel" form:"dictLabel" binding:"required,max=200"`
	DictValue  string `json:"dictValue" form:"dictValue" binding:"required,max=200"`
	LabelClass string `json:"labelClass" form:"labelClass" binding:"max=80"`
	Sort       int32  `json:"sort" form:"sort"`
	Remark     string `json:"remark" form:"remark" binding:"max=450"`
}

// ParamsAddSaveReq 参数管理添加修改请求
type ParamsAddSaveReq struct {
	Id         uint   `json:"id" form:"id"`
	ParamName  string `json:"paramName" form:"paramName" binding:"required,max=90"`
	ParamType  int8   `json:"paramType" form:"paramType" binding:"max=1"`
	ParamKey   string `json:"paramKey" form:"paramKey" binding:"required,max=90"`
	ParamValue string `json:"paramValue" form:"paramValue" binding:"required,max=1900"`
	Remark     string `json:"remark" form:"dictTypeId" binding:"max=200"`
}

// ParamsPageReq 参数管理分页请求
type ParamsPageReq struct {
	ParamType  string `json:"paramType" form:"paramType" binding:"max=1"`
	ParamKey   string `json:"paramKey" form:"paramKey" binding:"max=100"`
	ParamValue string `json:"paramValue" form:"paramValue" binding:"max=200"`
	typescom.PageOrderCommonReq
}

// LogLoginLogoutPageReq 登录登出日志分页请求
type LogLoginLogoutPageReq struct {
	Username string `json:"username" form:"username" binding:"max=50"`
	Address  string `json:"address" form:"address" binding:"max=30"`
	Status   string `json:"status" form:"status" binding:"max=1"`
	typescom.PageOrderCommonReq
}

// OperateLogPageReq 操作日志分页请求
type OperateLogPageReq struct {
	RealName string `json:"realName" form:"realName" binding:"max=50"`
	ReqUri   string `json:"reqUri" form:"reqUri" binding:"max=100"`
	typescom.PageOrderCommonReq
}

// AttachmentPageReq 附件记录分页请求
type AttachmentPageReq struct {
	Name     string `json:"name" form:"name" binding:"max=50"`
	Platform string `json:"platform" form:"platform" binding:"max=30"`
	typescom.PageOrderCommonReq
}

// UpdateSelfInfoReq 更新头像及密码
type UpdateSelfInfoReq struct {
	Avatar      string `json:"avatar" form:"avatar"`
	Password    string `json:"password" form:"password" binding:"required,min=4"`
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required,min=4"`
}
