package runtime

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func cleanupExecution(execInfo executionInfo) {
	err := os.RemoveAll(os.TempDir() + "/" + execInfo.Id)
	if err != nil {
		log.Println("Recursive deletion err", err)
	}
}

type parsedProcInfo struct {
	Ruid  int
	Euid  int
	State string
	Gid   int
}

func cleanupProcesses(signals chan os.Signal) {
	for {
		select {
		case <-signals:
			log.Println("Cleaning up processes")

			waitForPids := []int{}

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

				info, parsed := parseProcInfo(procInfo)
				if parsed == true {
					continue
				}

				if info.Ruid >= minRunnerUid && info.Ruid <= maxRunnerUid && info.Euid >= minRunnerUid && info.Euid <= maxRunnerUid {

					if info.State == "Z" {
						waitForPids = append(waitForPids, procId)
					} else {
						process, err := os.FindProcess(procId)
						if err != nil {
							log.Println("Couldn't find process id:", procId, err)
							continue
						}
						process.Kill()
					}
					log.Println("pid:", procId)
					log.Println("uid:", info.Ruid, info.Euid)
					log.Println("gid:", info.Gid)
					log.Println("state:", info.State)
				} else {
					log.Println("pid:", procId, "Not my process")
					continue
				}
			}

			for _, procId := range waitForPids {
				log.Println("waiting proc id", procId)
				var status unix.WaitStatus
				pid, err := unix.Wait4(procId, &status, unix.WNOHANG, nil)
				if err != nil {
					log.Println("Wait err:", err)
				}
				log.Println("Grandchild exited pid:", pid, status)
			}

			log.Println("Killed all my processes")
		}
	}

}

func parseProcInfo(procInfo string) (parsedProcInfo, bool) {
	r, err := regexp.Compile(`(?im)Uid:\s.*`)
	if err != nil {
		log.Println("Uid regex err", err)
		return parsedProcInfo{}, true
	}

	uids := strings.Fields(r.FindString(procInfo))

	ruid, err := strconv.Atoi(uids[1])
	if err != nil {
		log.Println("Ruid parse err", err)
		return parsedProcInfo{}, true
	}
	euid, err := strconv.Atoi(uids[2])
	if err != nil {
		log.Println("Euid parse err", err)
		return parsedProcInfo{}, true
	}

	r, err = regexp.Compile(`(?im)State:\s.*`)
	if err != nil {
		log.Println("State regex err", err)
		return parsedProcInfo{}, true
	}
	states := strings.Fields(r.FindString(procInfo))

	r, err = regexp.Compile(`(?im)Gid:\s.*`)
	if err != nil {
		log.Println("Gid regex err", err)
		return parsedProcInfo{}, true
	}

	gids := strings.Fields(r.FindString(procInfo))
	gid, err := strconv.Atoi(gids[1])
	if err != nil {
		log.Println("Gid parse err", err)
		return parsedProcInfo{}, true
	}

	return parsedProcInfo{
		Ruid:  ruid,
		Euid:  euid,
		State: states[1],
		Gid:   gid,
	}, false
}
