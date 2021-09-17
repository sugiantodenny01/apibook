package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"log"
)

var IsAuthenticated = jwtware.New(jwtware.Config{
	SigningKey: []byte("secret"),
	ErrorHandler: jwtError,
})

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}



func CheckAdmin(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	log.Println(claims)

	role := claims["role"].(float64)

	if role == 1 {
		log.Println("halo")
		return	c.Next()
	}else{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized",
		})
	}

}