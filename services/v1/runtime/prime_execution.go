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

func primeExecution(lang pkgInfo, code string) (executionInfo, error) {
	execId := uuid.New().String()

	uid := runnerUid
	gid := runnerGid

	// TODO: uncomment this
	// tmpDir := os.TempDir() + execId
	tmpDir := "/tmp/" + execId

	if err := os.Mkdir(tmpDir, 0644); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.MKDIR_FAILED + ":" + execId)
	}

	if err := chownR(tmpDir, int(uid), int(gid)); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_CHOWN_DIR + ":" + execId)
	}

	if err := os.WriteFile(tmpDir+"/"+execId+"."+lang.Extension, []byte(code), 0444); err != nil {
		log.Println(err)
		return executionInfo{}, errors.New(constants.CANNOT_WRITE_FILE + ":" + execId)
	}

	// increment uid and gid
	runnerUid++
	runnerGid++
	runnerUid %= 1500 - 1000 + 1
	runnerGid %= 1500 - 1000 + 1

	return executionInfo{Id: execId, Uid: uid, Gid: gid}, nil
}

func chownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}
