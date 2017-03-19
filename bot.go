package seabotserver

type BotService interface {
	Battle() interface{}
	SetBattle(b interface{})

	SetDBBot(*DBBot)
	DBBot() *DBBot

	Send(d interface{})
	SendError(err string)

	Sender()
	Handler()

	Disconnect()
}
