package entity

import (
	"github.com/google/uuid"
)

type User struct {
	UserId   uuid.UUID `json:"user_id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Auditable
}

func NewUser(fullname, email, phone string) *User {
	return &User{
		UserId:    uuid.New(),
		Fullname:  fullname,
		Email:     email,
		Phone:     phone,
		Auditable: NewAuditable(),
	}
}
