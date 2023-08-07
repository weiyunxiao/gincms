package sys

import (
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var RoleCtl = new(roleCtl)

// 角色管理
type roleCtl struct {
}

// AddRole 添加
func (r *roleCtl) AddRole(c *gin.Context) {
	var req types.RoleAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.RoleService.AddRole(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOk(c)
}

// UpdateRole 更新
func (r *roleCtl) UpdateRole(c *gin.Context) {
	var req types.RoleAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if req.Id <= 0 {
		jsonresp.JsonFailWithMessage("请指定数据", c)
		return
	}

	err := sys.RoleService.UpdateRole(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// GetRole 获取一条角色信息
func (r *roleCtl) GetRole(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	oneData, err := sys.RoleService.GetRole(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(oneData, c)
}

// RolePage 分页数据
func (r *roleCtl) RolePage(c *gin.Context) {
	var req types.RolePageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.RoleService.RolePage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// RoleUserPage 某个角色拥有的用户分页列表
func (r *roleCtl) RoleUserPage(c *gin.Context) {
	var req types.RoleHaveUserPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.RoleService.RoleUserPage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// RoleList 获取可用角色列表数据
func (r *roleCtl) RoleList(c *gin.Context) {
	list, err := sys.RoleService.RoleList(c)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(list, c)
}

// DelRole 根据id数组删除
func (r *roleCtl) DelRole(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.RoleService.DelRole(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// DelRoleUser 移除某个角色下的用户
func (r *roleCtl) DelRoleUser(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	roleId := cast.ToInt(c.DefaultQuery("roleId", ""))
	if roleId <= 0 {
		jsonresp.JsonFailWithMessage("请传递角色id", c)
		return
	}

	err := sys.RoleService.DelRoleUser(c, roleId, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// AddRoleUser 某个角色关联多个用户的操作
func (r *roleCtl) AddRoleUser(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	roleId := cast.ToInt(c.DefaultQuery("roleId", ""))
	if roleId <= 0 {
		jsonresp.JsonFailWithMessage("请传递角色id", c)
		return
	}
	err := sys.RoleService.AddRoleUser(c, roleId, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// RoleMenu 添加编辑角色时，显示的所有系统菜单
func (r *roleCtl) RoleMenu(c *gin.Context) {
	tree, err := sys.RoleService.RoleMenu(c)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(tree, c)
}
