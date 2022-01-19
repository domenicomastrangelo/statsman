package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Credentials) CheckCredentials() bool {
	stmt, err := SqliteDatabase.Prepare("select password from users where username = ?")

	if err != nil {
		log.Println(err.Error())
		return false
	}

	res := stmt.QueryRow(c.Username)

	var hashedPass []byte
	err = res.Scan(&hashedPass)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(c.Password))
	log.Println(err)
	return err == nil
}
