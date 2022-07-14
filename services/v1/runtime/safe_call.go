package runtime

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/structs"

	"golang.org/x/sys/unix"
)

type CmdOutput struct {
	StdOut string
	StdErr string
}

func readFromOutPipe(result *[]byte, ioPipe io.ReadCloser) {
	buf := make([]byte, 256)
	for {
		n, err := ioPipe.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		if len(*result) >= 1024 {
			msg := "\nOutput buffer limit execeeded..."
			*result = append(*result, []byte(msg)...)
			log.Println()
			break
		}
		*result = append(*result, buf[:n]...)
	}
}

func SafeCallLibrary(reqBody *structs.ExecuteCodeReqBody) (CmdOutput, error) {

	pkgInfo := packages[reqBody.Language]

	execInfo, err := primeExecution(pkgInfo, reqBody.Code)
	if err != nil {
		log.Println(err)
		return CmdOutput{}, err
	}

	defer cleanupExecution(execInfo)

	tmpDir := os.TempDir() + "/" + execInfo.Id

	cmdArgs := []string{
		"/usr/bin/nice",
		"/usr/bin/prlimit",
		fmt.Sprintf("--nproc=%d", config.RuntimeMaxProcessCount),
		fmt.Sprintf("--nofile=%d", config.RuntimeMaxOpenFiles),
		fmt.Sprintf("--fsize=%d", config.RuntimeMaxFileSize),
		// fmt.Sprintf("--as=%d", config.RuntimeMaxMemoryLimit),
		"/bin/bash",
		"./run_pkg.sh",
		pkgInfo.Cmd,
		tmpDir + "/" + execInfo.Id + "." + pkgInfo.Extension,
	}

	log.Println(cmdArgs)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = "."
	cmd.SysProcAttr = &unix.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: execInfo.Uid,
			Gid: execInfo.Gid,
		},
		Setpgid: true,
		Pgid:    0,
	}
	cmd.Env = append(os.Environ(), pkgInfo.EnvData)

	stdOut := []byte{}
	errOut := []byte{}

	stdOutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return CmdOutput{}, err
	}
	stdErrPipe, err := cmd.StderrPipe()
	if err != nil {
		return CmdOutput{}, err
	}

	defer stdOutPipe.Close()
	defer stdErrPipe.Close()

	go readFromOutPipe(&stdOut, stdOutPipe)
	go readFromOutPipe(&errOut, stdErrPipe)

	if err = cmd.Start(); err != nil {
		return CmdOutput{
			StdOut: string(stdOut),
			StdErr: string(errOut),
		}, err
	}

	log.Println("Executing with pid: ", cmd.Process.Pid)

	// 3 second timeout
	timer := time.AfterFunc(time.Second*3, func() {
		log.Println("Timeout exceeded, commencing Killings.")
		cmd.Process.Kill()
	})

	if err = cmd.Wait(); err != nil {
		timer.Stop()
		// if SIGKILL
		if err.Error() == "signal: killed" {
			return CmdOutput{}, errors.New("Execution Timeout exceeded")
		}
		// if other error
		return CmdOutput{
			StdOut: string(stdOut),
			StdErr: string(errOut),
		}, err
	} else {
		timer.Stop()
	}

	log.Println(len(stdOut))

	return CmdOutput{
		StdOut: string(stdOut),
		StdErr: string(errOut),
	}, nil
}
