package v1

import (
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
	if ok := requests.Validate(ctx, &request, requests.SignupPhoneExist); !ok {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
