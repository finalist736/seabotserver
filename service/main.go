package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/finalist736/seabotserver/storage/config"
	"github.com/finalist736/seabotserver/tcpserver"
	"github.com/pkg/profile"
)

var config_path *string = flag.String("config", "../config.json", "config file path")

func main() {

	flag.Parse()
	config.SetConfigFile(*config_path)
	conf := config.GetConfiguration()

	now := time.Now()
	if conf.Profiling {
		profilePath := fmt.Sprintf("./nav_profile_%d_%d_%d_%d_%d", // store profiles in current directory
			now.Day(), now.Month(),
			now.Hour(), now.Minute(), now.Second())
		p := profile.Start(profile.CPUProfile, profile.ProfilePath(profilePath), profile.NoShutdownHook)
		defer p.Stop()
	}

	rand.Seed(now.UnixNano())

	s := tcpserver.NewServer()
	err := s.StartListen(conf.Port)
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
