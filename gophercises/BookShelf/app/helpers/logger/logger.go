package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

//Logger - Global instance of logger
var logger *logrus.Logger

//InitLogging - Initialize logging parameters
func InitLogging() {

	logger = logrus.New()

	logFile, err := os.OpenFile("./logs/Server.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	logger.Out = logFile

}

//LogError - Method to log errors
func LogError(args ...interface{}) {
	logger.Error(args)
}

//LogInfo - Method to log errors
func LogInfo(args ...interface{}) {
	logger.Info(args)
}

//LogWran - Method to log errors
func LogWran(args ...interface{}) {
	logger.Warn(args)
}

//LogDebug - Method to log errors
func LogDebug(args ...interface{}) {
	logger.Debug(args)
}
