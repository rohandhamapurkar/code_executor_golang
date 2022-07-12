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
	Id  string
	Uid uint32
	Gid uint32
}

func primeExecution(pkg pkgInfo, code string) (executionInfo, error) {
	execId := uuid.New().String()

	uid := minRunnerUid + int(runnerIncrementUid)
	gid := minRunnerGid + int(runnerIncrementGid)

	tmpDir := os.TempDir() + "/" + execId

	if err := os.Mkdir(tmpDir, 0700); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.MKDIR_FAILED + ":" + execId)
	}

	log.Println(uid, gid)
	if err := chownR(tmpDir, int(uid), int(gid)); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_CHOWN_DIR + ":" + execId)
	}

	if err := os.WriteFile(tmpDir+"/"+execId+"."+pkg.Extension, []byte(code), 0700); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_WRITE_FILE + ":" + execId)
	}

	if err := chownR(tmpDir, int(uid), int(gid)); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_CHOWN_DIR + ":" + execId)
	}

	// increment uid and gid
	runnerIncrementUid++
	runnerIncrementGid++
	runnerIncrementUid %= uint32(maxRunnerUid - minRunnerUid + 1)
	runnerIncrementGid %= uint32(maxRunnerGid - minRunnerGid + 1)

	return executionInfo{Id: execId, Uid: uint32(uid), Gid: uint32(gid)}, nil
}

func chownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}
