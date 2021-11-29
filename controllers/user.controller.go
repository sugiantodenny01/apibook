package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sugiantodenny01/apibook/models"
	"strconv"
	"time"
)


func FetchAllUserfunc(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims);

	result, err  := models.FetchAllUser()
	if err != nil {
		return c.JSON(map[string]string{"message" : err.Error()})
	}

	return c.JSON(result)
}

func StoreUser(c *fiber.Ctx) error  {
	name:=c.FormValue("name")
	address:=c.FormValue("address")
	telp:=c.FormValue("telp")
	password :=c.FormValue("password")
	role := c.FormValue("role")

	passwordHash,err := models.HashPassword(password)

	if err != nil {
		return c.JSON(err.Error())
	}


	//string to int
	roleInt,err := strconv.Atoi(role)

	if err != nil {
		return c.JSON(err.Error())
	}


	result, err := models.StoreUser(name,address,telp, passwordHash, roleInt)

	if err != nil {
		return c.JSON(map[string]string{"message" : err.Error()})
	}

	return c.JSON(result)

}

func UpdateUser(c *fiber.Ctx) error  {
	id := c.FormValue("id")
	name:=c.FormValue("name")
	address:=c.FormValue("address")
	telp:=c.FormValue("telp")
	password := c.FormValue("password")
	role:= c.FormValue("role")


	//string to int
	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(err.Error())
	}

	//string to int
	convertRole,err := strconv.Atoi(role)

	if err != nil {
		return c.JSON(err.Error())
	}


	passwordHash,err :=models.HashPassword(password)

	if err != nil {
		return c.JSON(err.Error())
	}


	result, err := models.UpdateUser(conv_id,name,address,telp,passwordHash,convertRole)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(result)

}

func DeleteUser(c *fiber.Ctx)error  {
	id := c.FormValue("id")

	//string to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(err.Error())
	}

	result, err := models.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(result)

}

func LoginUser(c *fiber.Ctx)error  {
	username := c.FormValue("name")
	password := c.FormValue("password")

	res, obj, err := models.LoginUser(username, password)
	if err != nil {

		//error default fiber -> tidak dikembalikan status 200
		//code := fiber.StatusUnauthorized
		//c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		//c.Status(code).SendString(err.Error())
		//

		return c.JSON(map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return fiber.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["role"]     = obj.Role
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON( map[string]string{"messages": err.Error()})
	}

	return c.JSON(map[string]string{
		"token": t,
	})

}

