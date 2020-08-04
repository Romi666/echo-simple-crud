package controllers

import (
	"github.com/labstack/echo"
	"../models"
	"net/http"
	"strconv"
)

func FetchAllItem(c echo.Context) error{
	result, err := models.FetchAllProducts()
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchById(c echo.Context) error{
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	result, err := models.FetchById(idConv)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateItem(c echo.Context) error{
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return err
	}
	result, err := models.CreateItems(item.Name, item.Price)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateItem(c echo.Context) error{
	id := c.Param("id")
	name := c.FormValue("name")
	price := c.FormValue("price")
	idConv ,_ := strconv.Atoi(id)
	priceConv, _ := strconv.Atoi(price)

	result, err := models.UpdateItem(idConv, name, priceConv)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteItem(c echo.Context) error{
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	result, err := models.DeleteItem(idConv)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}