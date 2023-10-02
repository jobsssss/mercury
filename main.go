package main

import (
	"flag"
	"fmt"
	"mercury/bootstrap"
	btsConfig "mercury/config"
	"mercury/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "load .env, such as --env=testing, load are .env.testing")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()

	router := gin.New()
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port", "3000"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
