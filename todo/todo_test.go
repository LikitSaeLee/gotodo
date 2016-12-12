package todo

import (
	"testing"
)

func TestSetCompleted(t *testing.T) {
	todo := GetTodo(2)
	todo.SetCompleted()

	if todo.Completed != true {
		t.Fatal("Expect todo to be completed, but got:", todo)
	}
}
