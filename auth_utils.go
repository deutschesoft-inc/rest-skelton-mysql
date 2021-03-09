package main

import (
	"main/config"

	jwt "github.com/dgrijalva/jwt-go"
)

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(config.Secret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
