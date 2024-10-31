package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
	"time"
)

var logObj *logrus.Logger
var once sync.Once

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logObj = createLogger()
	})

	return logObj
}

func createLogger() *logrus.Logger {
	fmt.Println("Creating logger instance")

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	if err := os.MkdirAll("./logs", 0755); err != nil {
		fmt.Println("Failed to create log directory: ", err)
	}

	timeStr := time.Now().Format("2006-01-02")

	file, err := os.OpenFile("./logs/"+timeStr+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Out = os.Stderr
		logger.WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Warn("Failed to log to file, using default stderr")
	} else {
		logger.Out = io.MultiWriter(os.Stderr, file)
	}

	return logger
}
