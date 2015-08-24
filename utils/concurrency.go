package utils

import (
	"os"
	"runtime"
	"strconv"
)

// GetProcessCount - Get Process count defined by ENV or by NumCPU()
func GetProcessCount() uint {
	pc, err := strconv.Atoi(os.Getenv("PU_GO_MAX_PROCS"))

	if err != nil {
		pc = runtime.NumCPU()
	}

	return uint(pc)
}

// GetConcurrencyCount - Get Process count defined by ENV or by NumCPU()
func GetConcurrencyCount() uint {
	pc, err := strconv.Atoi(os.Getenv("PU_GO_MAX_CONCURRENCY"))

	if err != nil {
		pc = runtime.NumCPU()
	}

	return uint(pc)
}
