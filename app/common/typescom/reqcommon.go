package typescom

import "mime/multipart"

type IDReq struct {
	ID int64 `json:"id" form:"id"`
}

type IDArrReq struct {
	IDArr []int64 `json:"id_arr" form:"id_arr"`
}

// PageOrderCommonReq 分页及排序请求
type PageOrderCommonReq struct {
	Order string `json:"order" form:"order"`
	Asc   bool   `json:"asc" form:"asc"`
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
}

// FileUploadReq 上传单个文件请求
type FileUploadReq struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required"`
}
