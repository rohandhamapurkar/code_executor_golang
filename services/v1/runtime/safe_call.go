package runtime

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"rohandhamapurkar/code-executor/core/structs"
)

type CmdOutput struct {
	StdOut string
	StdErr string
}

func SafeCallLibrary(reqBody *structs.ExecuteCodeReqBody) (CmdOutput, error) {
	pkgInfo := packages[reqBody.Language]

	execInfo, err := primeExecution(pkgInfo, reqBody.Code)
	if err != nil {
		log.Println(err)
		return CmdOutput{}, err
	}
	log.Println(execInfo)

	// tmpDir := os.TempDir() + "/" + execInfo.Id
	tmpDir := os.TempDir() + "/" + execInfo.Id

	cmd := exec.Command("bash", "run_pkg.sh", pkgInfo.Cmd, tmpDir+"/"+execInfo.Id+"."+pkgInfo.Extension)
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd.Dir = "."
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: execInfo.Uid,
			Gid: execInfo.Gid,
		},
		Setsid:     true,
		Foreground: false,
	}

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, pkgInfo.EnvData)

	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err = cmd.Run()

	if err != nil {
		log.Println(err)
		log.Println(errOut.String())
	}

	fmt.Println(out.String())

	return CmdOutput{
		StdOut: out.String(),
		StdErr: errOut.String(),
	}, nil
}
