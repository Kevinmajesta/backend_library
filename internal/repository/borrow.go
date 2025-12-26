package repository

import (
	"github.com/Kevinmajesta/backend_library/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BorrowRepository interface {
	Create(tx *gorm.DB, borrow *entity.BorrowingRecord) error
	FindActiveBorrowByBookID(tx *gorm.DB, bookID uuid.UUID) (*entity.BorrowingRecord, error)
	FindByID(id uuid.UUID) (*entity.BorrowingRecord, error)
	Update(tx *gorm.DB, borrow *entity.BorrowingRecord) error
}

type borrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowRepository{db: db}
}

func (r *borrowRepository) Create(tx *gorm.DB, borrow *entity.BorrowingRecord) error {
	if tx != nil {
		return tx.Create(borrow).Error
	}
	return r.db.Create(borrow).Error
}

func (r *borrowRepository) FindActiveBorrowByBookID(
	tx *gorm.DB,
	bookID uuid.UUID,
) (*entity.BorrowingRecord, error) {

	var borrow entity.BorrowingRecord

	query := r.db
	if tx != nil {
		query = tx
	}

	err := query.
		Where("book_id = ? AND returned_at IS NULL", bookID).
		First(&borrow).Error

	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (r *borrowRepository) FindByID(id uuid.UUID) (*entity.BorrowingRecord, error) {
	var borrow entity.BorrowingRecord
	if err := r.db.Where("borrow_id = ?", id).First(&borrow).Error; err != nil {
		return nil, err
	}
	return &borrow, nil
}

func (r *borrowRepository) Update(tx *gorm.DB, borrow *entity.BorrowingRecord) error {
	if tx != nil {
		return tx.Save(borrow).Error
	}
	return r.db.Save(borrow).Error
}
