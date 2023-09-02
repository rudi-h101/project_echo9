package dto

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Id  int `json:"id"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
