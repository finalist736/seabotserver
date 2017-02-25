package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/finalist736/seabotserver/tcpserver"
	"github.com/pkg/profile"
)

func main() {

	now := time.Now()
	//	cfg := profile.Config{
	//		CPUProfile: true,
	//		MemProfile: true,
	//		ProfilePath: fmt.Sprintf("./nav_profile_%d_%d_%d_%d_%d", // store profiles in current directory
	//			now.Day(), now.Month(),
	//			now.Hour(), now.Minute(), now.Second()),
	//		NoShutdownHook: true, // do not hook SIGINT
	//	}
	profilePath := fmt.Sprintf("./nav_profile_%d_%d_%d_%d_%d", // store profiles in current directory
		now.Day(), now.Month(),
		now.Hour(), now.Minute(), now.Second())
	p := profile.Start(profile.CPUProfile, profile.ProfilePath(profilePath), profile.NoShutdownHook)
	defer p.Stop()

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
