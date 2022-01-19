package authenticate

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/domenicomastrangelo/statsman/internal/database"
	"github.com/golang-jwt/jwt"
)

var key = []byte("thisisthekey")

func authenticateJWT(creds database.Credentials) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = creds.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		log.Println("There was an error signing the key")
		log.Println(err.Error())
		return "", fmt.Errorf("an error occurred")
	}

	return tokenString, nil
}

func validateJWT(tokenFromRequest string) (string, error) {
	jwtToken, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Println("Error parsing token")
		log.Println(err.Error())
		return "", fmt.Errorf("error parsing token")
	}

	claims := jwtToken.Claims.(jwt.MapClaims)

	var tm time.Time
	switch t := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(t), 0)
	case json.Number:
		v, _ := t.Int64()
		tm = time.Unix(v, 0)
	}

	if time.Since(tm).Minutes() > -5 {
		return refreshJWTToken(claims)
	}

	return tokenFromRequest, nil
}

func refreshJWTToken(claims jwt.MapClaims) (string, error) {
	return authenticateJWT(database.Credentials{Username: claims["username"].(string)})
}
