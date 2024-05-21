package controllers

import (
	"fmt"
	database "go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	match, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, p.Username)
	if !match {
		return c.Status(fiber.StatusBadRequest).SendString("ชื่อผู้ใช้งานต้องประกอบด้วยตัวอักษร a-z, A-Z, ตัวเลข 0-9, หรือเครื่องหมาย _ หรือ - เท่านั้น")
	}

	if len(p.Password) < 6 || len(p.Password) > 20 {
		return c.Status(fiber.StatusBadRequest).SendString("ความยาวของรหัสผ่านต้องมากกว่า 6 และไม่เกิน 20 ตัวอักษร")
	}

	if len(p.Phon) < 10 {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณากรอกเบอร์โทรศัพท์ให้ถูกต้อง (อย่างน้อย 10 ตัวอักษร)")
	}

	if p.Business_type == "" {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณากรอกประเภทธุรกิจ")
	}

	if p.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณากรอกชื่อเว็บไซต์")
	}

	log.Println(p.Username)
	log.Println(p.Password)
	log.Println(p.Line)
	log.Println(p.Phon)
	log.Println(p.Business_type)
	log.Println(p.Url)

	str := fmt.Sprintf("Email: %s\nUsername: %s\nPassword: %s\nLine: %s\nPhon: %s\nBusiness Type: %s\nUrl: %s",
		p.Email, p.Username, p.Password, p.Line, p.Phon, p.Business_type, p.Url)

	return c.JSON(fiber.Map{
		"message": str,
	})
}

func DogIDGreaterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func ShowDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)

	return c.Status(200).JSON(dogs)
}

// 7_0
func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs // Adjust model name according to your project
	var dataResults []m.DogsRes

	// Fetch dogs from the database
	db.Find(&dogs) // Assuming this retrieves 10 dogs

	// Process each dog
	for _, v := range dogs {
		typeStr := ""
		switch v.DogID {
		case 111:
			typeStr = "red"
		case 113:
			typeStr = "green"
		case 999:
			typeStr = "pink"
		default:
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	// Create the result data
	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs),
	}

	// Return the result as JSON
	return c.Status(200).JSON(r)
}

// 7_2
func GetDogJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes

	colorCounters := map[string]int{
		"red":      0,
		"green":    0,
		"pink":     0,
		"no color": 0,
	}
	//1 inet 112 //2 inet1 113
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		colorCounters[typeStr]++

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:         dataResults,
		Name:         "golang-test",
		Count:        len(dogs), // Total number of dogs
		RedCount:     colorCounters["red"],
		GreenCount:   colorCounters["green"],
		PinkCount:    colorCounters["pink"],
		NoColorCount: colorCounters["no color"],
	}
	return c.Status(200).JSON(r)
}

// 7_1
func GetDogsScope(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	// Use the scope to filter the dogs
	db.Scopes(m.DogIDBetween50And100).Find(&dogs)

	return c.Status(200).JSON(dogs)

}

// 7_2

// company
// Create a new company
func CreateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&company)
	return c.Status(201).JSON(company)
}

// Get all companies
func GetAllCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company
	db.Find(&company)
	return c.Status(200).JSON(company)
}

// Get a single company by ID
func ReadSomeCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Company

	result := db.Find(&company, "com_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(company)

}

// Update a company
func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

// Delete a company
func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// Profile
// Create
func CreateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&profile)
	fmt.Println(profile)
	return c.Status(201).JSON(profile)
}

// Read
func GetAllProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile []m.Profile
	db.Find(&profile)

	return c.Status(200).JSON(profile)
}

func ReadSomeProfile(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	fmt.Println(search)
	var profile []m.Profile
	// result := db.Find(&company, "com_id = ?", search)
	result := db.Find(&profile, "emp_id = ? or name like ? or last_name like ?", search, search, search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(profile)
}

// Update
func UpdateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.Profile
	id := c.Params("id")

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&profile)
	return c.Status(200).JSON(profile)
}

// Delete
func RemoveProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile m.Profile

	result := db.Delete(&profile, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// Json
func GetJsonProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile []m.Profile

	db.Find(&profile)

	var dataResults []m.GetUserProfile

	genType := map[string]int{
		"GenZ":            0,
		"GenY":            0,
		"GenX":            0,
		"Baby Boomer":     0,
		"G.I. Generation": 0,
	}

	for _, v := range profile {
		typeStr := ""
		if v.Age < 24 {
			typeStr = "GenZ"
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "GenY"
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "GenX"
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Baby Boomer"
		} else if v.Age > 75 {
			typeStr = "G.I. Generation"
		} else {
			typeStr = "error"
		}

		genType[typeStr]++

		emp := m.GetUserProfile{
			EmpID:    v.EmpID,
			Name:     v.Name,
			LastName: v.LastName,
			Age:      v.Age,
			Type:     typeStr,
		}
		dataResults = append(dataResults, emp)
	}

	res := m.ResProfile{
		Data:       dataResults,
		Name:       "golang-test",
		Count:      len(profile), // Total number of dogs
		GenZ:       genType["GenZ"],
		GenX:       genType["GenX"],
		GenY:       genType["GenY"],
		BabyBoomer: genType["Baby Boomer"],
		GI:         genType["G.I. Generation"],
	}
	// return c.Status(200).JSON(res)
	return c.Status(200).JSON(res)
}

/*
func GetDogJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes

	colorCounters := map[string]int{
		"red":      0,
		"green":    0,
		"pink":     0,
		"no color": 0,
	}
	//1 inet 112 //2 inet1 113
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		colorCounters[typeStr]++

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:         dataResults,
		Name:         "golang-test",
		Count:        len(dogs), // Total number of dogs
		RedCount:     colorCounters["red"],
		GreenCount:   colorCounters["green"],
		PinkCount:    colorCounters["pink"],
		NoColorCount: colorCounters["no color"],
	}
	return c.Status(200).JSON(r)
}
*/
