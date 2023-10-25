package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json: "id`
	Item      string `json: "item"`
	Completed bool   `json: "completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Buy milk", Completed: false},
	{ID: "2", Item: "Buy eggs", Completed: false},
	{ID: "3", Item: "Buy bread", Completed: false},
	{ID: "4", Item: "Buy cheese", Completed: false},
	{ID: "5", Item: "Buy butter", Completed: false},
	{ID: "6", Item: "Buy flour", Completed: false},
	{ID: "6", Item: "Buy candy", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}
func addTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodosById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo Not Found")

}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(context *gin.Context) {
	var updatedTodo Todo
	if err := context.BindJSON(&updatedTodo); err != nil {
		return
	}
	todo, err := getTodosById(updatedTodo.ID)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}
	todo.Item = updatedTodo.Item
	todo.Completed = updatedTodo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
func main() {
	router := gin.Default()
	go router.GET("/todos",  getTodos)
	go router.GET("/todos/:id", getTodo)
	go router.POST("/todos", addTodo)
	go router.PATCH("/todos", updateTodo)
	 router.Run("localhost:8080")
}
