package routes

import (
	ctlV1 "mercury/app/controllers/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		signupCtl := new(ctlV1.SignupController)
		v1.POST("/auth/signup/phone/exist", signupCtl.IsPhoneExist)
		v1.POST("/auth/signup/email/exist", signupCtl.IsEmailExist)
		v1.POST("/auth/signup/using-phone", signupCtl.SignupUsingPhone)
		v1.POST("/auth/signup/using-email", signupCtl.SignupUsingEmail)

		// 发送验证码
		vcCtl := new(ctlV1.VerifyCodeController)
		// 图片验证码，需要加限流
		v1.POST("/auth/verify-codes/captcha", vcCtl.ShowCaptcha)
		v1.POST("/auth/verify-codes/phone", vcCtl.SendUsingPhone)
		v1.POST("/auth/verify-codes/email", vcCtl.SendUsingEmail)
	}
}
