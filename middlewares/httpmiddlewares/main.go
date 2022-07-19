package httpmiddlewares

import "log"

func init() {
	defer log.Println("Initialized http middleware cleanups")
	go cleanupVisitors()
}
