package config

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
)

// to store the app env variables
var Host string
var Port string
var PostgresDsn string

var LanguagePackagesDir string

var AwsCognitoRegion string
var AwsCognitoPoolID string
var AwsCognitoClientID string
var AwsCognitoJwksUrl string
var AwsCognitoIssuer string
var AwsCognitoJwtCachedPublicKey *rsa.PublicKey

var RuntimeMinRunnerUid int
var RuntimeMaxRunnerUid int
var RuntimeMinRunnerGid int
var RuntimeMaxRunnerGid int
var RuntimeMaxProcessCount int
var RuntimeMaxOpenFiles int
var RuntimeMaxFileSize int
var RuntimeMaxMemoryLimit int

// to load the env variables from .env
func Init() {

	// appMode := flag.String("mode", "", "app env mode\nlocal for .env.local\nprod for .env.prod\ndev for .env.dev")
	// flag.Parse()

	// if contains([]string{"local", "dev", "prod"}, appMode) == false {
	// 	log.Fatalln("app env mode invalid")
	// }

	// err := godotenv.Load(fmt.Sprintf(".env.%s", *appMode))
	// if err != nil {
	// 	log.Fatal(fmt.Sprintf("Error loading .env.%s file", *appMode))
	// }

	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")

	PostgresDsn = os.Getenv("POSTGRES_DSN")

	LanguagePackagesDir = os.Getenv("PKG_DIR_PATH")

	RuntimeMinRunnerUid = parseInt(os.Getenv("RUNTIME_MIN_RUNNER_UID"))
	RuntimeMaxRunnerUid = parseInt(os.Getenv("RUNTIME_MAX_RUNNER_UID"))
	RuntimeMinRunnerGid = parseInt(os.Getenv("RUNTIME_MIN_RUNNER_GID"))
	RuntimeMaxRunnerGid = parseInt(os.Getenv("RUNTIME_MAX_RUNNER_GID"))
	RuntimeMaxProcessCount = parseInt(os.Getenv("RUNTIME_MAX_PROCESS_COUNT"))
	RuntimeMaxOpenFiles = parseInt(os.Getenv("RUNTIME_MAX_OPEN_FILES"))
	RuntimeMaxFileSize = parseInt(os.Getenv("RUNTIME_MAX_FILE_SIZE"))
	RuntimeMaxMemoryLimit = parseInt(os.Getenv("RUNTIME_MAX_MEMORY_LIMIT"))

	AwsCognitoRegion = os.Getenv("AWS_COGNITO_REGION")
	AwsCognitoPoolID = os.Getenv("AWS_COGNITO_POOL_ID")
	AwsCognitoClientID = os.Getenv("AWS_COGNITO_CLIENT_ID")
	AwsCognitoIssuer = fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", AwsCognitoRegion, AwsCognitoPoolID)
	AwsCognitoJwksUrl = fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", AwsCognitoRegion, AwsCognitoPoolID)

	var err error
	AwsCognitoJwtCachedPublicKey, err = cacheAWSCognitoJWK()
	if err != nil {
		log.Fatalln("Could not cache AWS Cognito jwks")
	} else {
		log.Println("Cached AWS Cognito jwks")
	}

}

func parseInt(str string) int {
	parsedInt, err := strconv.Atoi(str)
	if err != nil {
		log.Println("Error while parsing", str)
		log.Fatalln(err)
	}
	return parsedInt
}

func contains(s []string, e *string) bool {
	for _, a := range s {
		if a == *e {
			return true
		}
	}
	return false
}
