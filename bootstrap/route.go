package bootstrap

import (
	"mercury/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	registerGlobalMiddleWare(router)

	routes.RegisterAPIRoutes(router)

	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
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
