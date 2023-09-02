package utils

import (
	"errors"
	"os"
	"r101/models"
	"r101/models/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetUserID(c echo.Context) (int, error) {
	authHeader := c.Request().Header.Get("Authorization")

	tokenString := authHeader[7:]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// Access the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Error accessing claims")
	}

	id := int(claims["id"].(float64))

	return id, nil
}

func Error(s string) {
	panic("unimplemented")
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