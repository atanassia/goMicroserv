// Golang program to handle different types of signals
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	SigChan := make(chan os.Signal, 1)

	signal.Notify(SigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-SigChan

	switch sig {
	case syscall.SIGHUP:
		fmt.Println("\nSIGHUP signal generated")

	case syscall.SIGINT:
		fmt.Println("\nSIGINT signal generated")

	case syscall.SIGTERM:
		fmt.Println("\nSIGTERM signal generated")

	case syscall.SIGQUIT:
		fmt.Println("\nSIGQUIT signal generated")

	default:
		fmt.Println("\nUNKNOWN signal generated")
	}
}