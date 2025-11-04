package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json: "id"`
	Item      string `json: "item"`
	Completed bool   `json: "completed"` //this make struct to json
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record my movie", Completed: false},
}

func getTodos(context *gin.Context) { //context is all things from request, response
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	//declare initialize at the same time and if context is not fit on newTodo it will throw error
	//if it fit, newTodo get json
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

// (param) (return this, or error)
func getTodoById(id string) (*todo, error) {
	// for index, value := slice
	for index, t := range todos {
		if t.ID == id {
			//return value, and error which is nil
			return &todos[index], nil
		}
	}

	return nil, errors.New("todo not found")
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// front *behind mean front is behind's pointer
// go's return default setting is void don't need to write void or ()
// if curly bracket is one, it means param
func getTodo(context *gin.Context) {
	id := context.Param("id")
	//todo or error will get
	todo, err := getTodoById(id)
	//go doesn't use curly brackets in if condition
	if err != nil {
		//HashMap by gin
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
