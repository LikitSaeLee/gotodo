package todo

import (
	"testing"
)

func TestGetMockData(t *testing.T) {
	expTodos := []Todo{
		Todo{ID: 1, Description: "Go Buy Milk", Completed: false},
		Todo{ID: 2, Description: "Go Make Apps", Completed: false},
		Todo{ID: 3, Description: "Go Sleepy", Completed: false},
	}

	todos := GetMockData()

	for i, todo := range todos {
		if todo.ID != expTodos[i].ID ||
			todo.Description != expTodos[i].Description ||
			todo.Completed != expTodos[i].Completed {

			t.Fatal("Expected to get todo, but got:", todo)
		}
	}
}

func TestSetMockData(t *testing.T) {
	expTodos := []Todo{
		Todo{ID: 0, Description: "Gorilla", Completed: false},
	}

	SetMockData(expTodos)

	todos := GetMockData()
	for i, todo := range todos {
		if todo.ID != expTodos[i].ID ||
			todo.Description != expTodos[i].Description ||
			todo.Completed != expTodos[i].Completed {

			t.Fatal("Expected to get todo, but got:", todo)
		}
	}

	Todos = []Todo{}
}

func TestGetTodo(t *testing.T) {
	todo := GetTodo(2)

	if todo.ID != 2 {
		t.Fatal("Expected to get todo, but got:", todo)
	}
}
