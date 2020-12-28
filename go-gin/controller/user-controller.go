package controller

import (
	"golang-gin-poc/entity"
	"golang-gin-poc/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) entity.User
	DeleteRecord(id int) []entity.User
	UpdateRecord(ctx *gin.Context) []entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Save(user)
	return user
}

func (c *controller) DeleteRecord(id int) []entity.User {

	return c.service.DeleteRecord(id)
}

func (c *controller) UpdateRecord(ctx *gin.Context) []entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	return c.service.UpdateRecord(user)
}
