package runtime

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/structs"
	"syscall"
)

type CmdOutput struct {
	StdOut string
	StdErr string
}

func SafeCallLibrary(reqBody *structs.ExecuteCodeReqBody) (CmdOutput, error) {
	lang := packages[reqBody.Language]
	log.Println("lang", lang)

	execInfo, err := primeExecution(lang, reqBody.Code)
	if err != nil {
		log.Println(err)
		return CmdOutput{}, err
	}

	// TODO: uncomment this
	// tmpDir := os.TempDir() + execId
	tmpDir := "/tmp/" + execInfo.Id
	cmd := exec.Command("bash", "run_pkg.sh", lang.Cmd, tmpDir+"/"+execInfo.Id+"."+lang.Extension)
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd.Dir = "."
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid:         execInfo.Uid,
			Gid:         execInfo.Gid,
			Groups:      nil,
			NoSetGroups: true,
		},
		Setsid: true,
	}

	envData, err := os.ReadFile(config.LanguagePackagesDir + "/" + lang.SrcFolder + "/" + ".env")
	if err != nil {
		log.Println(err)
		return CmdOutput{}, err
	}
	log.Println("envData", envData)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, string(envData))

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
