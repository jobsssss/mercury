package main

import (
	"fmt"
	"mercury/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
