package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/domenicomastrangelo/statsman/internal/database"
	"github.com/domenicomastrangelo/statsman/internal/handlers"
	"github.com/domenicomastrangelo/statsman/internal/handlers/authenticate"
	"github.com/gorilla/mux"
	"golang.org/x/term"
)

func main() {
	startFlag := flag.Bool("start", false, "Start the statsman daemon")
	addUserFlag := flag.Bool("user", false, "Adds a user to the database")

	flag.Parse()

	if *addUserFlag {
		creds := database.Credentials{}
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Username: ")
		creds.Username, _ = reader.ReadString('\n')

		fmt.Printf("Password: ")
		pass, err := term.ReadPassword(0)

		if err != nil {
			fmt.Println("There's been an error reading the password")
			return
		}

		creds.Password = string(pass)
		fmt.Println()

		creds.Username = strings.TrimSuffix(creds.Username, "\n")
		creds.Password = strings.TrimSuffix(creds.Password, "\n")

		database.AddUser(creds)
	}

	if *startFlag {
		router()
	}

	if !*startFlag && !*addUserFlag {
		flag.Usage()
	}
}

func router() {
	r := mux.NewRouter()
	rr := r.NewRoute().Subrouter()
	rr.Use(authenticate.Validate)

	rr.HandleFunc("/processes", handlers.ProcessesList).Methods("GET")
	rr.HandleFunc("/processes/details/{pid}", handlers.ProcessDetails).Methods("GET")
	rr.HandleFunc("/memory", handlers.Meminfo).Methods("GET")
	rr.HandleFunc("/disks", handlers.DisksSpace).Methods("GET")

	r.HandleFunc("/auth", authenticate.Auth).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
