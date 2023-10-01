package v1

import (
	"fmt"
	"mercury/app/models/user"
	"mercury/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	BaseAPIController
}

func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		fmt.Println(err.Error())

		return
	}
	errs := requests.ValidateSignupPhoneExist(&request, ctx)
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
