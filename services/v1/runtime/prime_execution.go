package runtime

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"rohandhamapurkar/code-executor/core/constants"

	"github.com/google/uuid"
)

type executionInfo struct {
	ID  string
	Uid uint32
	Gid uint32
}

func primeExecution(pkg pkgInfo, code string) (executionInfo, error) {
	execID := uuid.New().String()

	uid := minRunnerUid + int(runnerIncrementUid)
	gid := minRunnerGid + int(runnerIncrementGid)

	tmpDir := os.TempDir() + "/" + execID

	if err := os.Mkdir(tmpDir, 0700); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.MKDIR_FAILED + ":" + execID)
	}

	if err := chownR(tmpDir, int(uid), int(gid)); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_CHOWN_DIR + ":" + execID)
	}

	if err := os.WriteFile(tmpDir+"/"+execID+"."+pkg.Extension, []byte(code), 0700); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_WRITE_FILE + ":" + execID)
	}

	if err := chownR(tmpDir, int(uid), int(gid)); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_CHOWN_DIR + ":" + execID)
	}

	runnerIncrementMutex.Lock()
	// increment uid and gid
	runnerIncrementUid++
	runnerIncrementGid++
	runnerIncrementUid %= uint32(maxRunnerUid - minRunnerUid + 1)
	runnerIncrementGid %= uint32(maxRunnerGid - minRunnerGid + 1)
	runnerIncrementMutex.Unlock()

	return executionInfo{ID: execID, Uid: uint32(uid), Gid: uint32(gid)}, nil
}

func chownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, _ os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}
