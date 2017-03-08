package models

import (
	"time"
)

type Note struct {
	ID         int       `schema:"-"`
	Title      string    `schema:"title"`
	Body       string    `schema:"body"`
	Tags       string    `schema:"tags"`
	CategoryID int       `schema:"category_id"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}
