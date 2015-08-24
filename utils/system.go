package utils

import "syscall"

// ShutdownSignal - Will send syscall sigterm.
func ShutdownSignal() {
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}
