package queue

import (
	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/battle"
)

func Handle(bot *seabotserver.TcpBot, bvb *seabotserver.FBBvb) {
	qChan <- &seabotserver.QueueData{bot, bvb, false}
}

func Exit(bot *seabotserver.TcpBot) {
	qChan <- &seabotserver.QueueData{bot, nil, true}
}

func channelHandler() {
	var data *seabotserver.QueueData
	for {
		select {
		case data = <-qChan:
			//fmt.Printf("setting queue goroutine: %+v\n", data)
			switch data.Exit {
			case true:
				if first == nil {
					continue
				}
				if data.Bot == nil {
					continue
				}
				if first.Bot.ID == data.Bot.ID {
					first = nil
				}
			case false:
				// set to queue!

				if data.Bot == nil || data.Bvb == nil {
					continue
				}
				// queue is empty let set bot!
				if first == nil {
					first = data

					tb := seabotserver.ToBot{}
					tb.Bvb = &seabotserver.TBBvb{}
					tb.Bvb.Wait = 1
					data.Bot.Send(tb)
				} else {
					if first.Bot.ID == data.Bot.ID {
						continue
					}
					battle.Create(first, data)
					first = nil
				}
			}
		}
	}
}
