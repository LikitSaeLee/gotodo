package server

import (
	"github.com/likitsaelee/gotodo/todo"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
)

// PostAddTodo add a todo
func PostAddTodo(c *gin.Context) {
	todos := todo.GetMockData()
	newID := todos[len(todos)-1].ID + 1
	newTodo := todo.Todo{ID: newID, Description: c.PostForm("todo"), Completed: false}
	todo.SetMockData(append(todos, newTodo))
	c.Redirect(http.StatusMovedPermanently, "/")
}

// PostDeleteTodo delete todo
func PostDeleteTodo(c *gin.Context) {
	todos := todo.GetMockData()
	id, err := strconv.Atoi(c.PostForm("id"))
	if err == nil {
		for i, todoItem := range todos {
			if id == todoItem.ID {
				// delete the element
				todo.SetMockData(append(todos[:i], todos[i+1:]...))
			}
		}
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

// PostCompleteTodo complete todo
func PostCompleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.String(http.StatusNotFound, "Unable to find todo", nil)
	}
	todoItem := todo.GetTodo(id)
	todoItem.SetCompleted()

	c.Redirect(http.StatusMovedPermanently, "/")
}
