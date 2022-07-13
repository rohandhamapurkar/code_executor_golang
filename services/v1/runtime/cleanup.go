package runtime

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

func cleanupExecution(execInfo executionInfo) {

}

func cleanupProcesses(execInfo executionInfo) {
	log.Println("Cleaning up processes")

	processIds, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Println(err)
		return
	}

	for _, proc := range processIds {
		procName := proc.Name()
		procId, err := strconv.Atoi(procName)
		if err != nil {
			continue
		}
		fmt.Println("procId", procId)

		data, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", procId))
		if err != nil {
			log.Println(err)
			continue
		}
		procInfo := string(data)

		r, err := regexp.Compile(`(?im)Uid:\s.*`)
		if err != nil {
			log.Println(err)
			continue
		}

		uids := strings.Fields(r.FindString(procInfo))

		ruid, err := strconv.Atoi(uids[1])
		if err != nil {
			continue
		}
		euid, err := strconv.Atoi(uids[2])
		if err != nil {
			continue
		}

		r, err = regexp.Compile(`(?im)State:\s.*`)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("State", "pid", procId, r.FindString(procInfo))

		r, err = regexp.Compile(`(?im)Gid:\s.*`)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("pid", procId, "gid", r.FindString(procInfo))

		if execInfo.Uid != uint32(ruid) && execInfo.Uid != uint32(euid) {
			log.Println("Not my process", "uid", ruid, "pid", procId)
			continue
		}

		log.Println("killing", "uid", ruid, "pid", procId)
		syscall.Kill(procId, syscall.SIGSTOP)
		syscall.Kill(procId, syscall.SIGKILL)

	}

	log.Println("Killed all my processes")

}
