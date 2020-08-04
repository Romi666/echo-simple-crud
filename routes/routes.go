package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"../controllers"
	"../middleware"
)

func Init() *echo.Echo{
	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello there")
	})
	e.GET("/item", controllers.FetchAllItem, middleware.IsAuthenticated)
	e.GET("/item/:id", controllers.FetchById)
	e.POST("/item", controllers.CreateItem)
	e.PUT("/item/:id", controllers.UpdateItem)
	e.DELETE("/item/:id", controllers.DeleteItem)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/generate-hash", controllers.CheckLogin)
	return e
}