package utils

import (
	"os"
	"prakerja9/models"
	"prakerja9/models/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserID(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JwtCustomClaims)
	return claims.Id, nil
}

func GenerateJwt(user models.User) (string, error) {
	claims := &dto.JwtCustomClaims{
		Id: user.ID,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}