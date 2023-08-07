package types

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}
type SysDictTypeResp struct {
	DictType string `json:"dictType"`
	DataList []struct {
		DictLabel  string `json:"dictLabel"`
		DictValue  string `json:"dictValue"`
		LabelClass string `json:"labelClass"`
	} `json:"dataList"`
}
