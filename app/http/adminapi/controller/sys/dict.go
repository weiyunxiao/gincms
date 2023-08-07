package sys

import (
	"gincms/app"
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
