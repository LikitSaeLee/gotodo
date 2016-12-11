package todo

// Todo is a task of user
type Todo struct {
	ID          int
	Description string
}

// GetMockData returns a slice of fake todos
func GetMockData() []Todo {
	return []Todo{
		Todo{ID: 1, Description: "Go Buy Milk"},
		Todo{ID: 2, Description: "Go Make Apps"},
		Todo{ID: 3, Description: "Go Sleepy"},
	}
}
