package processes

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

type process map[string]string

func GetProcessesList() []process {
	var (
		processes   []process
		procs       []fs.DirEntry
		err         error
		procNameInt int
	)

	if procs, err = os.ReadDir("/proc"); err != nil {
		fmt.Println(err.Error())
		return []process{}
	}

	for _, p := range procs {
		if procNameInt, err = strconv.Atoi(p.Name()); err != nil {
			continue
		}

		var cmdLine []byte
		if cmdLine, err = os.ReadFile(fmt.Sprintf("/proc/%d/cmdline", procNameInt)); err != nil {
			log.Println(err.Error())
			log.Printf("Could not read file /proc/%d/cmdline", procNameInt)
		}

		processes = append(processes, process{
			"Pid":  fmt.Sprintf("%d", procNameInt),
			"Name": string(bytes.Trim(cmdLine, "\x00")),
		})

	}

	return processes
}

func GetProcessDetails(pid int) (map[string]string, error) {
	var (
		process    process = process{}
		procStatus []byte
		err        error
	)

	if procStatus, err = os.ReadFile(fmt.Sprintf("/proc/%d/status", pid)); err != nil {
		log.Println(err.Error())
		return map[string]string{}, fmt.Errorf("could not read status of PID %d", pid)
	}

	for _, str := range strings.Split(string(procStatus), "\n") {
		keyVal := strings.Split(str, ":")

		if len(keyVal) > 1 {
			process[strings.Trim(keyVal[0], "\t")] = strings.ReplaceAll(strings.TrimSpace(keyVal[1]), "\t", " ")
		} else if len(strings.Trim(keyVal[0], "\t")) > 0 {
			process[strings.Trim(keyVal[0], "\t")] = "-"
		}
	}

	return process, nil
}
