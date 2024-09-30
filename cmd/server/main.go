package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigSystem := make(chan os.Signal, 1)

	signal.Notify(sigSystem, syscall.SIGINT, syscall.SIGTERM)

	println("Hello")

	// your code ...

	<-sigSystem
}
