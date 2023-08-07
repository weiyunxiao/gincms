package sys

import (
	"gincms/app"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PostService = new(postService)

type postService struct {
}

// AddPost 添加岗位操作
func (p *postService) AddPost(c *gin.Context, req *types.PostAddSaveReq) (err error) {
	addData := &model.SysPost{
		PostCode: req.PostCode,
		PostName: req.PostName,
		Sort:     req.Sort,
		Creator:  c.GetInt64("uid"),
		CreateTime: carbon.DateTime{
			Carbon: carbon.Now(),
		},
		Updater: c.GetInt64("uid"),
		UpdateTime: carbon.DateTime{
			Carbon: carbon.Now(),
		},
	}
	err = app.DB().Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// UpdatePost 更新岗位操作
func (p *postService) UpdatePost(c *gin.Context, req *types.PostAddSaveReq) (err error) {
	var updateData = map[string]any{
		"post_code": req.PostCode,
		"post_name": req.PostName,
		"sort":      req.Sort,
		"status":    req.Status,
		"updater":   c.GetInt64("uid"),
		"update_time": carbon.DateTime{
			Carbon: carbon.Now(),
		},
	}
	err = app.DB().Model(&model.SysPost{}).Where("id=?", req.Id).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// GetPost 查询一条岗位
func (p *postService) GetPost(c *gin.Context, req *typescom.IDReq) (post *model.SysPost, err error) {
	post = &model.SysPost{}
	err = app.DB().Select("id,post_code,post_name,sort,status,create_time").Where("id=?", req.ID).Take(&post).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// DelPost 根据id数组删除岗位
func (p *postService) DelPost(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysPost{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// PostPage 获取岗位分页数据
func (p *postService) PostPage(c *gin.Context, req *types.PostPageReq) (total int64, list []model.SysPost, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc)
	list = make([]model.SysPost, 0)

	query := app.DB().Model(&model.SysPost{}).Where("deleted=0")
	if len(req.PostName) > 0 {
		query.Where("post_name like ?", "%"+req.PostName+"%")
	}
	if len(req.PostCode) > 0 {
		query.Where("post_code like ?", "%"+req.PostCode+"%")
	}
	if len(req.Status) > 0 {
		query.Where("status=?", cast.ToInt(req.Status))
	}

	err = query.Count(&total).Select("id,post_code,post_name,sort,status").Order(sortStr).
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).
		Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// PostList 查询岗位列表数据
func (p *postService) PostList(c *gin.Context) (list []model.SysPost, err error) {
	list = make([]model.SysPost, 0)
	sortStr := pkg.SortStr("", true)

	err = app.DB().Where("deleted=0 and status=1").Select("id,post_code,post_name,sort,status,create_time").Order(sortStr).
		Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
