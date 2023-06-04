package domain

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Name     string
	Password string
	Email    string
	Color    string
	Image    interface{}
}

func (u User) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) // this hash generation already uses a salt in it
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (u User) VerifyPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type UserRepo interface {
	CreateUser(ctx context.Context, user User) (string, error)
	FindAll(ctx context.Context) []User
}
