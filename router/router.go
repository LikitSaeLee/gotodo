package router

import (
	"github.com/likitsaelee/gotodo/server"
	"gopkg.in/gin-gonic/gin.v1"
)

// Load loads all the routes
func Load() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", server.GetIndex)
	r.GET("/login", server.GetLogin)
	r.POST("/todo/add", server.PostAddTodo)
	r.POST("/todo/delete", server.PostDeleteTodo)
	r.POST("/todo/complete", server.PostCompleteTodo)

	return r
}
