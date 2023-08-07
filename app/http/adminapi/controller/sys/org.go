package sys

import (
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
)

var OrgCtl = new(orgCtl)

// 机构管理
type orgCtl struct {
}

// AddOrg 添加机构
func (p *orgCtl) AddOrg(c *gin.Context) {
	var req types.OrgAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.OrgService.AddOrg(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// UpdateOrg 更新一个机构
func (p *orgCtl) UpdateOrg(c *gin.Context) {
	var req types.OrgAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.OrgService.UpdateOrg(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// GetOrg 得到一个机构
func (p *orgCtl) GetOrg(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	org, err := sys.OrgService.GetOrg(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(org, c)
}

// DelOrg 根据id删除机构
func (p *orgCtl) DelOrg(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.OrgService.DelOrg(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// OrgList 获取可用机构列表数据
func (p *orgCtl) OrgList(c *gin.Context) {
	var req typescom.PageOrderCommonReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	listTree, err := sys.OrgService.OrgList(c, req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(listTree, c)
}
