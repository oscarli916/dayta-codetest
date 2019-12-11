package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var port = ":3030"

var secretkey = os.Getenv("SECRET_KEY")

func main() {
	/*
		TODO:
		Implement a GRPC server that accepts a rpc call Decode Token.
		function verifyToken() is given below.
	*/
	fmt.Println("GRPC server running on port", port)
}

func verifyToken(token string) (map[string]interface{}, error) {
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretkey, nil
	})
	mappedClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Token claims map error")
	}
	return mappedClaims, nil
}
