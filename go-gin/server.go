package main

import (
	"golang-gin-poc/controller"
	"golang-gin-poc/middlewares"
	"golang-gin-poc/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func main() {
	logOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	subRouterAuthenticated := server.Group("/user", gin.BasicAuth(gin.Accounts{
		"admin": "adminpass",
	}))

	subRouterAuthenticated.GET("/get-users", func(ctx *gin.Context) {
		user, apiErr := userController.FindAll()
		if apiErr != nil {
			ctx.JSON(apiErr.StatusCode, apiErr)
		}
		ctx.JSON(200, user)
	})

	subRouterAuthenticated.GET("/get-user/:id", func(ctx *gin.Context) {
		user, apiErr := userController.FindUser(ctx)
		if apiErr != nil {
			ctx.JSON(apiErr.StatusCode, apiErr)
		}
		ctx.JSON(200, user)
	})

	subRouterAuthenticated.POST("/add-user", func(ctx *gin.Context) {
		res, apiErr := userController.Save(ctx)
		if apiErr != nil {
			ctx.JSON(apiErr.StatusCode, apiErr)
		}
		ctx.JSON(200, res)
	})

	subRouterAuthenticated.DELETE("/delete-user/:id", func(ctx *gin.Context) {

		res, apiErr := userController.DeleteRecord(ctx)
		if apiErr != nil {
			ctx.JSON(apiErr.StatusCode, apiErr)
		}
		ctx.JSON(200, res)

	})

	subRouterAuthenticated.PUT("/update-user/:id", func(ctx *gin.Context) {
		res, apiErr := userController.UpdateRecord(ctx)
		if apiErr != nil {
			ctx.JSON(apiErr.StatusCode, apiErr)
		}
		ctx.JSON(200, res)
	})

	server.Run(":8080")
}

func logOutput() {
	f, _ := os.Create("golang-gin-poc.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
