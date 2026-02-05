package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// NewLogger creates a new application logger
func NewLogger(verbose bool) *AppLogger {
	return &AppLogger{
		logger:  log.New(os.Stderr, "", 0),
		verbose: verbose,
	}
}

// formatMessage formats log messages with timestamp and level
func (l *AppLogger) formatMessage(level, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
}

// Debug logs debug messages (only if verbose is enabled)
func (l *AppLogger) Debug(format string, args ...interface{}) {
	if l.verbose {
		message := fmt.Sprintf(format, args...)
		l.logger.Println(l.formatMessage(LOG_LEVEL_DEBUG, message))
	}
}

// Info logs informational messages
func (l *AppLogger) Info(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Println(l.formatMessage(LOG_LEVEL_INFO, message))
}

// Warn logs warning messages
func (l *AppLogger) Warn(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Println(l.formatMessage(LOG_LEVEL_WARN, message))
}

// Error logs error messages
func (l *AppLogger) Error(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Println(l.formatMessage(LOG_LEVEL_ERROR, message))
}

// Fatal logs a fatal error and exits
func (l *AppLogger) Fatal(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Println(l.formatMessage(LOG_LEVEL_FATAL, message))
	os.Exit(1)
}
