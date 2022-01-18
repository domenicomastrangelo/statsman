package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/domenicomastrangelo/statsman/internal/disk"
)

func DisksSpace(w http.ResponseWriter, r *http.Request) {
	var jsonResponse []byte

	if diskData, err := disk.GetDiskDetails(); err != nil {
		log.Println(err.Error())
		log.Println("Could not get disks info data")
		return
	} else if jsonResponse, err = json.Marshal(diskData); err != nil {
		log.Println(err.Error())
		log.Println("Could not make meminfo into json")
		return
	}

	w.Write(jsonResponse)
}
