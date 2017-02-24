package battle

import (
	"time"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/logs"
)

func (s *Battle) Handle(bot *seabotserver.TcpBot, turn *seabotserver.FBTurn) {
	s.BattleChannel <- &BattleChannelData{bot, turn, false}
}

func (s *Battle) Exit(bot *seabotserver.TcpBot) {
	s.BattleChannel <- &BattleChannelData{bot, nil, true}
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

			switch data.Exit {
			case true:
				// need to close this battle!
				errMsg := &seabotserver.ToBot{
					Error: &seabotserver.TBError{Error: "player disconnected"}}
				if s.Bot1.ID == data.Bot.ID {
					s.Bot2.Send(errMsg)
					s.Bot2.Done <- true
				} else {
					s.Bot1.Send(errMsg)
					s.Bot1.Done <- true
				}
				return

			case false:
				//fmt.Printf("some data from bot: %d\n\t%+v\n", data.Bot.ID, data.Turn)
				if data.Turn == nil || data.Bot == nil {
					continue
				}

				if s.CurrentTurnID != data.Bot.ID {
					data.Bot.Send(&seabotserver.ToBot{Error: &seabotserver.TBError{Error: "incorrect turn"}})
					continue
				}

				if data.Turn.Shot[0] < 0 || data.Turn.Shot[0] > 9 ||
					data.Turn.Shot[1] < 0 || data.Turn.Shot[1] > 9 {
					data.Bot.Send(&seabotserver.ToBot{Error: &seabotserver.TBError{Error: "incorrect point"}})
					continue
				}
				var field *int
				var opppole *[10][10]int
				var oppShips *[10]*Ship
				if data.Bot.ID == s.Bot1.ID {
					field = &s.Pole2[data.Turn.Shot[0]][data.Turn.Shot[1]]
					opppole = s.Pole2
					opponent = s.Bot2
					oppShips = s.Ships2
				} else {
					field = &s.Pole1[data.Turn.Shot[0]][data.Turn.Shot[1]]
					opppole = s.Pole1
					opponent = s.Bot1
					oppShips = s.Ships1
				}
				if *field == 0 {
					tb.Turn.Result = -1
					*field = -10
				} else if *field == -10 {
					// calc misscount
					// and ban bot when count == 5
				} else if *field > -5 && *field < 0 {
					// shot in dead ship
					// and ban bot when count == 5
				} else {
					// check for ship dead and send 2
					tb.Turn.Result = checkShipDestroy(oppShips, data.Turn.Shot, *field)
					*field *= -1

					// check for battle end
					// -------------------------
					if checkFleet(opppole) {
						// end battle
						tbEnd := seabotserver.ToBot{}
						tbEnd.End = &seabotserver.TBEnd{}
						tbEnd.End.Winner = data.Bot.ID
						tbEnd.End.Ships = FormatShips(opppole)
						data.Bot.Send(tbEnd)
						if opppole == s.Pole1 {
							tbEnd.End.Ships = FormatShips(s.Pole2)
						} else {
							tbEnd.End.Ships = FormatShips(s.Pole1)
						}
						opponent.Send(tbEnd)

						//fmt.Println("POLE 1")
						//PrintPole(s.Pole1, s.Bot1.ID)

						//fmt.Println("POLE 2")
						//PrintPole(s.Pole2, s.Bot2.ID)

						s.Bot1.Battle = nil
						s.Bot2.Battle = nil

						s.Bot1.Done <- true
						s.Bot2.Done <- true

						// save battle result to log
						s.Log.Winner = tbEnd.End.Winner
						logs.SaveToFile(s.Log)
						// statistics save to DB
						return
					}

				}
				// save to log
				s.Log.Turns = append(s.Log.Turns,
					&seabotserver.LogTurn{
						data.Bot.ID,
						data.Turn.Shot,
						tb.Turn.Result})

				// send shot result
				data.Bot.Send(tb)
				// send opponent shot result
				// -------------------------
				tbOppTurn.Turn.Opponent.Shot = data.Turn.Shot
				tbOppTurn.Turn.Opponent.Result = tb.Turn.Result
				opponent.Send(tbOppTurn)

				// send next turn
				if tb.Turn.Result == -1 {
					tbNextTurn.Turn.ID = opponent.ID
					s.CurrentTurnID = opponent.ID
				} else {
					tbNextTurn.Turn.ID = data.Bot.ID
					s.CurrentTurnID = data.Bot.ID
				}
				s.Bot1.Send(tbNextTurn)
				s.Bot2.Send(tbNextTurn)
			}
		case <-timer:
			//fmt.Println("tick")
		}
	}
}
