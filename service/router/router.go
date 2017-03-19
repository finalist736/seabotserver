package router

import (
	"fmt"
	"runtime"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/gameplay/battle"
	"github.com/finalist736/seabotserver/gameplay/queue"
	"github.com/finalist736/seabotserver/storage/database/dbsql"
	"github.com/finalist736/seabotserver/tcpserver/bots/aibot"
)

func Dispatch(bot seabotserver.BotService, fbot *seabotserver.FromBot) {
	if bot.DBBot().ID == 0 {
		if fbot.Auth == "" {
			return
		} else {
			dbBotService := dbsql.NewDBBotService()
			dbbot, err := dbBotService.Auth(fbot.Auth)
			//fmt.Printf("after auth: %+v\n", dbbot)
			if err != nil {
				bot.SendError("auth error, register on http://finalistx.com/ for new key")
				bot.Disconnect()
				return
			}
			bot.SetDBBot(dbbot)

			tb := &seabotserver.ToBot{}
			tb.Auth = &seabotserver.TBAuth{}
			tb.Auth.OK = true
			tb.Auth.ID = bot.DBBot().ID
			tb.Auth.User = bot.DBBot().User

			bot.Send(tb)
		}
	} else if bot.Battle() == nil {
		if fbot.Bvb != nil {
			// stay to queue
			fmt.Printf("setting to queue: %+v\n", fbot)
			queue.Handle(bot, fbot.Bvb)
		} else if fbot.Bvr != nil {
			f := &seabotserver.QueueData{bot, fbot.Bvr, false}
			s := &seabotserver.QueueData{aibot.NewBot(), &seabotserver.FBBvb{0, nil}, false}
			battle.Create(f, s)
		} else if fbot.Exit {
			bot.Disconnect()
			return
		} else if fbot.Profile != nil {
			prof := &seabotserver.TBProfile{}
			prof.Gnum = runtime.NumGoroutine()
			bot.Send(&seabotserver.ToBot{Profile: prof})
		}
	} else {
		if fbot.Turn == nil {
			return
		}
		if bot.Battle() == nil {
			return
		}
		btl := bot.Battle().(*battle.Battle)
		if btl == nil {
			return
		}
		btl.Handle(bot, fbot.Turn)
		// need to handle exit?
	}

}
