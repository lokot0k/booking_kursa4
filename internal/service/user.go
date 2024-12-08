package service

import (
	"errors"
	"meeting-room-booking/internal/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(id int) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Create(user domain.User) error
	Update(user domain.User) error
	Delete(id int) error
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetAllUsers() ([]domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *UserService) GetUserByID(id int) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserService) GetUserByUsername(username string) (*domain.User, error) {
	return u.userRepo.GetByUsername(username)
}

func (u *UserService) CreateUser(user domain.User) error {
	return u.userRepo.Create(user)
}

func (u *UserService) UpdateUser(user domain.User) error {
	return u.userRepo.Update(user)
}

func (u *UserService) DeleteUser(id int) error {
	return u.userRepo.Delete(id)
}

func (u *UserService) Authenticate(username, password string) (*domain.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
