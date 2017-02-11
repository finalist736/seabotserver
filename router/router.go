package router

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/battle"
	"github.com/finalist736/seabotserver/queue"
)

var id int64

func Dispatch(bot *seabotserver.TcpBot) {

	fbot := &seabotserver.FromBot{}
	err := json.Unmarshal(bot.Buffer, fbot)
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
			// TODO
			// DATABASE AUTH AND ERROR
			tb := seabotserver.ToBot{}
			tb.Auth = &seabotserver.TBAuth{}
			tb.Auth.OK = true
			tb.Auth.ID = atomic.AddInt64(&id, 1)
			// remember bot id
			bot.ID = tb.Auth.ID

			bot.Send(tb)
		}
	} else if bot.Battle == nil {
		if fbot.Bvb != nil {
			// stay to queue
			fmt.Printf("setting to queue: %+v\n", fbot)
			queue.Handle(bot, fbot.Bvb)
		} else if fbot.Exit {
			bot.Done <- true
			queue.Exit(bot)
			return
		}
	} else {
		if fbot.Turn == nil {
			return
		}
		btl := bot.Battle.(*battle.Battle)
		if btl == nil {
			return
		}
		btl.Handle(bot, fbot.Turn)
		// need to handle exit?
	}

}
