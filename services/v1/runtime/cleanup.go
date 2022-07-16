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
	err := os.RemoveAll(os.TempDir() + "/" + execInfo.ID)
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

			processIDs, err := ioutil.ReadDir("/proc")
			if err != nil {
				log.Println(err)
				return
			}

			for _, proc := range processIDs {
				procName := proc.Name()
				procID, err := strconv.Atoi(procName)
				if err != nil {
					continue
				}
				fmt.Println("procID", procID)

				data, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", procID))
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
						waitForPids = append(waitForPids, procID)
					} else {
						process, err := os.FindProcess(procID)
						if err != nil {
							log.Println("Couldn't find process id:", procID, err)
							continue
						}
						process.Kill()
					}
					// log.Println("pid:", procID)
					// log.Println("uid:", info.Ruid, info.Euid)
					// log.Println("gid:", info.Gid)
					// log.Println("state:", info.State)
				} else {
					log.Println("pid:", procID, "Not my process")
					continue
				}
			}

			for _, procID := range waitForPids {
				log.Println("waiting proc id", procID)
				var status unix.WaitStatus
				pid, err := unix.Wait4(procID, &status, unix.WNOHANG, nil)
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
