package queue

import (
	"fmt"

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
			fmt.Printf("1 queue: %v\n", data.Bot.ID)
			switch data.Exit {
			case true:
				fmt.Printf("2 queue exit: %v\n", data.Bot.ID)
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
				fmt.Printf("2 queue set: %v\n", data.Bot.ID)
				// queue is empty let set bot!
				if first == nil {
					fmt.Printf("3 queue first: %v\n", data.Bot.ID)
					first = data

					tb := seabotserver.ToBot{}
					tb.Bvb = &seabotserver.TBBvb{}
					tb.Bvb.Wait = 1
					data.Bot.Send(tb)
				} else {
					fmt.Printf("4 queue second: %v\n", data.Bot.ID)
					if first.Bot.ID == data.Bot.ID {
						continue
					}
					battle.Create(first, data)
					fmt.Printf("4 queue created : %v + %v\n", data.Bot.ID, first.Bot.ID)
					first = nil
				}
			}
		}
	}
}
