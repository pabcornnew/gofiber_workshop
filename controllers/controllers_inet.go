package controllers

import (
	m "go-fiber-test/models"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func PostNameandPass(c *fiber.Ctx) error {
	p := new(m.Person) // เก็บข้อมูล

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error")
		// return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func GetUserByName(c *fiber.Ctx) error {
	var str string = "hello : " + c.Params("name")
	return c.JSON(str)
}

func QuerySearch(c *fiber.Ctx) error {
	a := c.Query("search") // ถ้า Search มาจาก หน้าบ้าน จะเก็บใน a
	str := "my search is  " + a
	return c.JSON(str)
}

// 5_1
func FactNum51(c *fiber.Ctx) error {
	num, err := c.ParamsInt("number")
	log.Printf("number is %v", num)
	if err != nil {
		return c.JSON(err)
	}
	res := 1

	for i := 1; i <= num; i++ {
		res *= i
	}
	return c.JSON(res)
}

func PostStatus(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	// * nil = null

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	//  * แบบลดรูป
	// if errors := validate.Struct(user); errors != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	// }

	return c.JSON(user)
}
