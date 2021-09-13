package middleware

import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetToken(username string) (string) {
	signingKey := []byte("secret-library")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expired": time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("secret-library")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}