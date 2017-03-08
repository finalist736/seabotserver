package router

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/battle"
	"github.com/finalist736/seabotserver/database/dbsql"
	"github.com/finalist736/seabotserver/queue"
)

func Dispatch(bot *seabotserver.TcpBot) {

	fbot := &seabotserver.FromBot{}
	err := json.Unmarshal(bot.Buffer, fbot)
	if err != nil {
		bot.Disconnect()
		fmt.Printf("json parse error: %s", err)
		return
	}

	if bot.DBBot.ID == 0 {
		if fbot.Auth == "" {
			fmt.Printf("authkey empty: %s", bot.RemoteAddr())
			// send error to bot
			return
		} else {
			dbBotService := dbsql.NewDBBotService()
			bot.DBBot, err = dbBotService.Auth(fbot.Auth)
			if err != nil {
				bot.SendError("auth error, register on http://finalistx.com/ for new key")
				bot.Disconnect()
				return
			}

			tb := &seabotserver.ToBot{}
			tb.Auth = &seabotserver.TBAuth{}
			tb.Auth.OK = true
			tb.Auth.ID = bot.DBBot.ID
			tb.Auth.User = bot.DBBot.User

			bot.Send(tb)
		}
	} else if bot.Battle == nil {
		if fbot.Bvb != nil {
			// stay to queue
			fmt.Printf("setting to queue: %+v\n", fbot)
			queue.Handle(bot, fbot.Bvb)
		} else if fbot.Exit {
			bot.Disconnect()
			queue.Exit(bot)
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
		if bot.Battle == nil {
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
