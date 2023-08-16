package sys

import (
	"gincms/app/common/auth"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
)

var MenuCtl = new(menuCtl)

type menuCtl struct {
}

// AddMenu 添加
func (m *menuCtl) AddMenu(c *gin.Context) {
	var req types.MenuAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.MenuService.AddMenu(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOk(c)
}

// GetMenu 获取单条信息
func (m *menuCtl) GetMenu(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	oneData, err := sys.MenuService.GetMenu(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(oneData, c)
}

// UpdateMenu 更新
func (m *menuCtl) UpdateMenu(c *gin.Context) {
	var req types.MenuAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if req.Id <= 0 {
		jsonresp.JsonFailWithMessage("请指定数据", c)
		return
	}

	err := sys.MenuService.UpdateMenu(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// MenuList 系统所有菜单-树结构列表
func (m *menuCtl) MenuList(c *gin.Context) {
	menuList, err := sys.MenuService.MenuList(c)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(&menuList, c)

}

// DelMenu 删除单条
func (m *menuCtl) DelMenu(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.MenuService.DelMenu(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)

}

// Authority 获取用户的所有权限
func (m *menuCtl) Authority(c *gin.Context) {
	authArr := []string{
		"sms:platform:save",
		"sys:role:page",
		"sys:user:export",
		"schedule:run",
		"sys:attachment:delete",
		"sms:platform:page",
		"sys:operate:all",
		"sys:user:page",
		"sys:role:list",
		"sys:post:save",
		"sys:dict:info",
		"sys:org:info",
		"sys:post:delete",
		"sys:role:save",
		"monitor:user:all",
		"sys:org:update",
		"monitor:cache:all",
		"sys:role:update",
		"sys:log:login",
		"sys:org:list",
		"sys:post:page",
		"sys:user:save",
		"schedule:delete",
		"sms:platform:update",
		"monitor:server:all",
		"sys:menu:update",
		"sys:menu:delete",
		"schedule:info",
		"sys:dict:delete",
		"sys:menu:list",
		"sms:platform:info",
		"sys:params:all",
		"sys:dict:page",
		"sys:user:import",
		"sys:user:delete",
		"sys:org:save",
		"sys:menu:info",
		"sys:user:update",
		"schedule:log",
		"sys:role:menu",
		"sys:post:info",
		"sys:post:update",
		"schedule:update",
		"sys:menu:save",
		"schedule:page",
		"sms:platform:delete",
		"sys:org:delete",
		"sys:role:info",
		"sys:attachment:page",
		"sys:user:info",
		"sys:attachment:save",
		"sys:role:delete",
		"schedule:save",
		"online:table:all",
		"sys:dict:update",
		"sms:log",
		"sys:dict:save",
	}
	jsonresp.JsonOkWithData(authArr, c)
}

// Nav  获取用户的菜单列表
func (m *menuCtl) Nav(c *gin.Context) {
	menuList, err := auth.GetUserMenu(c, c.GetInt64("uid"), "id,pid,name,url,icon,sort,create_time")
	if err != nil {
		jsonresp.JsonFailWithMessage("获取菜单出错", c)
		return
	}
	jsonresp.JsonOkWithData(&menuList, c)
}
