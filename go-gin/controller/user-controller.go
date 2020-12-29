package controller

import (
	"golang-gin-poc/entity"
	"golang-gin-poc/helpers"
	"golang-gin-poc/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() (*[]entity.User, *helpers.ApplicationError)
	FindUser(ctx *gin.Context) (*entity.User, *helpers.ApplicationError)
	Save(ctx *gin.Context) (*helpers.Notification, *helpers.ApplicationError)
	DeleteRecord(ctx *gin.Context) (*helpers.Notification, *helpers.ApplicationError)
	UpdateRecord(ctx *gin.Context) (*helpers.Notification, *helpers.ApplicationError)
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() (*[]entity.User, *helpers.ApplicationError) {
	return c.service.FindAll()
}

func (c *controller) FindUser(ctx *gin.Context) (*entity.User, *helpers.ApplicationError) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		return c.service.FindUser(id)
	}
	return nil, &helpers.ApplicationError{
		Message:    "UserId must be number",
		StatusCode: http.StatusBadRequest,
	}
}

func (c *controller) Save(ctx *gin.Context) (*helpers.Notification, *helpers.ApplicationError) {
	var user entity.User
	ctx.BindJSON(&user)
	return c.service.Save(user)
}

func (c *controller) DeleteRecord(ctx *gin.Context) (*helpers.Notification, *helpers.ApplicationError) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		return c.service.DeleteRecord(id)
	}
	return nil, &helpers.ApplicationError{
		Message:    "UserId must be number",
		StatusCode: http.StatusBadRequest,
	}

}

func (c *controller) UpdateRecord(ctx *gin.Context) (*helpers.Notification, *helpers.ApplicationError) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		var user entity.User
		ctx.BindJSON(&user)
		if user.ID != id {
			return nil, &helpers.ApplicationError{
				Message:    "UserId mismatch",
				StatusCode: http.StatusBadRequest,
			}
		}
		return c.service.UpdateRecord(user)
	}
	return nil, &helpers.ApplicationError{
		Message:    "UserId must be number",
		StatusCode: http.StatusBadRequest,
	}
}
