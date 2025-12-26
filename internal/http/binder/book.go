package binder

type BookCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}
