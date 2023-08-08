package sys

import (
	"gincms/app"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var DictCtl = new(dictCtl)

type dictCtl struct {
}

// AddDictType 添加
func (r *dictCtl) AddDictType(c *gin.Context) {
	var req types.DictTypeAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.DictService.AddDictType(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOk(c)
}

// GetDictType 获取单条
func (d *dictCtl) GetDictType(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	oneData, err := sys.DictService.GetDictType(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(oneData, c)
}

// DictTypePage 获取列表分页
func (d *dictCtl) DictTypePage(c *gin.Context) {
	var req types.DictPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.DictService.TypePage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// DelDictType 删除多条字典类型
func (d *dictCtl) DelDictType(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.DictService.DelType(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// UpdateDictType 更新
func (m *dictCtl) UpdateDictType(c *gin.Context) {
	var req types.DictTypeAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if req.Id <= 0 {
		jsonresp.JsonFailWithMessage("请指定数据", c)
		return
	}

	err := sys.DictService.UpdateDictType(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// DictTypeAll 获取所有的字典数据
func (d *dictCtl) DictTypeAll(c *gin.Context) {
	dictTypes, err := sys.DictService.DictTypeAll(c)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	output, _ := sonic.Marshal(dictTypes)
	var resp = make([]types.SysDictTypeResp, 0)
	err = sonic.Unmarshal(output, &resp)
	if err != nil {
		app.Logger.Error("解析json失败", zap.Error(err))
	}
	jsonresp.JsonOkWithData(resp, c)
}

// DictDataPage 获取字典二级数据列表分页
func (d *dictCtl) DictDataPage(c *gin.Context) {
	var req types.DictDataPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.DictService.DictDataPage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// DelDictData 删除多条二级字典数
func (d *dictCtl) DelDictData(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.DictService.DelDictData(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}

// AddDictData 添加
func (r *dictCtl) AddDictData(c *gin.Context) {
	var req types.DictDataAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.DictService.AddDictData(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOk(c)
}

// GetDictData 获取单条
func (d *dictCtl) GetDictData(c *gin.Context) {
	var req typescom.IDReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	oneData, err := sys.DictService.GetDictData(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOkWithData(oneData, c)
}

// UpdateDictData 更新
func (m *dictCtl) UpdateDictData(c *gin.Context) {
	var req types.DictDataAddSaveReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if req.Id <= 0 {
		jsonresp.JsonFailWithMessage("请指定数据", c)
		return
	}

	err := sys.DictService.UpdateDictData(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}
