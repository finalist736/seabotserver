package seabotserver

import (
	"encoding/json"
	"net"
)

type TcpBot struct {
	net.Conn
	ID      int64
	AuthKey string

	Battle interface{}

	SendChannel chan []byte
	Done        chan bool
	Buffer      []byte
}

func NewTcpBot(c net.Conn) *TcpBot {
	bot := &TcpBot{}
	bot.Conn = c
	bot.SendChannel = make(chan []byte, 2)
	bot.Done = make(chan bool, 2)
	bot.Buffer = make([]byte, 0)
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
