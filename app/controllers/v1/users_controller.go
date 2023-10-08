package v1

import (
	"mercury/app/models/user"
	"mercury/app/requests"
	"mercury/pkg/auth"
	"mercury/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(ctx *gin.Context) {
	userModel := auth.CurrentUser(ctx)
	response.Data(ctx, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}
	data, pager := user.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}
