package service

import (
	"errors"

	"github.com/Kevinmajesta/backend_library/internal/entity"
	"github.com/Kevinmajesta/backend_library/internal/repository"
)

type UserService interface {
	CreateUser(user *entity.User) (*entity.User, error)
	EmailExists(email string) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {

	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if user.Fullname == "" {
		return nil, errors.New("fullname cannot be empty")
	}
	if user.Phone == "" {
		return nil, errors.New("phone cannot be empty")
	}

	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) EmailExists(email string) bool {
	_, err := s.userRepository.FindUserByEmail(email)
	return err == nil
}
