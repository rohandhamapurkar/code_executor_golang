package runtime

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func cleanupExecution() {

}

func cleanupProcesses() error {
	// processes := []int{1}
	log.Println("Cleaning up processes")

	processIds, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Println(err)
		return err
	}

	for _, proc := range processIds {
		procId := proc.Name()
		if _, err := strconv.Atoi(procId); err != nil {
			continue
		}
		fmt.Println(procId)

		dat, err := os.ReadFile(fmt.Sprintf("/proc/%s/status", procId))
		if err != nil {
			continue
		}
		fmt.Print(string(dat))

	}

	return nil
}
