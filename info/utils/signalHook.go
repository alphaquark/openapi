package myutil

import (
	"os"
	"os/signal"
	"syscall"
)

// CleanupSignalHook return channel:
// Notify os.Interrupt | syscall.SIGTERM
func CleanupSignalHook() chan os.Signal {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	return c
}
