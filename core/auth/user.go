package auth

import jwt "github.com/dgrijalva/jwt-go"

type User struct {
	UserID  int    `json:"userId"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	jwt.StandardClaims
}
