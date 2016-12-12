package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetIndexRenderTodos(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.LoadHTMLGlob("../template/*")
	r.GET("/", GetIndex)

	todos := []string{
		"Go Buy Milk",
		"Go Make Apps",
		"Go Sleepy",
	}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	for _, todo := range todos {
		if !strings.Contains(res.Body.String(), todo) {
			t.Fatal("Expect response body to contain: ", todo)
		}
	}
}

func TestGetLoginRenderLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.LoadHTMLGlob("../template/*")
	r.GET("/login", GetLogin)

	req, _ := http.NewRequest("GET", "/login", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if !strings.Contains(res.Body.String(), "Login") {
		t.Fatal("Expect response body to render login")
	}
}
