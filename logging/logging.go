package logging

import (
	"fmt"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/powerunit-io/platform/utils"
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

	level := DefaultLoggingLevel

	if lvl := os.Getenv(levelEnv); lvl != "" {
		level = lvl
	}

	lvl, err := logrus.ParseLevel(level)

	if err != nil {
		return fmt.Errorf("Could not set logging level due to (err: %s)", err)
	}

	logrus.SetLevel(lvl)

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
func (l *Logger) Print(args ...interface{}) {
	logrus.Print(args...)
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
func New(conf map[string]interface{}) *Logger {
	logger := Logger{}

	forceColors := FormatterForceColors
	timestampFormat := FormatterTimestampFormat

	if utils.KeyInSlice("formatter_force_colors", conf) {
		forceColors = conf["formatter_force_colors"].(bool)
	}

	if utils.KeyInSlice("formatter_timestamp_format", conf) {
		timestampFormat = conf["formatter_timestamp_format"].(string)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     forceColors,
		TimestampFormat: timestampFormat,
	})

	if utils.KeyInSlice("output", conf) {
		logger.SetOutput(conf["output"].(io.Writer))
	}

	if utils.KeyInSlice("level", conf) {
		logger.SetLevel(conf["level"].(string))
	}

	return &logger
}
