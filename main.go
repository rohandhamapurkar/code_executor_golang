package main

import (
	"log"
	"net/http"
	appConfig "rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/db"
	httpMiddleware "rohandhamapurkar/code-executor/middlewares/http"
	v1 "rohandhamapurkar/code-executor/routers/v1"
	runtimeService "rohandhamapurkar/code-executor/services/v1/runtime"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	appConfig.Init()
	db.InitPostgresDBConn()
	runtimeService.Init()
	go httpMiddleware.CleanupVisitors()
	router = gin.Default()
	v1.SetV1Routes(router)
}

func main() {
	log.Println("Server Running on: ", appConfig.Host+":"+appConfig.Port)
	http.ListenAndServe(appConfig.Host+":"+appConfig.Port, router)
}
