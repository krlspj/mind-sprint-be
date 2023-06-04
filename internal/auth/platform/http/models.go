package handler

import "github.com/krlspj/mind-sprint-be/internal/auth/domain"

type userLoginReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u userLoginReq) toDomainUser() domain.User {
	return domain.User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
	}
}

type userRegisterReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u userRegisterReq) toDomainUser() domain.User {
	return domain.User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
	}
}

type userLoginResp struct {
}
