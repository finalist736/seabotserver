package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/finalist736/seabotserver/tcpserver"
)

func main() {

	rand.Seed(time.Now().UnixNano())

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
