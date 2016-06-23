package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signalChan, syscall.SIGHUP)

	go func() {
		for {
			sig := <-signalChan
			fmt.Println(sig)
			switch sig {
			case syscall.SIGTERM,
				syscall.SIGQUIT:
				done <- true
			case syscall.SIGHUP:
				fmt.Println("handling SIGHUP")
			}
		}
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
