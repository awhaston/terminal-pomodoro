package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleExit() {
	fmt.Print("\033[?25h")
	os.Exit(1)
}

func initCleanup() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
	}()
}

func newLoopPrint() {
	fmt.Print("\033[2J\033[H\033[?25l")
}
