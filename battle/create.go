package battle

import (
	"fmt"
	"math/rand"

	"github.com/finalist736/seabotserver"
)

func Create(q1, q2 *seabotserver.QueueData) {
	nb := NewBattle()
	nb.Bot1 = q1.Bot
	nb.Bot2 = q2.Bot

	// first bot's ships
	if q1.Bvb.Place == 0 {
		// place ships by server
		nb.Pole1, nb.Ships1 = PlaceShips()
		//PrintPole(nb.Pole1)
	} else {
		// need to check placement by bot
		// feature! place ships by server anyway
		nb.Pole1, nb.Ships1 = PlaceShips()
	}

	// second bot's ships
	if q2.Bvb.Place == 0 {
		// place ships by server
		nb.Pole2, nb.Ships2 = PlaceShips()
		//PrintPole(nb.Pole2)
	} else {
		// need to check
		// if places incorrect then break battle and disconnect bots;
		// feature! place ships by server anyway
		nb.Pole2, nb.Ships2 = PlaceShips()
	}

	//	for _, ship := range nb.Ships1 {
	//		fmt.Printf("%+v\n", ship)
	//	}
	//fmt.Printf("%+v\n", nb.Ships2)

	// response init
	tb := seabotserver.ToBot{}
	tb.Bvb = &seabotserver.TBBvb{}
	// to first bot
	tb.Bvb.ID = q2.Bot.DBBot.ID
	tb.Bvb.Name = fmt.Sprintf("bot_%d", tb.Bvb.ID)
	tb.Bvb.Ships = FormatShips(nb.Pole1)
	q1.Bot.Send(tb)
	// to second bot
	tb.Bvb.ID = q1.Bot.DBBot.ID
	tb.Bvb.Name = fmt.Sprintf("bot_%d", tb.Bvb.ID)
	tb.Bvb.Ships = FormatShips(nb.Pole2)
	q2.Bot.Send(tb)

	q1.Bot.Battle = nb
	q2.Bot.Battle = nb

	if rand.Int31n(2) == 0 {
		nb.CurrentTurnID = q1.Bot.DBBot.ID
	} else {
		nb.CurrentTurnID = q2.Bot.DBBot.ID
	}

	tb.Bvb = nil
	tb.Turn = &seabotserver.TBTurn{}
	tb.Turn.ID = nb.CurrentTurnID
	q1.Bot.Send(tb)
	q2.Bot.Send(tb)

	nb.Log.Sides[0] = &seabotserver.LogSides{}
	nb.Log.Sides[1] = &seabotserver.LogSides{}

	nb.Log.Sides[0].ID = nb.Bot1.DBBot.ID
	nb.Log.Sides[0].Name = "test"
	nb.Log.Sides[0].Sea = FormatShips(nb.Pole1)

	nb.Log.Sides[1].ID = nb.Bot2.DBBot.ID
	nb.Log.Sides[1].Name = "test2"
	nb.Log.Sides[1].Sea = FormatShips(nb.Pole2)

	go nb.Listener()
}
