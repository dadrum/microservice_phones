package file_logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// path and file name constants
const (
	logDir          = "logs"
	logFile         = "logrus.log"
	defaultLogLevel = logrus.WarnLevel
)

// struct for dependency injection
type FileLogger struct {
	*logrus.Entry
}

// --------------------------------------------------------------------------------------
// initialize custom logger with selected log level
// levelName - codename of logging level
func Init(levelName *string) (*FileLogger, error) {
	// Creates a new logger
	l := logrus.NewEntry(logrus.New())
	logger := FileLogger{l}

	// Log as text ASCII formatter.
	logger.Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// parse log level
	logLevel := parseLoglevel(levelName)
	logger.Logger.SetLevel(logLevel)

	// Create dirrectory to log
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Println("Can't create log dir. File loggin is not initialized")
		logger.Println("Logger was initialized with `" + logLevel.String() + "` level")
		return &logger, err
	}

	// Create log file
	file, err := os.OpenFile(logDir+"/"+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// file not available
		fmt.Println("Access denied to " + logDir + "/" + logFile + ", using default stderr")
		logger.Logger.SetOutput(os.Stdout)
		logger.Println("Logger was initialized with `" + logLevel.String() + "` level")
		return &logger, err
	}

	// sets the logger output
	logger.Logger.SetOutput(file)

	logger.Println("Logger was initialized with `" + logLevel.String() + "` level")
	return &logger, nil

}

// --------------------------------------------------------------------------------------
// determinate initialized log level
// return const `defaultLogLevel` on error or unknown `src` value
func parseLoglevel(src *string) logrus.Level {
	if src == nil {
		return defaultLogLevel
	}
	level, err := logrus.ParseLevel(*src)
	if err != nil {
		return defaultLogLevel
	} else {
		return level
	}
}
