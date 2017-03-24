package battle

import (
	"fmt"
	"time"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/storage/database/dbsql"
	"github.com/finalist736/seabotserver/storage/logs/mongodb"
)

func (s *Battle) Handle(bot seabotserver.BotService, turn *seabotserver.FBTurn) {
	s.BattleChannel <- &BattleChannelData{bot, turn, false}
}

func (s *Battle) Exit(bot seabotserver.BotService) {
	s.BattleChannel <- &BattleChannelData{bot, nil, true}
}

func (s *Battle) Listener() {
	var data *BattleChannelData
	var lastActive time.Time = time.Now()
	timer := time.Tick(time.Second * 10)
	tb := &seabotserver.ToBot{}
	tb.Turn = &seabotserver.TBTurn{}

	tbNextTurn := &seabotserver.ToBot{}
	tbNextTurn.Turn = &seabotserver.TBTurn{}

	tbOppTurn := &seabotserver.ToBot{}
	tbOppTurn.Turn = &seabotserver.TBTurn{}
	tbOppTurn.Turn.Opponent = &seabotserver.TBOpponentTurn{}

	var opponent seabotserver.BotService
	var isBattleEnd bool = false
	for {
		select {
		case data = <-s.BattleChannel:
			lastActive = time.Now()
			switch data.Exit {
			case true:
				// need to close this battle!
				errMsg := &seabotserver.ToBot{
					Error: &seabotserver.TBError{Error: "player disconnected"}}
				if s.Bot1.DBBot().ID == data.Bot.DBBot().ID {
					s.Bot2.Send(errMsg)
					s.Bot2.Disconnect()
				} else {
					s.Bot1.Send(errMsg)
					s.Bot1.Disconnect()
				}
				return

			case false:
				//fmt.Printf("some data from bot: %d\n\t%+v\n", data.Bot.DBBot().ID, data.Turn)
				if data.Turn == nil || data.Bot == nil {
					continue
				}

				if s.CurrentTurnID != data.Bot.DBBot().ID {
					data.Bot.Send(
						&seabotserver.ToBot{
							Error: &seabotserver.TBError{
								Error: fmt.Sprintf("not your turn! now turn bot id: %d. but your bot id: %d",
									s.CurrentTurnID, data.Bot.DBBot().ID)}})
					continue
				}

				if data.Turn.Shot[0] < 0 || data.Turn.Shot[0] > 9 ||
					data.Turn.Shot[1] < 0 || data.Turn.Shot[1] > 9 {
					data.Bot.Send(
						&seabotserver.ToBot{
							Error: &seabotserver.TBError{
								Error: "incorrect point"}})
					continue
				}
				var field *int
				var opppole *[10][10]int
				var oppShips *[10]*Ship
				if data.Bot.DBBot().ID == s.Bot1.DBBot().ID {
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
						isBattleEnd = true
					}

				}
				// save to log
				s.Log.Turns = append(s.Log.Turns,
					&seabotserver.LogTurn{
						data.Bot.DBBot().ID,
						data.Turn.Shot,
						tb.Turn.Result})

				// send shot result
				data.Bot.Send(tb)
				// send opponent shot result
				// -------------------------
				tbOppTurn.Turn.Opponent.Shot = data.Turn.Shot
				tbOppTurn.Turn.Opponent.Result = tb.Turn.Result
				opponent.Send(tbOppTurn)

				if isBattleEnd {
					// end battle
					tbEnd := &seabotserver.ToBot{}
					tbEnd.End = &seabotserver.TBEnd{}
					tbEnd.End.Winner = data.Bot.DBBot().ID
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

					s.Bot1.SetBattle(nil)
					s.Bot2.SetBattle(nil)

					//s.Bot1.Disconnect()
					//s.Bot2.Disconnect()

					// save battle result to log
					s.Log.Winner = tbEnd.End.Winner
					s.Log.EndTime = time.Now().Unix()
					// statistics save to DB
					logserv := mongodb.NewLoggingService()
					logserv.Store(s.Log)
					// sandbox counters
					if data.Bot.DBBot().ID < 0 || opponent.DBBot().ID < 0 {

						id := data.Bot.DBBot().ID
						if id < 0 {
							id = opponent.DBBot().ID
						}
						if id > 0 {
							bvsaiService := dbsql.NewDBBVsaiService()
							bvsaiData := bvsaiService.Get(id)
							bvsaiData.Bvr++
							bvsaiService.Store(bvsaiData)
						}
					} else {
						var dbsbx *seabotserver.DBSandbox
						sbservice := dbsql.NewDBSandboxService()
						dbsbx = sbservice.Get(data.Bot.DBBot().ID)
						dbsbx.Wins++
						sbservice.Store(dbsbx)
						dbsbx = sbservice.Get(opponent.DBBot().ID)
						dbsbx.Loses++
						sbservice.Store(dbsbx)
					}

					return
				} else {
					// send next turn
					if tb.Turn.Result == -1 {
						tbNextTurn.Turn.ID = opponent.DBBot().ID
						s.CurrentTurnID = opponent.DBBot().ID
					} else {
						tbNextTurn.Turn.ID = data.Bot.DBBot().ID
						s.CurrentTurnID = data.Bot.DBBot().ID
					}
					s.Bot1.Send(tbNextTurn)
					s.Bot2.Send(tbNextTurn)
				}
			}
		case tick := <-timer:
			//fmt.Printf("tick: %s; diff: %s\n", tick, tick.Sub(lastActive))
			if tick.Sub(lastActive) > time.Minute*60 {
				// check and set lose guilty
				fmt.Printf("%d - timeout GUILTY\n", s.CurrentTurnID)
				return
			}
		}
	}
}
