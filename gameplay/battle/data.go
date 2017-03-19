package battle

import "github.com/finalist736/seabotserver"

type BattleChannelData struct {
	Bot  seabotserver.BotService
	Turn *seabotserver.FBTurn
	Exit bool
}

type Ship struct {
	Count int
	Place [][2]int
}

type Battle struct {
	ID int64

	Bot1 seabotserver.BotService
	Bot2 seabotserver.BotService

	BattleChannel chan *BattleChannelData

	CurrentTurnID int64
	Pole1         *[10][10]int
	Ships1        *[10]*Ship
	Pole2         *[10][10]int
	Ships2        *[10]*Ship

	Log *seabotserver.LogBattle
}

func NewBattle() *Battle {
	return &Battle{
		BattleChannel: make(chan *BattleChannelData, 2),
		Log:           &seabotserver.LogBattle{}}
}
