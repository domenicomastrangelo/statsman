package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/domenicomastrangelo/statsman/internal/processes"
	"github.com/gorilla/mux"
)

func ProcessesList(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		jsonResponse []byte
	)

	processesList := processes.GetProcessesList()

	if jsonResponse, err = json.Marshal(processesList); err != nil {
		log.Println(err.Error())
		log.Println("Could not get list of processes")
		return
	}

	w.Write(jsonResponse)
}

func ProcessDetails(w http.ResponseWriter, r *http.Request) {
	var (
		pid          int
		err          error
		params       map[string]string
		jsonResponse []byte
	)

	params = mux.Vars(r)

	if pid, err = strconv.Atoi(params["pid"]); err != nil {
		log.Printf("Could not convert process PID to int. %s\n", err.Error())
		return
	}

	if details, err := processes.GetProcessDetails(pid); err != nil {
		log.Println(err.Error())
		log.Printf("Could not get details of PID %d\n", pid)
	} else {
		if jsonResponse, err = json.Marshal(details); err != nil {
			log.Println(err.Error())
			log.Printf("Could not marshal JSON for details of PID %d\n", pid)
			return
		}

		w.Write(jsonResponse)
	}
}
