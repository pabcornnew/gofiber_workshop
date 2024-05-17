package main

import (
	"go-fiber-test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // * ใช้ libary (fiber)
	routes.InetRoutes(app)
	/*
		! Fiber.Ctx = Context
		* เช่น BodyParser

		BodyParser = การรับค่าข้อมูลมาจากทาง KeyBorad
	*/

	/* (Status Code)
	! 400 : Fail
	! 401 : Token is invalid
	! 402 : Login is true but you not use this page (Access denied)
	! 403 : Record not found
	! 404 : Not found

	* 200 : Success
	* 201 : Create Success

	* 500 :  Internal Server Error
	* 502 : Sever Fail
	*/

	// * กำหนด Post (เปิด Server)
	app.Listen(":3000")
}
