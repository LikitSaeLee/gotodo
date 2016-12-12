package todo

// Todo is a task of user
type Todo struct {
	ID          int
	Description string
	Completed   bool
}

// SetCompleted set a Todo as completed.
func (todo *Todo) SetCompleted() {
	todo.Completed = true
}
