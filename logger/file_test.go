package logger

import (
	"testing"
)

func TestFileLog(t *testing.T) {
	var m map[string]string = map[string]string{
		"log_path": "/home/golang/log",
		"log_file": "test.log",
		"level":    "0",
	}
	log := NewFileLog(m)
	log.DeBug("ERROR: code %d", 32222)
	log.Error("ERROR: %s", "test error")
	log.Close()
	// log := NewConsole(5)
	// log.DeBug("server is running DeBug")
	// log.Trace("server is running Trace")
	// log.Error("server is running Error")
}

func TestConsole(t *testing.T) {
	log := NewConsole(LeDeBug)
	log.DeBug("server is runing")
	log.Error("server is runing Error")
}
