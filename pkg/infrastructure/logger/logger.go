package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

// Ultra basic implementation
type BasicLogger struct {
	logger *log.Logger
}

func New(prefix string) *BasicLogger {

	p := fmt.Sprintf("%s: ", prefix)
	logger := log.New(os.Stderr, p, log.LstdFlags)

	return &BasicLogger{
		logger: logger,
	}
}

func (l *BasicLogger) Debug(message string, args ...interface{}) {
	p := fmt.Sprintf("DEBUG: %s", message)
	l.logger.Printf(p, args)
}

func (l *BasicLogger) Info(message string, args ...interface{}) {
	p := fmt.Sprintf("INFO: %s", message)
	l.logger.Printf(p, args)
}

func (l *BasicLogger) Warn(message string, args ...interface{}) {
	p := fmt.Sprintf("WARNING: %s", message)
	l.logger.Printf(p, args)
}

func (l *BasicLogger) Error(message string, args ...interface{}) {
	p := fmt.Sprintf("ERROR: %s", message)
	l.logger.Printf(p, args)

}

func (l *BasicLogger) Fatal(message string, args ...interface{}) {
	p := fmt.Sprintf("FATAL: %s", message)
	l.logger.Fatalf(p, args)

}
