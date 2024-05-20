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
type GetDogsJson struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

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
	Data  []DogsRes `json:"data"`
	Name  string    `json:"name"`
	Count int       `json:"count"`
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
