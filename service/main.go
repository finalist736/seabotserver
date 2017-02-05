package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/finalist736/seabotserver/tcpserver"
)

func main() {
	s := tcpserver.NewServer()
	err := s.StartListen(":11000")
	if err != nil {
		panic(err)
	}

	signalToClose := make(chan os.Signal, 1)
	signal.Notify(signalToClose,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGKILL)
	for {
		<-signalToClose
		fmt.Println("signal received, closing...")
		return
	}
}
