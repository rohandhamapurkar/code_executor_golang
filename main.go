package main

import (
	"fmt"
	"net/http"
	"rohandhamapurkar/code-executor/core/config"
	v1 "rohandhamapurkar/code-executor/routers/v1"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	config.Init()
	router = gin.Default()
	v1.SetV1Routes(router)
}

func main() {
	fmt.Println("Server Running on: ", config.AppConfig.Host+":"+config.AppConfig.Port)
	http.ListenAndServe(config.AppConfig.Host+":"+config.AppConfig.Port, router)
}
