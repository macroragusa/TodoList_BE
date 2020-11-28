package models

// todoModel describes a todoModel type
type Todos struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// CreateTodoInput forces data validation for input because create need only one field
type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}
