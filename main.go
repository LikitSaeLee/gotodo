package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.Run()
}
