package service

import (
	"context"
	"errors"
	"log"

	"github.com/krlspj/mind-sprint-be/internal/auth/domain"
)

var (
	ErrUsernameAlreadyExist = errors.New("username already exist")
	ErrEmailAlreadyExist    = errors.New("email already exists")
)

type AuthService interface {
	CreateUser(ctx context.Context, user domain.User) (string, error)
	CheckNewUserConflicts(ctx context.Context, user domain.User) error
	Login(ctx context.Context, user domain.User) (domain.User, error)
}

type authService struct {
	userRepo domain.UserRepo
}

func NewAuthService(ur domain.UserRepo) *authService {
	return &authService{
		userRepo: ur,
	}
}

func (s *authService) CreateUser(ctx context.Context, user domain.User) (string, error) {
	log.Println("user pass:", user.Password)
	hashedPass, err := user.HashPassword()
	if err != nil {
		return "", err
	}
	user.Password = hashedPass
	log.Println("user Hashed pass", user)
	log.Println("--------", s.userRepo.FindAll(ctx))

	return s.userRepo.CreateUser(ctx, user)
}

func (s *authService) CheckNewUserConflicts(ctx context.Context, user domain.User) error {
	// Find user with the same username, or same email.
	// if user found, do not allow to create user with the same value

	users := s.userRepo.FindAll(ctx)
	for _, u := range users {
		if u.Name == user.Name {
			return ErrUsernameAlreadyExist

		} else if u.Email == user.Email {
			return ErrEmailAlreadyExist
		}
	}

	return nil

}

func (s *authService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{ID: "137", Name: "user0"}, nil
}
