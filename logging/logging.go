package logging

import (
	"fmt"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/powerunit-io/platform/constants"
)

// Logger -
type Logger struct {
	logrus.Logger
}

// SetFormatter -
func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

// SetOutput -
func (l *Logger) SetOutput(output io.Writer) {
	logrus.SetOutput(output)
}

// SetLevel -
func (l *Logger) SetLevel(levelEnv string) error {

	level := constants.DEFAULT_LOGGING_LEVEL

	if lvl := os.Getenv(levelEnv); lvl != "" {
		level = lvl
	}

	if lvl, err := logrus.ParseLevel(level); err != nil {
		return fmt.Errorf("Could not set logging level due to (err: %s)", err)
	} else {
		logrus.SetLevel(lvl)
	}

	return nil
}

// Error -
func (l *Logger) Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Warning -
func (l *Logger) Warning(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

// Info -
func (l *Logger) Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Fatal -
func (l *Logger) Fatal(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Debug -
func (l *Logger) Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Print -
func (l *Logger) Print(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

// Panic -
func (l *Logger) Panic(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// GetContextLogger -
func (l *Logger) GetContextLogger(fields map[string]interface{}) *logrus.Entry {
	return l.WithFields(fields)
}

// New -
func New() *Logger {
	logger := Logger{}
	return &logger
}
