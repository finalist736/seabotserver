package tcpserver

import (
	"fmt"
	"net"

	"github.com/finalist736/seabotserver"
)

type Server struct {
	seabotserver.TcpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartListen(port string) error {
	var err error
	s.Listener, err = net.Listen("tcp", port)
	if err != nil {
		return err
	}
	fmt.Println("server listen started", s.Listener.Addr())
	go s.acceptConnections()
	return nil
}

func (s *Server) acceptConnections() {
	for {
		c, err := s.Listener.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("connected: %v\n", c.RemoteAddr())
		bot := seabotserver.NewTcpBot(c)
		go handle(bot)
		go sender(bot)
	}
}
