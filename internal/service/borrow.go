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

const MaxBorrowQuota = 3

func (s *borrowService) BorrowBook(userID, bookID uuid.UUID) (*entity.BorrowingRecord, error) {

	var result *entity.BorrowingRecord

	err := s.db.Transaction(func(tx *gorm.DB) error {
		//saya pakai locking agar tidak terjadi race condition saat menguabah stok buku
		//takutnya sistem lagi banyak yang pakai dan terjadi race condition
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
		//ini saya buat var baru untuk menghitung jumlah max peminjaman aktif
		//terus saya cek apakah sudah melebihi kuota atau belum, kalau belum bisa lanjut, klo belum return error
		var activeBorrowCount int64
		if err := tx.Model(&entity.BorrowingRecord{}).
			Where("user_id = ? AND returned_at IS NULL", userID).
			Count(&activeBorrowCount).Error; err != nil {
			return err
		}

		if activeBorrowCount >= MaxBorrowQuota {
			return ErrBorrowQuotaExceeded
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

		//locking sama kaya di atas, takut race condition
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
		//ngembalikan stok buku
		book.Stock++
		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		return nil
	})
}
