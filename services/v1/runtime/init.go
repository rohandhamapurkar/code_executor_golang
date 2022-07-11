package runtime

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type pkgInfo struct {
	Language  string
	Version   string
	Extension string
	SrcFolder string `yaml:"src_folder"`
	Cmd       string
}

var packages map[string]pkgInfo

var runnerUid uint32
var runnerGid uint32

func Init() {
	buf, err := ioutil.ReadFile("./languages.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(buf, &packages)
	if err != nil {
		log.Fatalln(err)
	}

	// intialize runner uid and gid
	runnerUid = 0
	runnerGid = 0

}
