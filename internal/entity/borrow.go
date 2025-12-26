package entity

import (
	"time"
	"github.com/google/uuid"
)

type BorrowingRecord struct {
	BorrowId   uuid.UUID  `json:"borrow_id" gorm:"primaryKey"`
	UserId     uuid.UUID  `json:"user_id"`
	BookId     uuid.UUID  `json:"book_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
}

func NewBorrowingRecord(userId, bookId uuid.UUID) *BorrowingRecord {
	return &BorrowingRecord{
		BorrowId:   uuid.New(),
		UserId:     userId,
		BookId:     bookId,
		BorrowedAt: time.Now(),
	}
}