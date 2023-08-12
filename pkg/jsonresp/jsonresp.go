package jsonresp

import (
	"gincms/app"
	"gincms/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	success_code      = 0
	param_erro_code   = 400 //参数错误
	common_err_code   = 7
	RefreshToken_code = 409 //与前端对接，刷新token验证token时为这个验证码返回
)

type baseJsonResp struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	ReqKey string      `json:"reqKey"`
}

func JsonResult(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, baseJsonResp{
		code,
		data,
		msg,
		pkg.GetReqKey(c),
	})
}

func JsonOk(c *gin.Context) {
	JsonResult(success_code, map[string]interface{}{}, "操作成功", c)
}

func JsonOkWithMessage(message string, c *gin.Context) {
	JsonResult(success_code, map[string]interface{}{}, message, c)
}

func JsonOkWithData(data interface{}, c *gin.Context) {
	JsonResult(success_code, data, "查询成功", c)
}

func JsonOkWithDetailed(data interface{}, message string, c *gin.Context) {
	JsonResult(success_code, data, message, c)
}

func JsonFail(c *gin.Context) {
	JsonResult(common_err_code, map[string]interface{}{}, "操作失败", c)
}
func JsonFailParame(c *gin.Context, err error) {
	if app.Config.App.LogParamErr {
		pkg.ParamErrLog(c, err)
	}
	if app.Config.App.LogParamShowClient {
		JsonResult(param_erro_code, map[string]interface{}{}, "参数错误:"+err.Error(), c)
		return
	}
	JsonResult(param_erro_code, map[string]interface{}{}, "参数错误", c)
}

func JsonFailParameWithMsg(message string, c *gin.Context) {
	JsonResult(param_erro_code, map[string]interface{}{}, message, c)
}

func JsonFailWithMessage(message string, c *gin.Context) {
	JsonResult(common_err_code, map[string]interface{}{}, message, c)
}

func JsonFailWithDetailed(data interface{}, message string, c *gin.Context) {
	JsonResult(common_err_code, data, message, c)
}

func JsonFail500(c *gin.Context) {
	JsonResult(http.StatusInternalServerError, map[string]interface{}{}, "操作失败", c)
}
