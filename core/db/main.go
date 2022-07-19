package db

import "log"

func init() {
	defer log.Println("Initialized database connections")
	initPostgresDBConn()
}
