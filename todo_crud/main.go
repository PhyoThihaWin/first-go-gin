package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID       string `json:"id"`
	Item     string `json:"item"`
	Complete bool   `json:"complete"`
}

type DataResponse[T any] struct {
	Data T `json:"item"`
}

var todoList = []Todo{
	{ID: "1", Item: "Clean Room", Complete: false},
	{ID: "2", Item: "Wash Clothes", Complete: false},
	{ID: "3", Item: "GoLang Homework", Complete: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, DataResponse[[]Todo]{Data: todoList})
	// context.IndentedJSON(http.StatusOK, map[string]interface{}{"data": todoList})
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, DataResponse[string]{Data: "Bad Error"})
	} else {
		todoList = append(todoList, newTodo)
		context.IndentedJSON(http.StatusOK, DataResponse[Todo]{Data: newTodo})
	}

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
