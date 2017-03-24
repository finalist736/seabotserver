package aibot

import (
	"math/rand"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/gameplay/battle"
)

type AIBot struct {
	dbbot *seabotserver.DBBot
	btl   interface{}

	trns []*seabotserver.FBTurn
}

func NewBot() seabotserver.BotService {
	bot := &AIBot{}
	bot.dbbot = &seabotserver.DBBot{}

	bot.dbbot.ID = -1
	bot.dbbot.User = -1

	bot.FieldInit()
	return bot
}

func (s *AIBot) FieldInit() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			s.trns = append(s.trns, &seabotserver.FBTurn{[2]int{i, j}})
		}
	}
}

func (s *AIBot) Battle() interface{} {
	return s.btl
}

func (s *AIBot) SetBattle(b interface{}) {
	s.btl = b
}

func (s *AIBot) DBBot() *seabotserver.DBBot {
	return s.dbbot
}

func (s *AIBot) SetDBBot(b *seabotserver.DBBot) {
	s.dbbot = b
}

func (s *AIBot) Send(d interface{}) {
	if d == nil {
		return
	}
	tb := d.(*seabotserver.ToBot)
	if tb == nil {
		return
	}
	if tb.End != nil {
		//fmt.Printf("end: %+v\n", tb.End)
		return
	}
	if tb.Turn == nil {
		return
	}
	//fmt.Printf("send: %+v\n", tb.Turn)
	if tb.Turn.Result != 0 {
		//fmt.Printf("result: %+v\n", tb.Turn.Result)
		return
	}
	if tb.Turn.Opponent != nil {
		//fmt.Printf("opp: %+v\n", tb.Turn.Opponent)
		return
	}
	if tb.Turn.ID != s.dbbot.ID {
		return
	}
	if s.Battle() == nil {
		return
	}
	btl := s.Battle().(*battle.Battle)
	if btl == nil {
		return
	}
	var shot *seabotserver.FBTurn
	ln := len(s.trns)
	if ln > 0 {
		item := rand.Intn(ln)
		shot = s.trns[item]
		s.trns = append(s.trns[:item], s.trns[item+1:]...)
	} else {
		shot = &seabotserver.FBTurn{[2]int{0, 0}}
	}
	//fmt.Printf("aishoot: %+v\n", shot)
	btl.Handle(s, shot)
}

func (s *AIBot) SendError(err string) {

}

func (s *AIBot) Disconnect() {

}

func (p *AIBot) Sender() {

}

func (p *AIBot) Handler() {
	//	defer func() {
	//		if p.Battle() == nil {
	//			return
	//		}
	//		btl := p.Battle().(*battle.Battle)
	//		if btl == nil {
	//			return
	//		}
	//		btl.Exit(p)
	//	}()

}
