package todo

import (
	context "context"
)

// Controller for todos
type Controller struct {
}

// Todo struct
type Todo struct {
	ID int `json:"id"`
}
