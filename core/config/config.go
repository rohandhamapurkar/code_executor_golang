package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// to store the app env variables
var AppConfig *appConfig

type appConfig struct {
	Port string
	Host string
}

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

	AppConfig = &appConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

}

func contains(s []string, e *string) bool {
	for _, a := range s {
		if a == *e {
			return true
		}
	}
	return false
}
