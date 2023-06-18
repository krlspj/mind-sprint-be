package inmemory

import (
	"strconv"

	"github.com/krlspj/mind-sprint-be/internal/auth/domain"
)

type userDB struct {
	id       int
	name     string
	password string
	email    string
}

func (u userDB) toDomainUser() domain.User {
	return domain.User{
		ID:       strconv.Itoa(u.id),
		Name:     u.name,
		Password: u.password,
		Email:    u.email,
	}
}

func toUserDB(user domain.User) (userDB, error) {
	var id int
	if user.ID == "" {
		id = 0
	} else {
		var err error
		id, err = strconv.Atoi(user.ID)
		if err != nil {
			return userDB{}, err
		}
	}

	return userDB{
		id:       id,
		name:     user.Name,
		password: user.Password,
		email:    user.Email,
	}, nil
}

func toDomainUserList(users []userDB) []domain.User {
	result := make([]domain.User, 0, len(users))
	for _, u := range users {
		result = append(result, u.toDomainUser())
	}
	return result
}
