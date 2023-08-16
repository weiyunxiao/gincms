package jwt

import (
	"errors"
	"gincms/app"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {

	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID       uint64 `json:"user_id"`
	ExpireAtTime int64  `json:"expire_time"`
	IsFreshToken bool   `json:"is_fresh_token"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(app.Config.Http.JwtAccessSecret),
	}
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenStr, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenStr)

	// 2. 解析出错
	if err != nil {
		if errors.Is(err, jwtpkg.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		} else if errors.Is(err, jwtpkg.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) IssueToken(c *gin.Context, userID uint64) (string, int64) {

	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := jwt.expireAtTime()
	now := jwtpkg.NewNumericDate(timenowInTimezone())
	claims := JWTCustomClaims{
		userID,
		expireAtTime,
		false,
		jwtpkg.RegisteredClaims{
			NotBefore: now,                                               // 签名生效时间
			IssuedAt:  now,                                               // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtpkg.NewNumericDate(time.Unix(expireAtTime, 0)), // 签名过期时间
			Issuer:    "gincms",                                          // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		app.Logger.Error("生成token失败", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return "", 0
	}

	return token, expireAtTime
}

// IssueRefreshToken 生成刷新token所使用的token
func (jwt *JWT) IssueRefreshToken(c *gin.Context, userID uint64) (string, int64) {
	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := jwt.freshTokenExpireAtTime()
	now := jwtpkg.NewNumericDate(timenowInTimezone())
	claims := JWTCustomClaims{
		userID,
		expireAtTime,
		true,
		jwtpkg.RegisteredClaims{
			NotBefore: now,                                               // 签名生效时间
			IssuedAt:  now,                                               // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtpkg.NewNumericDate(time.Unix(expireAtTime, 0)), // 签名过期时间
			Issuer:    "gincms",                                          // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		app.Logger.Error("生成刷新token失败", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return "", 0
	}

	return token, expireAtTime
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	timenow := timenowInTimezone()

	expireTime := app.Config.Http.JwtAccessExpire

	expire := time.Duration(expireTime) * time.Second
	return timenow.Add(expire).Unix()
}

// freshTokenExpireAtTime 过期时间
func (jwt *JWT) freshTokenExpireAtTime() int64 {
	timenow := timenowInTimezone()

	expireTime := app.Config.Http.JwtRefreshTokenExpire

	expire := time.Duration(expireTime) * time.Second
	return timenow.Add(expire).Unix()
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// getTokenFromHeader 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.DefaultQuery("access_token", "")
	if len(authHeader) == 0 {
		authHeader = c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			return "", ErrHeaderEmpty
		}
	}
	//// 按空格分割
	//parts := strings.SplitN(authHeader, " ", 2)
	//if !(len(parts) == 2 && parts[0] == "Bearer") {
	//	return "", ErrHeaderMalformed
	//}
	return authHeader, nil
}

func timenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(time.Local.String())
	return time.Now().In(chinaTimezone)
}
