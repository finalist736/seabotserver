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
	tb := seabotserver.ToBot{}
	tb.Turn = &seabotserver.TBTurn{}

	tbNextTurn := seabotserver.ToBot{}
	tbNextTurn.Turn = &seabotserver.TBTurn{}

	tbOppTurn := seabotserver.ToBot{}
	tbOppTurn.Turn = &seabotserver.TBTurn{}
	tbOppTurn.Turn.Opponent = &seabotserver.TBOpponentTurn{}

	var opponent *seabotserver.TcpBot
	for {
		select {
		case data = <-s.BattleChannel:
			//fmt.Printf("some data from bot: %d\n\t%+v\n", data.Bot.ID, data.Turn)
			if data.Turn == nil || data.Bot == nil {
				continue
			}
			if data.Turn.Shot[0] < 0 || data.Turn.Shot[0] > 9 ||
				data.Turn.Shot[1] < 0 || data.Turn.Shot[1] > 9 {
				data.Bot.Send(&seabotserver.ToBot{Error: &seabotserver.TBError{Error: "incorrect point"}})
				continue
			}
			var field *int
			var opppole *[10][10]int
			if data.Bot.ID == s.Bot1.ID {
				field = &s.Pole2[data.Turn.Shot[0]][data.Turn.Shot[1]]
				opppole = s.Pole2
				opponent = s.Bot2
			} else {
				field = &s.Pole1[data.Turn.Shot[0]][data.Turn.Shot[1]]
				opppole = s.Pole1
				opponent = s.Bot1
			}
			if *field == 0 {
				tb.Turn.Result = -1
			} else {
				tb.Turn.Result = 1
				*field *= -1
			}
			// send shot result
			data.Bot.Send(tb)
			// send opponent shot result
			// -------------------------
			tbOppTurn.Turn.Opponent.Shot = data.Turn.Shot
			tbOppTurn.Turn.Opponent.Result = tb.Turn.Result
			opponent.Send(tbOppTurn)

			// check for battle end
			// -------------------------
			if checkFleet(opppole) {
				// end battle
				tbEnd := seabotserver.ToBot{}
				tbEnd.End = &seabotserver.TBEnd{}
				tbEnd.End.Winner = data.Bot.ID

				fmt.Println("POLE 1")
				PrintPole(s.Pole1)

				fmt.Println("POLE 2")
				PrintPole(s.Pole2)

				data.Bot.Send(tbEnd)
				opponent.Send(tbEnd)

				s.Bot1.Done <- true
				s.Bot2.Done <- true

				return
			}

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
