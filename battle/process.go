package battle

import (
	"fmt"
	"time"

	"github.com/finalist736/seabotserver"
)

func (s *Battle) Handle(bot *seabotserver.TcpBot, turn *seabotserver.FBTurn) {
	s.BattleChannel <- &BattleChannelData{bot, turn}
}

func (s *Battle) Listener() {
	var data *BattleChannelData
	timer := time.Tick(time.Second * 10)
	for {
		select {
		case data = <-s.BattleChannel:
			fmt.Printf("some data from bot: %d\n\t%+v\n", data.Bot.ID, data.Turn)
		case <-timer:
			//fmt.Println("tick")
		}
	}
}
