package model

import "github.com/dgrijalva/jwt-go"


type Claims struct {
	Username           string `json:"username"`
	IsAdmin            bool   `json:"isAdmin"`
	jwt.StandardClaims `swaggerignore:"true"`
}
