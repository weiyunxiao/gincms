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
	authArr, err := auth.GetUserAuth(c, c.GetInt64("uid"))
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
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
