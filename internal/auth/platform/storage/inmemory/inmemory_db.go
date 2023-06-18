package inmemory

import (
	"context"
	"fmt"

	"github.com/krlspj/mind-sprint-be/internal/auth/domain"
)

type userRepoStub struct {
	users []userDB
}

func NewUserRepositoryStub() *userRepoStub {
	initUsers := []userDB{
		{1, "Carles", "12345", "krls@gmail.com"},
		{2, "Lais", "67890", "lais@gmail.com"},
	}
	return &userRepoStub{
		users: initUsers,
	}
}

func (aStub *userRepoStub) CreateUser(ctx context.Context, user domain.User) (string, error) {
	lastUserId := aStub.users[len(aStub.users)-1].id
	lastUserId++

	dbUser, err := toUserDB(user)
	if err != nil {
		return "", err
	}
	dbUser.id = lastUserId

	aStub.users = append(aStub.users, dbUser)

	return fmt.Sprint(lastUserId), nil

}

func (aStub *userRepoStub) FindAll(ctx context.Context) []domain.User {
	return toDomainUserList(aStub.users)
}
