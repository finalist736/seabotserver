package router

import "github.com/finalist736/seabotserver"

func Dispatch(bot *seabotserver.TcpBot) {
	//bot.Done <- true
	bot.Send <- []byte("12345")
}
