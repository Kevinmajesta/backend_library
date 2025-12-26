package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	BookId uuid.UUID `gorm:"type:uuid;primaryKey" json:"book_id"`
	Title  string    `json:"title"`
	Stock  int       `json:"stock"`
	Auditable
}

func NewBook(title string, stock int) *Book {
	return &Book{
		BookId:    uuid.New(),
		Title:     title,
		Stock:     stock,
		Auditable: NewAuditable(),
	}
}
