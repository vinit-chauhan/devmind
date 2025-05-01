package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var logger *log.Logger

func init() {
	log_path := "devmind.log"

	if runtime.GOOS == "windows" {
		log_path = os.Getenv("APPDATA") + "\\devmind.log"
	} else if runtime.GOOS == "linux" {
		log_path = "/var/logs/devmind.log"
	} else if runtime.GOOS == "darwin" {
		log_path = "/Library/Logs/devmind.log"
	}

	log_file, err := os.OpenFile(log_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening log file:", err)
		return
	}

	logger = log.Default()
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetPrefix("devmind: ")
	logger.SetOutput(log_file)
}

func Debug(msg string) {
	logger.Println("DEBUG: " + msg)
}

func Info(msg string) {
	logger.Println("INFO: " + msg)
}

func Warn(msg string) {
	logger.Println("WARN: " + msg)
}

func Error(msg string) {
	logger.Println("ERROR: " + msg)
}
