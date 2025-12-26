package binder

type UserCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}
