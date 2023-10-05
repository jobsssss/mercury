package bootstrap

import (
	"fmt"
	"mercury/app/middlewares"
	"mercury/pkg/config"
	"mercury/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Serve() {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)
	// gin 实例
	router := gin.New()
	setupRoute(router)
	err := router.Run(":" + config.Get("app.port", "3000"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func setupRoute(router *gin.Engine) {

	registerGlobalMiddleWare(router)

	routes.RegisterAPIRoutes(router)

	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "404 Not found")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "The request path not defined, please confirm url or request method is correct",
			})
		}
	})
}
