package battle

import (
	"time"

	"github.com/finalist736/seabotserver"
)

func (s *Battle) Handle(bot *seabotserver.TcpBot, turn *seabotserver.FBTurn) {
	s.BattleChannel <- &BattleChannelData{bot, turn}
}

func (s *Battle) Listener() {
	var data *BattleChannelData
	timer := time.Tick(time.Second * 10)
	tb := seabotserver.ToBot{}
	tb.Turn = &seabotserver.TBTurn{}

	tbNextTurn := seabotserver.ToBot{}
	tbNextTurn.Turn = &seabotserver.TBTurn{}
	var opponent *seabotserver.TcpBot
	for {
		select {
		case data = <-s.BattleChannel:
			//fmt.Printf("some data from bot: %d\n\t%+v\n", data.Bot.ID, data.Turn)
			if data.Turn == nil || data.Bot == nil {
				continue
			}
			var field *int
			if data.Bot.ID == s.Bot1.ID {
				field = &s.Pole2[data.Turn.Shot[0]][data.Turn.Shot[1]]
				opponent = s.Bot2
			} else {
				field = &s.Pole1[data.Turn.Shot[0]][data.Turn.Shot[1]]
				opponent = s.Bot1
			}
			if *field == 0 {
				tb.Turn.Result = -1
			} else {
				tb.Turn.Result = 1
			}
			// send shot result
			data.Bot.Send(tb)
			// send opponent shot result
			// -------------------------
			// send next turn
			if tb.Turn.Result == -1 {
				tbNextTurn.Turn.ID = opponent.ID
			} else {
				tbNextTurn.Turn.ID = data.Bot.ID
			}
			s.Bot1.Send(tbNextTurn)
			s.Bot2.Send(tbNextTurn)
		case <-timer:
			//fmt.Println("tick")
		}
	}
}
