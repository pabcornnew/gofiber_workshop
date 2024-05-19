package controllers

import (
	"fmt"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"

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

// 5_2
func AsciiQuery(c *fiber.Ctx) error {
	queryParam := c.Query("tax_id")

	var result string
	var text string

	for _, v := range queryParam {
		result = strconv.Itoa(int(v))
		text += " " + result
	}
	// str := "my search is  " + queryParam
	return c.JSON(text)
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

// 6
func Register(c *fiber.Ctx) error {
	p := new(m.Register)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Validate username
	match, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, p.Username)
	if !match {
		return c.Status(fiber.StatusBadRequest).SendString("ชื่อผู้ใช้งานต้องประกอบด้วยตัวอักษร a-z, A-Z, ตัวเลข 0-9, หรือเครื่องหมาย _ หรือ - เท่านั้น")
	}

	// Validate password length
	if len(p.Password) < 6 || len(p.Password) > 20 {
		return c.Status(fiber.StatusBadRequest).SendString("ความยาวของรหัสผ่านต้องมากกว่า 6 และไม่เกิน 20 ตัวอักษร")
	}

	// Validate phone number length
	if len(p.Phon) < 10 {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณากรอกเบอร์โทรศัพท์ให้ถูกต้อง (อย่างน้อย 10 ตัวอักษร)")
	}

	// Validate business type
	if p.Business_type == "" {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณากรอกประเภทธุรกิจ")
	}

	// Validate URL
	if p.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณากรอกชื่อเว็บไซต์")
	}

	// Log data
	log.Println(p.Username)
	log.Println(p.Password)
	log.Println(p.Line)
	log.Println(p.Phon)
	log.Println(p.Business_type)
	log.Println(p.Url)

	// Respond with success
	str := fmt.Sprintf("Email: %s\nUsername: %s\nPassword: %s\nLine: %s\nPhon: %s\nBusiness Type: %s\nUrl: %s",
		p.Email, p.Username, p.Password, p.Line, p.Phon, p.Business_type, p.Url)

	return c.JSON(fiber.Map{
		"message": str,
	})
}
