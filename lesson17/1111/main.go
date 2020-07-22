package main

import (
	"time"

	"xiao.com/golang/lesson/logger"
)

var log logger.Logger

func initLogger(level int) {
	log = logger.NewConsole(level)
	log.DeBug("init logger success")
}

func run() {
	for {
		log.Error("server is running")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger(1)
	run()
	return
}
