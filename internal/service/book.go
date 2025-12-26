package service

import (
	"errors"

	"github.com/Kevinmajesta/backend_library/internal/entity"
	"github.com/Kevinmajesta/backend_library/internal/repository"
	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(book *entity.Book) (*entity.Book, error)
	BookExists(title string) bool
	FindBookByID(id string) (*entity.Book, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) *bookService {

	return &bookService{
		bookRepository: bookRepository,
	}
}

func (s *bookService) CreateBook(book *entity.Book) (*entity.Book, error) {
	if book.Title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if book.Stock == 0 {
		return nil, errors.New("stock cannot be zero or less")
	}

	newBook, err := s.bookRepository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return newBook, nil
}

func (s *bookService) BookExists(title string) bool {
	_, err := s.bookRepository.FindBookByTitle(title)
	return err == nil
}

func (s *bookService) FindBookByID(id string) (*entity.Book, error) {
	book, err := s.bookRepository.FindBookByID(uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return book, nil
}
