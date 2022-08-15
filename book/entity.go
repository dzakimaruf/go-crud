package book

import "time"

type Book struct {
	ID          int
	title       string
	Description string
	Price       int
	Rating      int
	createdAt   time.Time
	updatedAt   time.Time
}
