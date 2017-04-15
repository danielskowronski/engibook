package models

import (
	"time"
)

type Notebook struct {
	ID         int       `schema:"-"`
	Title      string    `schema:"title"`
	CreatedAt time.Time  `schema:"-"`
	UpdatedAt time.Time  `schema:"-"`
}
