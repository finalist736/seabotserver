package seabotserver

import (
	"net"
)

type TcpBot struct {
	net.Conn
	ID      int64
	AuthKey string

	Send   chan []byte
	Done   chan bool
	Buffer []byte
}

func NewTcpBot(c net.Conn) *TcpBot {
	bot := &TcpBot{}
	bot.Conn = c
	bot.Send = make(chan []byte)
	bot.Done = make(chan bool)
	bot.Buffer = make([]byte, 1024)
	return bot
}

type TcpServer struct {
	Listener     net.Listener
	ClientsCount int64
}
