package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/apibook/controllers"
	"github.com/sugiantodenny01/apibook/middleware"
)

func Init()  *fiber.App {
	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	/*route user */
	r.Get("/users",middleware.IsAuthenticated,middleware.CheckAdmin,controllers.FetchAllUserfunc)
	r.Post("/users",middleware.CheckAdmin,controllers.StoreUser)
	r.Put("/users",middleware.CheckAdmin,controllers.UpdateUser)
	r.Delete("/users",middleware.CheckAdmin,controllers.DeleteUser)

	/*route book */
	r.Get("/books",controllers.FetchAllBookfunc)
	r.Post("/books",controllers.StoreBook)
	r.Put("/books",controllers.UpdateBook)
	r.Delete("/boooks",controllers.DeleteBook)

	/*route login */
	r.Post("/login", controllers.LoginUser)


	return r

}