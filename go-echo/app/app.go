package app

import (
	"go-echo-poc/app/controllers"
	"go-echo-poc/app/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	e = echo.New()
)

func StartApplication() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code = http.StatusInternalServerError
			key  = "ServerError"
			msg  string
		)

		if he, ok := err.(*helpers.HttpError); ok {
			code = he.Code
			key = he.Key
			msg = he.Message
		} else {
			msg = err.Error()
		}

		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD {
				err := c.NoContent(code)
				if err != nil {
					c.Logger().Error(err)
				}
			} else {
				err := c.JSON(code, helpers.NewHTTPError(code, key, msg))
				if err != nil {
					c.Logger().Error(err)
				}
			}
		}

	}

	e.POST("/create-user", controllers.CreateUser)
	e.GET("/get-user/:id", controllers.GetUser)
	e.PUT("/update-user/:uid", controllers.UpdateUser)
	e.DELETE("/delete-user/:uid", controllers.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
