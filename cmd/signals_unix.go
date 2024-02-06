//go:build unix

package cmd

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
}
