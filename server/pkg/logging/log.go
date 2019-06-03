package logging

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"cloud.google.com/go/logging"
)

type Level int

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
)

var (
	logger *log.Logger

	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup(projectId, logName string) {

	if len(projectId) != 0 {
		ctx := context.Background()

		client, err := logging.NewClient(ctx, projectId)
		if err != nil {
			log.Fatalf("Failed to create gcp stacklog client: %v", err)
		}
		defer client.Close()

		logger = client.Logger(logName).StandardLogger(logging.Info)
	}
}

func Debug(v ...interface{}) {
	logPrint(DEBUG, v)
}

func Info(v ...interface{}) {
	logPrint(INFO, v)
}

func Warn(v ...interface{}) {
	logPrint(WARNING, v)
}

func Error(v ...interface{}) {
	logPrint(ERROR, v)
}

func Fatal(v ...interface{}) {
	logPrint(FATAL, v)
}

func setPrefix(level Level) (logPrefix string) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	return
}

func logPrint(level Level, v ...interface{}) {
	prefixStr := setPrefix(level)

	if logger != nil {
		logger.SetPrefix(prefixStr)
		logger.Printf("%v", v)
	} else {
		fmt.Printf("%s %v", prefixStr, v)
	}

}
