package v1

import (
	"mercury/app/models/category"
	"mercury/app/requests"
	"mercury/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Store(ctx *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(ctx, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(ctx, categoryModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(ctx *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(ctx.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(ctx, &request, requests.CategorySave); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(ctx, categoryModel)
	} else {
		response.Abort500(ctx)
	}
}