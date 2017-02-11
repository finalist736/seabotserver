package queue

import "github.com/finalist736/seabotserver"

var qChan chan *seabotserver.QueueData
var first *seabotserver.QueueData

func init() {
	qChan = make(chan *seabotserver.QueueData, 10)
	go channelHandler()
}
