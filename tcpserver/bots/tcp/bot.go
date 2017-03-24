package tcp

import (
	"encoding/json"
	"io"
	"net"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/gameplay/battle"
	"github.com/finalist736/seabotserver/gameplay/queue"
	"github.com/finalist736/seabotserver/service/router"
)

type TcpBot struct {
	net.Conn

	dbbot *seabotserver.DBBot
	btl   interface{}

	sendChannel chan []byte
	Done        chan bool
	bfr         []byte
}

func NewBot(c net.Conn) seabotserver.BotService {
	bot := &TcpBot{}
	bot.dbbot = &seabotserver.DBBot{}
	bot.Conn = c
	bot.sendChannel = make(chan []byte, 2)
	bot.Done = make(chan bool, 2)
	bot.bfr = make([]byte, 0)
	return bot
}

func (s *TcpBot) Battle() interface{} {
	return s.btl
}

func (s *TcpBot) SetBattle(b interface{}) {
	s.btl = b
}

func (s *TcpBot) DBBot() *seabotserver.DBBot {
	return s.dbbot
}

func (s *TcpBot) SetDBBot(b *seabotserver.DBBot) {
	s.dbbot = b
}

func (s *TcpBot) Send(d interface{}) {
	ba, err := json.Marshal(d)
	if err != nil {
		return
	}
	s.sendChannel <- ba
}

func (s *TcpBot) SendError(err string) {
	s.Send(&seabotserver.ToBot{Error: &seabotserver.TBError{Error: err}})
}

func (s *TcpBot) Disconnect() {
	s.Done <- true
}

func (p *TcpBot) Sender() {
	defer p.Conn.Close()
	//defer func() { fmt.Printf("sender close: %v\n", p.RemoteAddr()) }()
	var err error
	var n int

	var size int
	var send_buff []byte

	for {
		select {
		case <-p.Done:
			return
		case msg := <-p.sendChannel:

			size = len(msg)
			send_buff = make([]byte, 0)
			send_buff = append(send_buff, byte(size>>24))
			send_buff = append(send_buff, byte(size>>16))
			send_buff = append(send_buff, byte(size>>8))
			send_buff = append(send_buff, byte(size))
			send_buff = append(send_buff, msg...)

			//conn.Client.Write(send_buff)

			//fmt.Printf("send to player:\n\t%s\n\n", msg)
			var total int
			needToSend := len(send_buff)
			for {
				n, err = p.Write(send_buff[total:])
				total += n
				if err != nil {
					//fmt.Printf("read error: %s\n", err.Error())
					// TODO
					// logging
					p.Disconnect()
					return
				}
				if total == needToSend {
					break
				}
			}
		}
	}
}

func (p *TcpBot) Handler() {
	//defer func() { fmt.Printf("handler close: %v\n", p.RemoteAddr()) }()
	defer func() { queue.Exit(p) }()
	defer func() {
		if p.Battle() == nil {
			return
		}
		btl := p.Battle().(*battle.Battle)
		if btl == nil {
			return
		}
		btl.Exit(p)
	}()

	tmp_buffer := make([]byte, 4)
	var numbytes, tmp_numbytes, size, atempts int
	var err error
	numbytes = 0
	fbot := &seabotserver.FromBot{}
	for {
		select {
		case <-p.Done:
			return
		default:
			// set read deadline for pvp battle or pve battle
			for {
				tmp_numbytes, err = p.Read(tmp_buffer)
				if err != nil {
					if err != io.EOF {
						//fmt.Printf("4 bytes read error: %s\n", err.Error())
					}
					p.Disconnect()
					return
				}
				numbytes += tmp_numbytes
				if numbytes < 4 {
					continue
				}
				break
			}

			size = 0
			for i := 0; i < 4; i++ {
				size = size*256 + int(tmp_buffer[i])
			}
			if size == 0 {
				atempts++
				continue
			}
			if size > 9999 {
				atempts++
				continue
			}

			p.bfr = make([]byte, size)
			numbytes = 0
			for {
				tmp_numbytes, err = p.Read(p.bfr[numbytes:])
				if err != nil {
					if err != io.EOF {
						//fmt.Printf("data read error: %s\n", err.Error())
					}
					p.Disconnect()
					return
				}
				numbytes += tmp_numbytes
				if numbytes < size {
					continue
				}
				break
			}

			err = json.Unmarshal(p.bfr, fbot)
			if err != nil {
				//fmt.Printf("json parse error: %s\n", err)
				p.Disconnect()
				return
			}

			router.Dispatch(p, fbot)
			//p.SetReadDeadline(time.Now().Add(time.Second * 5))
		}
	}
}
