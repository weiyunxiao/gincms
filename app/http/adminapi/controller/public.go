package controller

import (
	"gincms/app"
	"gincms/app/http/adminapi/service"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

var PublicCtl = new(publicCtl)

type publicCtl struct {
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// LoginCaptchaEnabled 是否打开验证码登录
func (p *publicCtl) LoginCaptchaEnabled(c *gin.Context) {
	jsonresp.JsonOkWithData(app.Config.Http.LoginCaptchaEnabled, c)
}

// Captcha 验证码
func (p *publicCtl) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	limitCaptchaNum := app.Config.AdminCaptcha.LimitCaptchaNum // 是否开启防爆次数
	LimitTimeOut := app.Config.AdminCaptcha.LimitTimeOut       // 缓存超时时间
	IpKey := c.ClientIP() + "adminCaptcha"
	v, ok := app.LocalCache.Get(IpKey)
	if !ok {
		app.LocalCache.Set(IpKey, 1, time.Duration(LimitTimeOut)*time.Second)
	} else {
		_ = app.LocalCache.Increment(IpKey, 1)
	}
	var allow bool
	if limitCaptchaNum == 0 || limitCaptchaNum > cast.ToInt(v) {
		allow = true
	}
	if !allow {
		jsonresp.JsonOkWithDetailed(types.SysCaptchaResponse{
			CaptchaId:     "",
			PicPath:       "",
			CaptchaLength: 0,
			OpenCaptcha:   allow,
		}, "请不要频繁请求验证码", c)
		return
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(app.Config.AdminCaptcha.ImgHeight, app.Config.AdminCaptcha.ImgWidth, app.Config.AdminCaptcha.CharLen, 0.2, 50)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		app.Logger.Error("生成验证码失败", zap.Error(err))
		jsonresp.JsonFailWithMessage("验证码获取失败", c)
		return
	}
	jsonresp.JsonOkWithDetailed(types.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: app.Config.AdminCaptcha.CharLen,
		OpenCaptcha:   allow,
	}, "验证码获取成功", c)
}

// Login 用户登录
func (p *publicCtl) Login(c *gin.Context) {
	var req types.LoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	if app.Config.Http.LoginCaptchaEnabled {
		if !store.Verify(req.CaptchaKey, req.Captcha, true) {
			jsonresp.JsonFailWithMessage("验证码错误", c)
			return
		}
	}

	token, expireUnix, _ := service.PublicService.Login(&req, c)
	if len(token) == 0 {
		jsonresp.JsonFailWithMessage("用户登录失败,请检查账号及密码，及用户是否被禁用", c)
		return
	}

	jsonresp.JsonOkWithData(gin.H{
		"access_token":   token,
		"expire_at_unix": expireUnix,
	}, c)
}
