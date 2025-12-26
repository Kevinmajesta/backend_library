package binder

import "github.com/google/uuid"

type BorrowRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	BookID uuid.UUID `json:"book_id" validate:"required"`
}
