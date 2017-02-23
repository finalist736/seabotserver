package battle

import (
	"math/rand"
	"time"

	"github.com/finalist736/seabotserver"
)

type BattleChannelData struct {
	Bot  *seabotserver.TcpBot
	Turn *seabotserver.FBTurn
	Exit bool
}

type Ship struct {
	Count int
	Place [][2]int
}

type Battle struct {
	ID int64

	Bot1 *seabotserver.TcpBot
	Bot2 *seabotserver.TcpBot

	BattleChannel chan *BattleChannelData

	CurrentTurnID int64
	Pole1         *[10][10]int
	Ships1        *[10]*Ship
	Pole2         *[10][10]int
	Ships2        *[10]*Ship

	Log *seabotserver.LogBattle
}

var rnd *rand.Rand

func NewBattle() *Battle {
	return &Battle{
		BattleChannel: make(chan *BattleChannelData, 2),
		Log:           &seabotserver.LogBattle{}}
}

func init() {
	rnd = rand.New(rand.NewSource(time.Now().Unix()))
}
