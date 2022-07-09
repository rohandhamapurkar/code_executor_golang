// exec_linux.go
//go:build linux
// +build linux

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("bash", "run.sh", "java", "--version")

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		log.Fatal(errOut.String())
	}

	fmt.Printf(out.String())

}
