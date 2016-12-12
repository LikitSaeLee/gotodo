package router

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

func TestGetIndexReturnStatusOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := Load("../template/*")

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Expect response to have status 200.")
	}
}

func TestGetLoginReturnStatusOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := Load("../template/*")

	req, _ := http.NewRequest("GET", "/login", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Expect response to have status 200.")
	}
}

func TestPostAddTodoReturnStatusRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := Load("../template/*")

	v := url.Values{}
	v.Set("todo", "Go Learn Programming")

	req, _ := http.NewRequest("POST", "/todo/add", bytes.NewBufferString(v.Encode()))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusMovedPermanently {
		t.Fatal("Expect response to have status redirect.")
	}
}

func TestPostDeleteTodoReturnStatusRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := Load("../template/*")

	v := url.Values{}
	v.Set("id", "2")

	req, _ := http.NewRequest("POST", "/todo/delete", bytes.NewBufferString(v.Encode()))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusMovedPermanently {
		t.Fatal("Expect response to have status redirect.")
	}
}

func TestPostCompleteTodoReturnStatusRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := Load("../template/*")

	v := url.Values{}
	v.Set("id", "2")

	req, _ := http.NewRequest("POST", "/todo/complete", bytes.NewBufferString(v.Encode()))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusMovedPermanently {
		t.Fatal("Expect response to have status redirect, but got:", res.Code)
	}
}
