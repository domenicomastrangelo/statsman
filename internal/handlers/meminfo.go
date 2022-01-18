package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/domenicomastrangelo/statsman/internal/memory"
)

func Meminfo(w http.ResponseWriter, r *http.Request) {
	var (
		memData      map[string]string
		err          error
		jsonResponse []byte
	)

	if memData, err = memory.GetMemoryDetails(); err != nil {
		log.Println(err.Error())
		log.Println("Could not get memory info data")
		return
	}

	if jsonResponse, err = json.Marshal(memData); err != nil {
		log.Println(err.Error())
		log.Println("Could not make meminfo into json")
		return
	}

	w.Write(jsonResponse)
}
