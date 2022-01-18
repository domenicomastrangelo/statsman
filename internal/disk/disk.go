package disk

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type diskDetails map[string]string

func GetDiskDetails() ([]diskDetails, error) {
	var (
		df              exec.Cmd
		out             bytes.Buffer
		diskDetailsList []diskDetails = []diskDetails{}
	)

	df = *exec.Command("df")
	df.Stdout = &out

	if err := df.Run(); err != nil {
		log.Println(err.Error())
		return []diskDetails{}, fmt.Errorf("could not run df command on system")
	}

	for i, str := range strings.Split(string(out.String()), "\n") {
		keyVal := strings.Fields(str)

		if i == 0 {
			continue
		}

		if len(keyVal) > 0 {
			diskDetailsList = append(diskDetailsList, diskDetails{
				"filesystem":     keyVal[0],
				"1k_blocks":      keyVal[1],
				"used":           keyVal[2],
				"available":      keyVal[3],
				"use_percentage": keyVal[4],
				"mounted_on":     keyVal[5],
			})
		}
	}

	return diskDetailsList, nil
}
