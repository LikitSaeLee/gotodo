package server

import (
	"bytes"
	"github.com/likitsaelee/gotodo/todo"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	todo.Todos = []todo.Todo{}
	os.Exit(m.Run())
}

func TestPostAddTodoShouldAddTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.LoadHTMLGlob("../template/*")
	r.POST("/todo/add", PostAddTodo)

	v := url.Values{}
	v.Set("todo", "Go Learn Programming")

	req, _ := http.NewRequest("POST", "/todo/add", bytes.NewBufferString(v.Encode()))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	todos := todo.GetMockData()

	if todos[len(todos)-1].Description != "Go Learn Programming" {
		t.Fatal("Expect to add new todo, but got:", todos)
	}
}
