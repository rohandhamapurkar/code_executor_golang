// exec_linux.go
//go:build linux
// +build linux

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("whoami")

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Dir = "/go/exec/"
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: 1001,
			Gid: 1001,
		},
		Pgid: 0,
	}
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PATH=/tmp/pkg/node_v14.20.0/bin:/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()

	if err != nil {
		log.Println(err)
		log.Fatalln(errOut.String())
	}

	fmt.Println(out.String())

}
