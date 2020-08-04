package controllers

import (
	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func CheckLogin(c echo.Context) error{
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message" : err.Error()})
	}

	if !res {
		return echo.ErrUnauthorized
	}
	//Create token
	token := jwt.New(jwt.SigningMethodHS256)

	//Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token" : t})
}