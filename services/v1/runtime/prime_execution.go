package runtime

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"rohandhamapurkar/code-executor/core/constants"

	"github.com/google/uuid"
)

func chownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}

func primeExecution(lang pkgInfo, code string) (string, error) {
	execId := uuid.New().String()

	// TODO: uncomment this
	// tmpDir := os.TempDir() + execId
	tmpDir := "/tmp/" + execId

	if err := chownR(tmpDir, int(runnerUid), int(runnerGid)); err != nil {
		return "", errors.New(constants.CANNOT_CHOWN_DIR + ":" + execId)
	}

	if err := os.WriteFile(tmpDir+"/"+execId+"."+lang.Extension, []byte(code), 0444); err != nil {
		return "", errors.New(constants.CANNOT_WRITE_FILE + ":" + execId)
	}

	// increment uid and gid
	runnerUid++
	runnerGid++
	runnerUid %= 1500 - 1000 + 1
	runnerGid %= 1500 - 1000 + 1

	log.Println(runnerUid, runnerGid)

	log.Println("execId", execId)

	return execId, nil
}
