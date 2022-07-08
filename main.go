package main

import (
	"log"
	"net/http"
	appConfig "rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/db"
	v1 "rohandhamapurkar/code-executor/routers/v1"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	appConfig.Init()
	db.InitPostgresDBConn()
	router = gin.Default()
	v1.SetV1Routes(router)
}

func main() {
	log.Println("Server Running on: ", appConfig.Host+":"+appConfig.Port)
	http.ListenAndServe(appConfig.Host+":"+appConfig.Port, router)
}
