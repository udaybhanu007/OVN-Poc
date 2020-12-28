package main

import (
	"fmt"
	"golang-gin-poc/controller"
	"golang-gin-poc/middlewares"
	"golang-gin-poc/service"
	"io"
	"os"
	"strconv"

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

	server.GET("/get-user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/add-user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	server.DELETE("/delete-user/:id", func(ctx *gin.Context) {
		fmt.Println(strconv.Atoi(ctx.Param("id")))
		id, err := strconv.Atoi(ctx.Param("id"))
		if err == nil {
			fmt.Println("inside")
			ctx.JSON(200, userController.DeleteRecord(id))
		}
		fmt.Println(err)
	})

	server.PUT("/update-user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.UpdateRecord(ctx))
	})

	server.Run(":8080")
}

func logOutput() {
	f, _ := os.Create("golang-gin-poc.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
