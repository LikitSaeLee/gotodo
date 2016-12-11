package main

import (
	"github.com/likitsaelee/gotodo/todo"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"strconv"
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
		newTodo := todo.Todo{ID: newID, Description: c.PostForm("todo"), Completed: false}
		todos = append(todos, newTodo)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/todo/delete", func(c *gin.Context) {
		id, err := strconv.Atoi(c.PostForm("id"))
		if err == nil {
			for i, todo := range todos {
				if id == todo.ID {
					// delete the element
					todos = append(todos[:i], todos[i+1:]...)
				}
			}
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/todo/complete", func(c *gin.Context) {
		id, err := strconv.Atoi(c.PostForm("id"))
		if err == nil {
			for i, todo := range todos {
				if id == todo.ID {
					todos[i].Completed = true
				}
			}
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	log.Println("Running on port 8080...")
	r.Run()
}
