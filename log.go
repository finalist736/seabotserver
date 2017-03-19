package seabotserver

type LogBattle struct {
	BattleID  int64        `bson:"battle"`
	StartTime int64        `bson:"start"`
	EndTime   int64        `bson:"end"`
	Winner    int64        `bson:"winner"`
	Sides     [2]*LogSides `bson:"sides"`
	Turns     []*LogTurn   `bson:"turns"`
}

type LogSides struct {
	ID   int64     `bson:"id"`
	Name string    `bson:"name"`
	Sea  *[100]int `bson:"sea"`
}

type LogTurn struct {
	ID     int64  `bson:"id"`
	Shot   [2]int `bson:"shot"`
	Result int    `bson:"result"`
}

type LoggingService interface {
	Store(*LogBattle) error
}
