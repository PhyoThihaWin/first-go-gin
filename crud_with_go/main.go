package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/crud_with_go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", GetUsers)
	router.GET("/user/:id", GetUserById)
	router.POST("/user", CreateUser)
	router.DELETE("/user/:id", DeleteUser)
	router.Run("localhost:9090")
}

func GetUsers(context *gin.Context) {
	users := models.GetAllUsers()
	context.IndentedJSON(http.StatusOK, models.DataResponse[[]models.User]{Data: users})
}

func GetUserById(context *gin.Context) {
	userId := context.Param("id")
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetail, _ := models.GetUserById(ID)
	context.IndentedJSON(http.StatusOK, models.DataResponse[models.User]{Data: *userDetail})
}

func CreateUser(context *gin.Context) {
	CreateUser := &models.User{}
	if err := context.BindJSON(&CreateUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, models.DataResponse[string]{Data: "Bad Error"})
	} else {
		data := CreateUser.CreateBook()
		context.IndentedJSON(http.StatusOK, models.DataResponse[models.User]{Data: *data})
	}
}

func DeleteUser(context *gin.Context) {
	userId := context.Param("id")
	fmt.Print("userId", userId)
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetail := models.DeleteUser(ID)
	context.IndentedJSON(http.StatusOK, models.DataResponse[models.User]{Data: *userDetail})
}

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var updateBook = &models.Book{}
// 	utils.ParseBody(r, updateBook)
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	bookDetails, db := models.GetBookById(ID)
// 	if updateBook.Name != "" {
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != "" {
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != "" {
// 		bookDetails.Publication = updateBook.Publication
// 	}
// 	db.Save(&bookDetails)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
