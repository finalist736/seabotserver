package router

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/finalist736/seabotserver"
)

var id int64

func Dispatch(bot *seabotserver.TcpBot) {

	fbot := &seabotserver.FromBot{}
	err := json.Unmarshal(bot.Buffer[0:bot.BufferLen], fbot)
	if err != nil {
		bot.Done <- true
		fmt.Printf("json parse error: %s", err)
		return
	}

	if bot.ID == 0 {
		if fbot.Auth == "" {
			fmt.Printf("authkey empty: %s", bot.RemoteAddr())
			// send error to bot
			return
		} else {
			tb := seabotserver.ToBot{}
			tb.Auth = &seabotserver.TBAuth{}
			tb.Auth.OK = true
			tb.Auth.ID = atomic.AddInt64(&id, 1)
			bot.Send(tb)
		}
	} else {
		if fbot.Turn != nil {
			// make turn
		} else if fbot.Bvb != nil {
			// stay to queue
		} else if fbot.Exit {
			bot.Done <- true
			return
		}
	}

}
