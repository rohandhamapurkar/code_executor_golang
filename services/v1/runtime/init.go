package runtime

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"rohandhamapurkar/code-executor/core/config"

	"golang.org/x/sys/unix"
	"gopkg.in/yaml.v2"
)

type pkgInfo struct {
	Language  string
	Version   string
	Extension string
	EnvData   string
	SrcFolder string `yaml:"src_folder"`
	Cmd       string
}

var packages map[string]pkgInfo

var runnerIncrementUid uint32
var runnerIncrementGid uint32

var minRunnerUid int
var minRunnerGid int
var maxRunnerUid int
var maxRunnerGid int

func Init() {
	buf, err := ioutil.ReadFile("./languages.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(buf, &packages)
	if err != nil {
		log.Fatalln(err)
	}

	for key, element := range packages {
		pkgEnv, err := os.ReadFile(config.LanguagePackagesDir + "/" + element.SrcFolder + "/" + ".env")
		if err != nil {
			log.Println("Error while loading pkg env for", element.Language)
			log.Fatalln(err)
		}
		if entry, ok := packages[key]; ok {
			entry.EnvData = string(pkgEnv)
			packages[key] = entry
		} else {
			log.Fatalln(errors.New("Could not assign env: " + element.Language))
		}
	}

	err = unix.Prctl(unix.PR_SET_CHILD_SUBREAPER, uintptr(1), 0, 0, 0)
	if err != nil {
		log.Println("Prctl err", err)
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, unix.SIGCHLD)

	go cleanupProcesses(sigs)

	// intialize runner uid and gid
	runnerIncrementUid = 0
	runnerIncrementGid = 0

	minRunnerUid = config.RuntimeMinRunnerUid
	maxRunnerUid = config.RuntimeMaxRunnerUid
	minRunnerGid = config.RuntimeMinRunnerGid
	maxRunnerGid = config.RuntimeMaxRunnerGid

}
