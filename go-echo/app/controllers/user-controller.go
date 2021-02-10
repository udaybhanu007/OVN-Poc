package controllers

import (
	"encoding/json"
	"fmt"
	"go-echo-poc/app/model"
	"go-echo-poc/app/services"
	"go-echo-poc/app/utils"
	"net/http"
	"strconv"
	"sync"

	"github.com/gocql/gocql"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	fmt.Print("creating user")
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, "INVALID", "Invalid Json")
	}
	resultMessage := new(utils.NotificationMessage)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		result, saveError := services.UsersService.CreateUser(*user)
		if saveError != nil {
			return
		}
		fmt.Println(result)
		resultMessage = result
	}()

	go utils.ProduceKafkaForUserCreate(user)

	wg.Wait()

	return c.JSON(http.StatusCreated, resultMessage)
}

func GetUser(c echo.Context) error {
	UUID, uidErr := gocql.ParseUUID(c.Param("id"))
	if uidErr != nil {
		services.LogActivity(c.Param("id"), "GET: /get-user/:id", strconv.Itoa(http.StatusBadRequest), "Invalid UUID")
		return utils.NewHTTPError(http.StatusBadRequest, "INVALID", "Invalid UUID")
	}

	user, getErr := services.UsersService.GetUser(UUID)
	if getErr != nil {
		services.LogActivity(c.Param("id"), "GET: /get-user/:id", strconv.Itoa(http.StatusBadRequest), getErr.Error())
		return getErr
	}
	services.LogActivity(c.Param("id"), "GET: /get-user/:id", strconv.Itoa(http.StatusOK), "Success! User found.")
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	UUID, uidErr := gocql.ParseUUID(c.Param("uid"))
	if uidErr == nil {
		var user model.User
		err := json.NewDecoder(c.Request().Body).Decode(&user)
		if err != nil {
			return err
		}
		_, getErr := services.UsersService.GetUser(UUID)
		if getErr != nil {
			return getErr
		}
		user.Uuid = UUID

		result, updateErr := services.UsersService.UpdateUser(user)
		if updateErr != nil {
			return updateErr
		}
		return c.JSON(http.StatusOK, result)
	} else {
		return utils.NewHTTPError(http.StatusBadRequest, "INVALID", "Invalid UUID")
	}
}

func DeleteUser(c echo.Context) error {
	UUID, uidErr := gocql.ParseUUID(c.Param("uid"))
	if uidErr != nil {
		return utils.NewHTTPError(http.StatusBadRequest, "INVALID", "Invalid UUID")
	}
	fmt.Println("delete controller start")
	user, getErr := services.UsersService.DeleteUser(UUID)
	if getErr != nil {
		return getErr
	}
	return c.JSON(http.StatusOK, user)
}
