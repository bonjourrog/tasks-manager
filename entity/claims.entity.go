package entity

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserName string `json:"user"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
