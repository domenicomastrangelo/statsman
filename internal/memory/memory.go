package memory

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type memory map[string]string

func GetMemoryDetails() (map[string]string, error) {
	var (
		mem     memory = memory{}
		memData []byte
		err     error
	)

	if memData, err = os.ReadFile("/proc/meminfo"); err != nil {
		log.Println(err.Error())
		return map[string]string{}, fmt.Errorf("could not read /proc/meminfo")
	}

	for _, str := range strings.Split(string(memData), "\n") {
		keyVal := strings.Split(str, ":")

		if len(keyVal) > 1 {
			mem[strings.Trim(keyVal[0], "\t")] = strings.ReplaceAll(strings.TrimSpace(keyVal[1]), "\t", " ")
		} else if len(strings.Trim(keyVal[0], "\t")) > 0 {
			mem[strings.Trim(keyVal[0], "\t")] = "-"
		}
	}

	return mem, nil
}
