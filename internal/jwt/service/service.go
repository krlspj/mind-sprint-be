package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userID, username string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateToken() (jwtClaims, error)
}

type jwtClaims struct {
	Username      string `json:"username,omitempty"`
	UserID        string `json:"id,omitempty"`
	StandarClaims jwt.Claims
}

type defJwtService struct {
	secret               string
	tokenLiftTime        int
	refreshTokenLifeTime int
}

func NewJwtService(hmacSampleSecret string, tkTime, refreshTkTime int) defJwtService {
	return defJwtService{
		secret:               hmacSampleSecret,
		tokenLiftTime:        tkTime,
		refreshTokenLifeTime: refreshTkTime,
	}
}

func (s *defJwtService) GenerateToken(aUserID, aUsername string) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(5) * time.Minute)),
		Subject:   aUserID,
	}
	//		UserID:   aUserID,
	//		Username: aUsername,
	//		StandarClaims: jwt.RegisteredClaims{
	//
	//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(5) * time.Minute)),
	//			//ExpiresAt: time.Now().Add(time.Duration(5) * time.Minute).Unix(),
	//		},
	//	}

	return jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(s.secret))

}
