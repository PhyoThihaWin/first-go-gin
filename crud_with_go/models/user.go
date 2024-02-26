package models

import (
	"example.com/m/crud_with_go/config"
	"gorm.io/gorm"
)

type User struct {
	CustomModel
	USERNAME  string `json:"username"`
	DEVICE_ID string `json:"device_id"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateBook() *User {
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) *User {
	var user User
	db.Where("id=?", ID).Delete(&user)
	return &user
}
