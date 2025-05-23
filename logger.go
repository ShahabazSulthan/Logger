package logger

import (
	"fmt"
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

// Debug logs a message at DEBUG level.
func Debug(args ...interface{}) {
	if currentLogLevel <= DEBUG && logEnabled {
		log.Println("[DEBUG]", fmt.Sprint(args...))
	}
}

// Info logs a message at INFO level.
func Info(args ...interface{}) {
	if currentLogLevel <= INFO && logEnabled {
		log.Println("[INFO]", fmt.Sprint(args...))
	}
}

// Warn logs a message at WARN level.
func Warn(args ...interface{}) {
	if currentLogLevel <= WARN && logEnabled {
		log.Println("[WARN]", fmt.Sprint(args...))
	}
}

// Error logs a message at ERROR level.
func Error(args ...interface{}) {
	if currentLogLevel <= ERROR && logEnabled {
		log.Println("[ERROR]", fmt.Sprint(args...))
	}
}
