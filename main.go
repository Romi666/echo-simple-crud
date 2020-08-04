package main

import (
	"./db"
	"./routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.PUT, echo.GET, echo.DELETE, echo.POST},
	}))
	e.Logger.Fatal(e.Start(":3000"))
}
