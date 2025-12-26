package repository

import (
	"github.com/Kevinmajesta/backend_library/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindBookByID(id uuid.UUID) (*entity.Book, error)
	FindBookByTitle(title string) (*entity.Book, error)
	CreateBook(book *entity.Book) (*entity.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) FindBookByID(id uuid.UUID) (*entity.Book, error) {
	book := new(entity.Book)
	if err := r.db.Where("book_id = ?", id).Take(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *bookRepository) FindBookByTitle(title string) (*entity.Book, error) {
	book := new(entity.Book)
	if err := r.db.Where("title = ?", title).Take(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *bookRepository) CreateBook(book *entity.Book) (*entity.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
