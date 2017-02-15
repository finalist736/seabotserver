package tcpserver

import (
	"fmt"
	"io"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/battle"
	"github.com/finalist736/seabotserver/queue"
	"github.com/finalist736/seabotserver/router"
)

func sender(p *seabotserver.TcpBot) {
	defer p.Close()
	defer func() { fmt.Printf("sender close: %v\n", p.RemoteAddr()) }()
	var err error
	var n int

	var size int
	var send_buff []byte

	for {
		select {
		case <-p.Done:
			return
		case msg := <-p.SendChannel:

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
	defer func() { queue.Exit(p) }()
	defer func() {
		if p.Battle == nil {
			return
		}
		btl := p.Battle.(*battle.Battle)
		if btl == nil {
			return
		}
		btl.Exit(p)
	}()

	tmp_buffer := make([]byte, 4)
	var numbytes, tmp_numbytes, size, atempts int
	var err error
	numbytes = 0
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
						fmt.Printf("4 bytes read error: %s\n", err.Error())
					}
					p.Done <- true
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

			p.Buffer = make([]byte, size)
			numbytes = 0
			for {
				tmp_numbytes, err = p.Read(p.Buffer[numbytes:])
				if err != nil {
					if err != io.EOF {
						fmt.Printf("data read error: %s\n", err.Error())
					}
					p.Done <- true
					return
				}
				numbytes += tmp_numbytes
				if numbytes < size {
					continue
				}
				break
			}
			router.Dispatch(p)
			//p.SetReadDeadline(time.Now().Add(time.Second * 5))
		}
	}
}
