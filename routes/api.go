package routes

import (
	ctlv1 "mercury/app/controllers/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			signupCtl := new(ctlv1.SignupController)
			authGroup.POST("/signup/phone/exist", signupCtl.IsPhoneExist)
		}
	}
}
