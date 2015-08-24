package utils

import (
	"os"
	"runtime"
	"strconv"
)

// GetProcessCount - Get Process count defined by ENV or by NumCPU()
func GetProcessCount(env string) int {

	envName := "PU_GO_MAX_PROCS"

	if env != "" {
		envName = env
	}

	pc, err := strconv.Atoi(os.Getenv(envName))

	if err != nil {
		pc = runtime.NumCPU()
	}

	return int(pc)
}

// GetConcurrencyCount - Get Process count defined by ENV or by NumCPU()
func GetConcurrencyCount(env string) int {

	envName := "PU_GO_MAX_CONCURRENCY"

	if env != "" {
		envName = env
	}

	pc, err := strconv.Atoi(os.Getenv(envName))

	if err != nil {
		pc = runtime.NumCPU()
	}

	return int(pc)
}
