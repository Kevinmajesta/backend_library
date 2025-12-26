package service

import "errors"

var (
	ErrBookOutOfStock      = errors.New("stok buku habis")
	ErrBorrowQuotaExceeded = errors.New("kuota peminjaman sudah penuh")
)
