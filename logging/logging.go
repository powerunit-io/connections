package logging

import (
	"github.com/Sirupsen/logrus"
)

// Logger -
type Logger struct {
	logrus.Logger
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
