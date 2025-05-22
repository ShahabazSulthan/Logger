package logger

import (
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	NONE
)

var (
	currentLogLevel LogLevel = INFO
	logEnabled      bool     = true
)

func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

func SetLogLevelFromString(levelStr string) {
	switch strings.ToLower(strings.TrimSpace(levelStr)) {
	case "debug":
		SetLogLevel(DEBUG)
	case "info":
		SetLogLevel(INFO)
	case "warn":
		SetLogLevel(WARN)
	case "error":
		SetLogLevel(ERROR)
	case "none":
		SetLogLevel(NONE)
	default:
		SetLogLevel(INFO)
	}
}

func SetLogEnabled(enabled bool) {
	logEnabled = enabled
}

func InitializeLoggerLevel() {
	logLevel := os.Getenv("ZTNA_LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	SetLogLevelFromString(logLevel)

	logEnabledStr := os.Getenv("LOG_ENABLED")
	if strings.ToLower(logEnabledStr) == "false" {
		SetLogEnabled(false)
	}
}

func Debug(format string, args ...interface{}) {
	if currentLogLevel <= DEBUG && logEnabled {
		log.Printf("[DEBUG] "+format, args...)
	}
}

func Info(format string, args ...interface{}) {
	if currentLogLevel <= INFO && logEnabled {
		log.Printf("[INFO] "+format, args...)
	}
}

func Warn(format string, args ...interface{}) {
	if currentLogLevel <= WARN && logEnabled {
		log.Printf("[WARN] "+format, args...)
	}
}

func Error(format string, args ...interface{}) {
	if currentLogLevel <= ERROR && logEnabled {
		log.Printf("[ERROR] "+format, args...)
	}
}
