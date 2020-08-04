package controllers

import (
	"../helpers"
	"github.com/labstack/echo"
	"net/http"
)

func GenerateHashPassword(c echo.Context)error{
	password := c.Param("password")

	result, err := helpers.HashPassword(password)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"Error" : err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"Token" : result})
}
