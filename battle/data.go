package battle

import (
	"math/rand"
	"time"

	"github.com/finalist736/seabotserver"
)

type BattleChannelData struct {
	Bot  *seabotserver.TcpBot
	Turn *seabotserver.FBTurn
}

type Battle struct {
	ID int64

	Bot1 *seabotserver.TcpBot
	Bot2 *seabotserver.TcpBot

	BattleChannel chan *BattleChannelData

	CurrentTurnID int64
	Pole1         [10][10]int
	Pole2         [10][10]int
}

var rnd *rand.Rand

func NewBattle() *Battle {
	return &Battle{BattleChannel: make(chan *BattleChannelData)}
}

func init() {
	rnd = rand.New(rand.NewSource(time.Now().Unix()))
}
