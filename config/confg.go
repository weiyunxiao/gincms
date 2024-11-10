package config

type Config struct {
	App          App  `mapstructure:"App"`
	Http         Http `mapstructure:"Http"`
	PathInfo     PathInfo
	AdminCaptcha AdminCaptcha  `mapstructure:"AdminCaptcha"`
	DBDefineList map[string]DB `mapstructure:"DBDefineList"`
	Redis        Redis         `mapstructure:"Redis"`
}
type App struct {
	Name                string   `mapstructure:"Name"`
	Env                 string   `mapstructure:"Env"`
	IsDemo              bool     `mapstructure:"IsDemo"`              //演示展示不可以更改数据
	IsDemoWhiteList     []string `mapstructure:"IsDemoWhiteList"`     //演示允许访问的白名单
	IsOpenRecordOperate bool     `mapstructure:"IsOpenRecordOperate"` //是否记录用户非get操作到记录表中

	LogPath            string `mapstructure:"LogPath"`
	OssType            string `mapstructure:"OssType"`
	UploadDir          string `mapstructure:"UploadDir"`
	UploadMaxM         int64  `mapstructure:"UploadMaxM"`         //文件上传限制多少M
	LogParamShowClient bool   `mapstructure:"LogParamShowClient"` //参数绑定错误是否显示回前端
	LogParamErr        bool   `mapstructure:"LogParamErr"`        //是否记录参数绑定错误到文件
	OpenRedis          bool   `mapstructure:"OpenRedis"`          //是否开启redis
}

type Http struct {
	WebSiteUrl            string `mapstructure:"WebSiteUrl"`
	UseCrossMiddleware    bool   `mapstructure:"UseCrossMiddleware"` //是否开启跨域中间件，nginx解决的话此可关闭
	AdminApiPort          string `mapstructure:"AdminApiPort"`
	JwtAccessSecret       string `mapstructure:"JwtAccessSecret"`
	JwtAccessExpire       int64  `mapstructure:"JwtAccessExpire"`
	JwtRefreshTokenExpire int64  `mapstructure:"JwtRefreshTokenExpire"`
}

type PathInfo struct {
	Root string
}
type DB struct {
	Dsn             string `mapstructure:"Dsn"`
	MaxIdleConns    int    `mapstructure:"MaxIdleConns"`
	MaxOpenConns    int    `mapstructure:"MaxOpenConns"`
	ConnMaxIdleTime int    `mapstructure:"ConnMaxIdleTime"`
	ConnMaxLifetime int    `mapstructure:"ConnMaxLifetime"`
}

type AdminCaptcha struct {
	CharLen         int `mapstructure:"CharLen"`
	ImgWidth        int `mapstructure:"ImgWidth"`
	ImgHeight       int `mapstructure:"ImgHeight"`
	LimitCaptchaNum int `mapstructure:"LimitCaptchaNum"`
	LimitTimeOut    int `mapstructure:"LimitTimeOut"` //秒
}
type Redis struct {
	Host  string `mapstructure:"Host"`
	Port  int    `mapstructure:"Port"`
	Pwd   string `mapstructure:"Pwd"`
	Dbuse int    `mapstructure:"Dbuse"`
}
