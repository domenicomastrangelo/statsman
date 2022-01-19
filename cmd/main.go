package main

import (
	"net/http"

	"github.com/domenicomastrangelo/statsman/internal/handlers"
	"github.com/domenicomastrangelo/statsman/internal/handlers/authenticate"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(authenticate.Validate)

	r.HandleFunc("/processes", handlers.ProcessesList).Methods("GET")
	r.HandleFunc("/processes/details/{pid}", handlers.ProcessDetails).Methods("GET")
	r.HandleFunc("/memory", handlers.Meminfo).Methods("GET")
	r.HandleFunc("/disks", handlers.DisksSpace).Methods("GET")

	r.HandleFunc("/auth", authenticate.Auth).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
