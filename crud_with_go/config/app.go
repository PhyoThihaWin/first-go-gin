package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")

	dsn := "dbeaver:dbeaver@tcp(localhost:3306)/showcase_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Mysql connected!!")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
