package seabotserver

import (
	"encoding/json"
	"net"
)

type TcpBot struct {
	net.Conn
	ID      int64
	AuthKey string

	SendChannel chan []byte
	Done        chan bool
	Buffer      []byte
	BufferLen   int
}

func NewTcpBot(c net.Conn) *TcpBot {
	bot := &TcpBot{}
	bot.Conn = c
	bot.SendChannel = make(chan []byte)
	bot.Done = make(chan bool)
	bot.Buffer = make([]byte, 1024)
	return bot
}

func (s *TcpBot) Send(d interface{}) {
	ba, err := json.Marshal(d)
	if err != nil {
		return
	}
	s.SendChannel <- ba
}

type TcpServer struct {
	Listener     net.Listener
	ClientsCount int64
}
