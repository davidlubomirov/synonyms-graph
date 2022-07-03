package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func signalHandler(sig os.Signal) {
	fmt.Printf("signal received: %d", sig)
}

func handleSystemSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				signalHandler(sig)
				os.Exit(int(syscall.SIGINT))
			case syscall.SIGTERM:
				signalHandler(sig)
				os.Exit(int(syscall.SIGTERM))
			default:
				log.Fatalf("not handled os signal. Signal received: %d", sig)
			}
		}
	}()

}
