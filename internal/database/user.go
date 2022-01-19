package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var SqliteDatabase *sql.DB

func init() {
	file, err := os.Open("/usr/local/var/statsman/db.db")

	if err != nil {
		os.Mkdir("/usr/local/var/statsman/", 0770)
		file, err = os.Create("/usr/local/var/statsman/db.db")

		if err != nil {
			log.Println(err.Error())
			log.Fatalln("Could not open database file")
			return
		}
	}

	if file != nil {
		file.Close()
	} else {
		log.Fatalln("Could not create file /usr/local/var/statsman/db.db")
		return
	}

	SqliteDatabase, _ = sql.Open("sqlite3", "/usr/local/var/statsman/db.db")

	createUsersTable()
}

func createUsersTable() {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" VARCHAR(255) UNIQUE,
		"password" VARCHAR(255)
	  );`

	log.Println("Create users table...")
	statement, err := SqliteDatabase.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("users table created")
}

func AddUser(creds Credentials) {
	log.Println("Inserting user record ...")
	insertStudentSQL := `INSERT INTO users(username, password) VALUES (?, ?)`
	statement, err := SqliteDatabase.Prepare(insertStudentSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var pass []byte
	pass, err = bcrypt.GenerateFromPassword([]byte(creds.Password), 14)
	creds.Password = string(pass)

	if err != nil {
		log.Fatalln("Could not hash user password")
	}

	_, err = statement.Exec(creds.Username, creds.Password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
