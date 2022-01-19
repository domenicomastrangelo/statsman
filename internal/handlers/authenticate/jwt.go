package authenticate

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("thisisthekey")

func authenticateJWT(creds credentials) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = creds.Username
	claims["exp"] = time.Now().Add(time.Minute * 30)

	tokenString, err := token.SignedString(key)

	if err != nil {
		log.Println("There was an error signing the key")
		log.Println(err.Error())
		return "", fmt.Errorf("an error occurred")
	}

	return tokenString, nil
}
