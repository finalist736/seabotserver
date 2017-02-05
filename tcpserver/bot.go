package tcpserver

import (
	"fmt"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/router"
)

func sender(p *seabotserver.TcpBot) {
	defer p.Close()
	defer func() { fmt.Printf("sender close: %v\n", p.RemoteAddr()) }()
	var err error
	var n int
	for {
		select {
		case <-p.Done:
			return
		case msg := <-p.Send:
			var total int
			needToSend := len(msg)
			for {
				n, err = p.Write(msg[total:])
				total += n
				if err != nil {
					fmt.Printf("read error: %s\n", err.Error())
					// TODO
					// logging
					p.Done <- true
					return
				}
				if total == needToSend {
					break
				}
			}
		}
	}
}

func handle(p *seabotserver.TcpBot) {
	defer func() { fmt.Printf("handler close: %v\n", p.RemoteAddr()) }()
	for {

		select {
		case <-p.Done:
			return
		default:
			n, err := p.Read(p.Buffer)
			if err != nil {
				fmt.Printf("read error: %s\n", err.Error())
				p.Done <- true
				return
			}
			if n == 0 {
				//fmt.Printf("zero read\n")
				continue
			}
			fmt.Printf("received: %s\n", p.Buffer)
			router.Dispatch(p)
		}
	}
}
