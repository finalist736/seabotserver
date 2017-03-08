package seabotserver

type LogBattle struct {
	Winner int64        `json:"winner"`
	Sides  [2]*LogSides `json:"sides"`
	Turns  []*LogTurn   `json:"turns"`
}

type LogSides struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	Sea  *[100]int `json:"sea"`
}

type LogTurn struct {
	ID     int64  `json:"id"`
	Shot   [2]int `json:"shot"`
	Result int    `json:"result"`
}

type LoggingService interface {
	Store(*LogBattle) error
}
