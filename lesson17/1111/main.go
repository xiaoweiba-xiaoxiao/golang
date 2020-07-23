package main

import (
	"fmt"
	"time"

	"xiao.com/golang/lesson/logger"
)

// var log logger.Logger

func initLogger(name, logPath, logFile, level string) (err error) {

	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_file"] = logFile
	config["level"] = level
	err = logger.InitLogger(name, config)
	if err != nil {
		fmt.Printf("Error： %s", err)
		return
	}
	logger.Debug("init logger success")
	return
}

func run() {
	for {
		logger.Error("server is running\n")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "/home/golang/log", "test.log", "debug")
	run()
	return
}
