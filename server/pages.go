package server

import (
	"github.com/likitsaelee/gotodo/todo"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

// GetIndex get the index page
func GetIndex(c *gin.Context) {
	todos := todo.GetMockData()
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"Todos": todos,
	})
}

// GetLogin return login page
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", nil)
}
