package typescom

// PageDataResp 分页返回数据
type PageDataResp struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}
