package routes

import (
	ctlV1 "mercury/app/controllers/v1"
	"mercury/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 注册用户
			signupCtl := new(ctlV1.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), signupCtl.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), signupCtl.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), signupCtl.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), signupCtl.SignupUsingEmail)

			// 发送验证码
			vcCtl := new(ctlV1.VerifyCodeController)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcCtl.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcCtl.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcCtl.SendUsingEmail)

			//登录
			lgCtl := new(ctlV1.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgCtl.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgCtl.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgCtl.RefreshToken)

			// 重置密码
			pwdCtl := new(ctlV1.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwdCtl.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwdCtl.ResetByEmail)
		}
	}
}
