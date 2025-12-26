package service

import (
	"errors"
	"time"

	"github.com/Kevinmajesta/backend_library/internal/entity"
	"github.com/Kevinmajesta/backend_library/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BorrowService interface {
	BorrowBook(userID, bookID uuid.UUID) (*entity.BorrowingRecord, error)
	ReturnBook(borrowID uuid.UUID) error
}

type borrowService struct {
	db             *gorm.DB
	bookRepository repository.BookRepository
	borrowRepo     repository.BorrowRepository
}

func NewBorrowService(
	db *gorm.DB,
	bookRepository repository.BookRepository,
	borrowRepo repository.BorrowRepository,
) BorrowService {
	return &borrowService{
		db:             db,
		bookRepository: bookRepository,
		borrowRepo:     borrowRepo,
	}
}

func (s *borrowService) BorrowBook(userID, bookID uuid.UUID) (*entity.BorrowingRecord, error) {

	var result *entity.BorrowingRecord

	err := s.db.Transaction(func(tx *gorm.DB) error {

		var book entity.Book
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("book_id = ?", bookID).
			First(&book).Error; err != nil {
			return err
		}

		if book.Stock <= 0 {
			return ErrBookOutOfStock
		}

		book.Stock--
		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		borrow := entity.NewBorrowingRecord(userID, bookID)
		if err := s.borrowRepo.Create(tx, borrow); err != nil {
			return err
		}

		result = borrow
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *borrowService) ReturnBook(borrowID uuid.UUID) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		borrow, err := s.borrowRepo.FindByID(borrowID)
		if err != nil {
			return errors.New("borrow record not found")
		}

		if borrow.ReturnedAt != nil {
			return errors.New("book already returned")
		}

		//locking
		var book entity.Book
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("book_id = ?", borrow.BookId).
			First(&book).Error; err != nil {
			return err
		}

		now := time.Now()
		borrow.ReturnedAt = &now

		if err := s.borrowRepo.Update(tx, borrow); err != nil {
			return err
		}

		book.Stock++
		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		return nil
	})
}
