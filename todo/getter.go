package todo

// Todos is a list of Todo
var Todos []Todo

// GetMockData returns a slice of fake todos
func GetMockData() []Todo {
	if len(Todos) == 0 {
		Todos = []Todo{
			Todo{ID: 1, Description: "Go Buy Milk", Completed: false},
			Todo{ID: 2, Description: "Go Make Apps", Completed: false},
			Todo{ID: 3, Description: "Go Sleepy", Completed: false},
		}
	}
	return Todos
}

// SetMockData set todo data
func SetMockData(todos []Todo) {
	Todos = todos
}

// GetTodo given an int id return a todo with corresponding id.
func GetTodo(id int) Todo {
	for _, todo := range GetMockData() {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}
