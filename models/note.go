package models

import (
	"time"
)

type Note struct {
	ID         int       `schema:"-"`
	Title      string    `schema:"title"`
	Body       string    `schema:"body"`
	Tags       string    `schema:"tags"`
	NotebookID int       `schema:"notebook_id"`
	CreatedAt time.Time  `schema:"-"`
	UpdatedAt time.Time  `schema:"-"`
}
