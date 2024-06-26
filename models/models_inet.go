package models

import "gorm.io/gorm"

// * เส้น POST (ส่งข้อมูล) [BodyParser : การรับข้อมูลในรูปแบบ JSON]
type Person struct {
	//ส่วนเก็บข้อมูลไว้ใน GO
	Name string `json:"name"`
	Pass string `json:"pass"`
}

// *Connect to database
type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

// Regiter
type Register struct {
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Line          string `json:"line"`
	Phon          string `json:"phon"`
	Business_type string `json:"business"`
	Url           string `json:"url"`
}

// 7_2
type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data         []DogsRes `json:"data"`
	Name         string    `json:"name"`
	Count        int       `json:"count"`
	RedCount     int       `json:"red_count"`
	GreenCount   int       `json:"green_count"`
	PinkCount    int       `json:"pink_count"`
	NoColorCount int       `json:"no_color_count"`
}

// 7_0
type Company struct {
	gorm.Model
	ComID   int    `json:"com_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

// 7_1
func DogIDBetween50And100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ? AND dog_id < ?", 50, 100)
}

// Final Project
type Profile struct {
	gorm.Model
	EmpID     int    `json:"employee_id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	BirthDay  string `json:"birthday"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Telephone string `json:"tel"`
}

type GetUserProfile struct {
	EmpID     int    `json:"employee_id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	BirthDay  string `json:"birthday"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Telephone string `json:"tel"`
	Type      string `json:"type"`
}

type ResProfile struct {
	Data       []GetUserProfile `json:"data"`
	Name       string           `json:"name"`
	Count      int              `json:"count"`
	GenZ       int              `json:"genz"`
	GenX       int              `json:"genx"`
	GenY       int              `json:"geny"`
	BabyBoomer int              `json:"babyboomer"`
	GI         int              `json:"gi"`
}
