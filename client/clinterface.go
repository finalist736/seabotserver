package client

//type pole int[20]

type Bot struct {
	sea          [][]int
	lastX, lastY int
}

type Shoot struct {
	x, y int
}

type PlayerInfo struct {
	ID   int64
	Name string
}

type BotInterface interface {
	Auth() string
	Start(*PlayerInfo)
	Turn() *Shoot
	TurnResult(*Shoot, int)
}
