package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krlspj/mind-sprint-be/internal/auth/service"
)

type AuthHandler struct {
	authService service.AuthService

	// user service
	// jwt service
}

func NewAuthHandler(as service.AuthService) AuthHandler {
	return AuthHandler{
		authService: as,
	}
}

func (h AuthHandler) RegisterNewUser(c *gin.Context) {
	var userReq userRegisterReq

	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.authService.CheckNewUserConflicts(c.Request.Context(), userReq.toDomainUser())
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	uID, err := h.authService.CreateUser(c.Request.Context(), userReq.toDomainUser())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, uID)

}

func (h AuthHandler) Login(c *gin.Context) {
	var userLogin userLoginReq

	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authService.Login(c.Request.Context(), userLogin.toDomainUser())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
	// generate Token
	// generate refresh refresh token

}

func (h AuthHandler) Logout(c *gin.Context) {

}
