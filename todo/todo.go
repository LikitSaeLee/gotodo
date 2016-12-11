package todo

// Todos is a list of Todo
var Todos []Todo

// Todo is a task of user
type Todo struct {
	ID          int
	Description string
	Completed   bool
}

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
