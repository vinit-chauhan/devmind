package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/vinit-chauhan/devmind/internal/constants"
)

var logger *log.Logger

func Init() {
	log_path := constants.LOG_PATH

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
