package runtime

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"rohandhamapurkar/code-executor/core/config"
	"sync"

	"golang.org/x/sys/unix"
	"gopkg.in/yaml.v2"
)

type pkgInfo struct {
	Version   string
	Extension string
	EnvData   string
	SrcFolder string `yaml:"src_folder"`
	Cmd       string
}

type language struct {
	Language  string `json:"language"`
	Version   string `json:"version"`
	Extension string `json:"extension"`
}

var Packages map[string]pkgInfo

var PackagesJSON *[]language

var runnerIncrementMutex sync.Mutex
var runnerIncrementUid uint32
var runnerIncrementGid uint32

var minRunnerUid int
var minRunnerGid int
var maxRunnerUid int
var maxRunnerGid int

func init() {
	defer log.Println("Initialized runtime service")
	buf, err := ioutil.ReadFile("./languages.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(buf, &Packages)
	if err != nil {
		log.Fatalln(err)
	}

	PackagesJSON = &[]language{}

	for key, element := range Packages {

		pkgEnv, err := os.ReadFile(config.LanguagePackagesDir + "/" + element.SrcFolder + "/" + ".env")

		if err != nil {
			log.Println("Error while loading pkg env for", key)
			log.Fatalln(err)
		}
		*PackagesJSON = append(*PackagesJSON, language{Language: key, Version: element.Version, Extension: element.Extension})

		if entry, ok := Packages[key]; ok {
			entry.EnvData = string(pkgEnv)
			Packages[key] = entry
		} else {
			log.Fatalln(errors.New("Could not assign env: " + key))
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
