package runtime

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"rohandhamapurkar/code-executor/core/config"

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

var minRunnerUid = 1001
var minRunnerGid = 1001
var maxRunnerUid = 1500
var maxRunnerGid = 1500

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

	// intialize runner uid and gid
	runnerIncrementUid = 0
	runnerIncrementGid = 0

}
