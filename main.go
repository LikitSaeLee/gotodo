package main

import (
	"fmt"
	"github.com/likitsaelee/gotodo/todo"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	todos := todo.GetMockData()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Todos": todos,
		})
	})

	r.POST("/todo/add", func(c *gin.Context) {
		newID := todos[len(todos)-1].ID + 1
		newTodo := todo.Todo{ID: newID, Description: c.PostForm("todo")}
		todos = append(todos, newTodo)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	fmt.Printf("Running on port 8080...")
	r.Run()
}
