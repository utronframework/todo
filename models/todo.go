package models

import (
	"time"
)

//Todo represent an item of todo list
type Todo struct {
	ID        int       `schema:"-"`
	Body      string    `schema:"body"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}
