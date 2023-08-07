package types

import "gincms/app/common/typescom"

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
