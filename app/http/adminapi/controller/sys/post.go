package sys

import (
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
)

var PostCtl = new(postCtl)

// 岗位管理
type postCtl struct {
}

// AddPost 添加岗位
func (p *postCtl) AddPost(c *gin.Context) {
	var req types.PostAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.PostService.AddPost(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// UpdatePost 更新一个岗位
func (p *postCtl) UpdatePost(c *gin.Context) {
	var req types.PostAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.PostService.UpdatePost(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// GetPost 得到一个岗位
func (p *postCtl) GetPost(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	post, err := sys.PostService.GetPost(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(post, c)
}

// DelPost 根据id数组删除岗位
func (p *postCtl) DelPost(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.PostService.DelPost(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// PostPage 岗位分页数据
func (p *postCtl) PostPage(c *gin.Context) {
	var req types.PostPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.PostService.PostPage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// PostList 获取可用岗位列表-比如用户添加时需要选择岗位
func (p *postCtl) PostList(c *gin.Context) {
	list, err := sys.PostService.PostList(c)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(list, c)

}
