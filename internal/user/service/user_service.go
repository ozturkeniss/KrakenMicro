package service

import (
	"context"
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gomicro/internal/user/model"
	"gomicro/internal/user/repository"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id uint) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint) error

	// For test compatibility
	GetUser(ctx context.Context, id uint) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	existingUser, err := s.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return s.repo.Create(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	existingUser, err := s.repo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.repo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	existingUser, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.repo.Delete(ctx, id)
}

func (s *userService) GetUser(ctx context.Context, id uint) (*model.User, error) {
	return s.GetUserByID(ctx, id)
} 