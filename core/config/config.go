package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// to store the app env variables
var Host string
var Port string
var AwsCognitoRegion string
var AwsCognitoClientId string
var AwsAccessKeyId string
var AwsSecretAccessKey string
var PostgresDsn string

// to load the env variables from .env
func Init() {

	appMode := flag.String("mode", "", "app env mode\nlocal for .env.local\nprod for .env.prod\ndev for .env.dev")
	flag.Parse()

	if contains([]string{"local", "dev", "prod"}, appMode) == false {
		log.Fatalln("app env mode invalid")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", *appMode))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env.%s file", *appMode))
	}

	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")

	AwsCognitoRegion = os.Getenv("AWS_COGNITO_REGION")
	AwsCognitoClientId = os.Getenv("AWS_COGNITO_CLIENT_ID")
	AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	PostgresDsn = os.Getenv("POSTGRES_DSN")

}

func contains(s []string, e *string) bool {
	for _, a := range s {
		if a == *e {
			return true
		}
	}
	return false
}
