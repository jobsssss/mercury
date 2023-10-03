package v1

import (
	"github.com/gin-gonic/gin"
	"mercury/app/models/user"
	"mercury/app/requests"
	"mercury/pkg/response"
)

type SignupController struct {
	BaseAPIController
}

func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.SignupPhoneExist); !ok {
		return
	}

	response.JSON(ctx, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
