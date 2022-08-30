package main

import (
	"fmt"

	"github.com/bankonly/gogin-init/src/configs"
	"github.com/bankonly/goutils/middlewares"
	"github.com/bankonly/goutils/response"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	app := gin.New()

	app.Use(gin.Logger())

	app.Use(gin.CustomRecovery(middlewares.ErrorRecovery)) // Error handler from panic call

	app.GET("/", func(ctx *gin.Context) {
		res := response.New(ctx)
		res.Success(response.H{})
	})

	fmt.Printf("Server is running on port: %s", configs.Constant.PORT)

	app.Run(fmt.Sprintf(":%s", configs.Constant.PORT))
}
