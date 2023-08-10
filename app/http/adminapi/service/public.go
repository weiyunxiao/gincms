package service

import (
	"gincms/app"
	"gincms/app/common/comservice"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"gincms/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var PublicService = new(publicService)

type publicService struct {
}

// LoginCaptchaEnabled 是否打开验证码登录
func (p publicService) LoginCaptchaEnabled(keyName string) bool {
	if v, ok := comservice.GetParam(keyName); ok {
		return cast.ToBool(v)
	}
	return false
}

// Login 用户登录逻辑
func (p *publicService) Login(req *types.LoginReq, c *gin.Context) (token string, expireAtUnix int64, err error) {
	var user model.SysUser
	err = app.DB().Model(&model.SysUser{}).Select("id,username,password").
		Where("username=? and status=1 and deleted=0", req.UserName).Take(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if user.ID == 0 {
		return
	}
	//密码错误
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}

	jwtObj := jwt.NewJWT()
	token, expireAtUnix = jwtObj.IssueToken(user.ID)
	return
}
