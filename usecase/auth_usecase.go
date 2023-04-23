package usecase

import (
	"main/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const SECRET_JWT = "nFyLUyG5aUUkjGE+vk8/BDxKwAG/j5hJ+Io72HArK7k="

type AuthUsecase interface {
	CreateToken(payload dto.CreateToken) (string, error)
	ExtractTokenUserId(payload echo.Context) int
	ExtractTokenAdminId(payload echo.Context) int
}

type authUsecase struct{}

func NewAuthUsecase() *authUsecase {
	return &authUsecase{}
}

func (s *authUsecase) CreateToken(payload dto.CreateToken) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"id":         payload.Id,
		"user_type":  payload.UserType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_JWT))

}

func (s *authUsecase) ExtractTokenUserId(payload echo.Context) int {
	user := payload.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		if claims["user_type"] == "user" {
			id := claims["id"].(float64)
			return int(id)
		}
		return 0

	}
	return 0
}

func (s *authUsecase) ExtractTokenAdminId(payload echo.Context) int {
	user := payload.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		if claims["user_type"] == "admin" {
			id := claims["id"].(float64)
			return int(id)
		}
		return 0
	}
	return 0

}
