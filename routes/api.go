package routes

import (
	ctlv1 "mercury/app/controllers/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		signupCtl := new(ctlv1.SignupController)
		v1.POST("/auth/signup/phone/exist", signupCtl.IsPhoneExist)
		v1.POST("/auth/signup/email/exist", signupCtl.IsEmailExist)

		// 发送验证码
		vcc := new(ctlv1.VerifyCodeController)
		// 图片验证码，需要加限流
		v1.POST("/auth/verify-codes/captcha", vcc.ShowCaptcha)
		v1.POST("/auth/verify-codes/phone", vcc.SendUsingPhone)
	}
}
