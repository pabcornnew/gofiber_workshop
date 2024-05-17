package routes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	//* วิธีสร้าง Group API
	api := app.Group("/api")
	// v1
	v1 := api.Group("/v1")
	// ! Result : /api/v1

	// * [Middleware && Basic Authentication]
	// Provide a minimal config
	// ! ต้องไว้ด้านบนเสมอ
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			//5_0
			"gofiber": "21022566", // * username && password
		},
	}))

	v1.Get("/", controllers.HelloWorld)

	v1.Post("/", controllers.PostNameandPass)

	// * [Params]
	v1.Get("/fact/:number", controllers.FactNum51)

	v1.Get("/user/:name", controllers.GetUserByName)

	// * [Query]
	v1.Post("/inet", controllers.QuerySearch)

	//* [Validation]
	v1.Post("/valid", controllers.PostStatus)
}
