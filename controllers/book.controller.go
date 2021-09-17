package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/apibook/models"
	"strconv"
)

func FetchAllBookfunc(c *fiber.Ctx) error {
	result, err  := models.FetchAllBook()
	if err != nil {
		return c.JSON(map[string]string{"message" : err.Error()})
	}

	return c.JSON(result)
}

func StoreBook(c *fiber.Ctx) error  {
	name:=c.FormValue("name")
	author:=c.FormValue("author")
	stock:=c.FormValue("stock")
	price :=c.FormValue("price")

	//string to int
	stockInt,err := strconv.Atoi(stock)

	if err != nil {
		return c.JSON(err.Error())
	}

	//string to int
	priceInt,err := strconv.Atoi(price)

	if err != nil {
		return c.JSON(err.Error())
	}



	result, err := models.StoreBook(name,author,stockInt, priceInt)

	if err != nil {
		return c.JSON(map[string]string{"message" : err.Error()})
	}

	return c.JSON(result)

}

func UpdateBook(c *fiber.Ctx) error  {
	id := c.FormValue("id")
	name:=c.FormValue("name")
	author:=c.FormValue("author")
	stock:=c.FormValue("stock")
	price :=c.FormValue("price")

	//string to int
	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(err.Error())
	}

	//string to int
	stockInt,err := strconv.Atoi(stock)

	if err != nil {
		return c.JSON(err.Error())
	}

	//string to int
	priceInt,err := strconv.Atoi(price)

	if err != nil {
		return c.JSON(err.Error())
	}


	result, err := models.UpdateBook(conv_id,name,author,stockInt,priceInt)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(result)

}

func DeleteBook(c *fiber.Ctx)error  {
	id := c.FormValue("id")

	//string to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(err.Error())
	}

	result, err := models.DeleteBook(conv_id)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(result)

}

