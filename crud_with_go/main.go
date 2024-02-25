package main

import (
	"net/http"

	"example.com/m/crud_with_go/config"
	"example.com/m/crud_with_go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.User{})

	router := gin.Default()
	router.GET("/user/list", GetUsers)
	router.Run("localhost:9090")
}

func GetUsers(context *gin.Context) {
	var Users []models.User
	db.Find(&Users)
	// return Users

	context.IndentedJSON(http.StatusOK, models.DataResponse[[]models.User]{Data: Users})
}
