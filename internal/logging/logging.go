package logging

import (
	log "github.com/hhkbp2/go-logging"
	"os"
)

var logger log.Logger = nil

// NewLogger zwraca nowy logger
func NewLogger() log.Logger {
	if logger != nil {
		return logger
	}
	filepath := "dataCollector.log"
	fileMode := os.O_APPEND

	fileHandler, err := log.NewFileHandler(filepath, fileMode, 0)
	if err != nil {
		panic(err.Error())
	}

	handler := log.NewStdoutHandler()
	format := "%(asctime)s %(levelname)s (%(filename)s:%(lineno)d) " +
		"%(name)s %(message)s"
	dateFormat := "%Y-%m-%d %H:%M:%S.%3n"
	formatter := log.NewStandardFormatter(format, dateFormat)
	handler.SetFormatter(formatter)
	fileHandler.SetFormatter(formatter)

	logger = log.GetLogger("dataCollector")
	logger.SetLevel(log.LevelDebug)
	logger.AddHandler(handler)
	logger.AddHandler(fileHandler)

	logger.Debug("Logger skonfigurowany")

	return logger
}

// Shutdown zamyka log
func Shutdown() {
	log.Shutdown()
}
