package sys

import (
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var UserCtl = new(userCtl)

type userCtl struct {
}

// AddUser 添加一个用户
func (u *userCtl) AddUser(c *gin.Context) {
	var req types.UserAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	isExist, err := sys.UserService.IsUserNameExist(c, req.Username)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	if isExist {
		jsonresp.JsonFailWithMessage("系统已存在相同的用户名了", c)
		return
	}

	err = sys.UserService.AddUser(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOk(c)
}

// GetUser 得到一个用户
func (u *userCtl) GetUser(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	user, err := sys.UserService.GetUser(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(user, c)
}

// UpdateUser 更新
func (u *userCtl) UpdateUser(c *gin.Context) {
	var req types.UserAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if req.Id <= 0 {
		jsonresp.JsonFailWithMessage("请指定用户", c)
		return
	}

	err := sys.UserService.UpdateUser(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// UserPage 分页数据
func (u *userCtl) UserPage(c *gin.Context) {
	var req types.UserPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.UserService.UserPage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// DelUser 根据id数组删除
func (u *userCtl) DelUser(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.UserService.DelUser(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// Info 用户登录时，获取用户信息给前端
func (u *userCtl) Info(c *gin.Context) {
	var uid int64
	if id, ok := c.Get("uid"); ok {
		uid = cast.ToInt64(id)
	}
	if uid == 0 {
		jsonresp.JsonFailWithMessage("无法获取获取用户", c)
		return
	}
	user, err := sys.UserService.Info(c, uid)
	if err != nil {
		jsonresp.JsonFail500(c)
		return
	}
	jsonresp.JsonOkWithData(user, c)
}

// UpdateSelfInfo 更新自己的信息
func (u *userCtl) UpdateSelfInfo(c *gin.Context) {
	var req types.UpdateSelfInfoReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if err := sys.UserService.UpdateSelfInfo(c, &req); err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}
