package routes

import (
	"go-fiber-test/controllers"
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	//* วิธีสร้าง Group API
	api := app.Group("/api")
	// v1
	v1 := api.Group("/v1")

	// v3
	v3 := api.Group("/v3")
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

	v1.Get("/", c.HelloWorld)

	v1.Post("/", c.PostNameandPass)

	// * [Params]
	v1.Get("/fact/:number", c.FactNum51)

	v1.Get("/user/:name", c.GetUserByName)

	// * [Query]
	v1.Post("/inet", c.QuerySearch)

	//* [Validation]
	v1.Post("/valid", c.PostStatus)

	//5_2
	v3.Post("/pab", c.AsciiQuery)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", controllers.GetDogs)
	dog.Get("/filter", controllers.GetDog)
	dog.Get("/json", controllers.GetDogsJson)
	dog.Post("/", controllers.AddDog)
	dog.Put("/:id", controllers.UpdateDog)
	dog.Delete("/:id", controllers.RemoveDog)

	//CRUD company
	company := v1.Group("/company")
	company.Get("", controllers.GetAllCompany)
	company.Get("/filter", controllers.ReadSomeCompany)
	company.Post("/", controllers.CreateCompany)
	// company.Put("/:e_id", controllers.UpdateCompany)
}
