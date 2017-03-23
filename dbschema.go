package seabotserver

type DBBot struct {
	ID      int64  `db:"id"`
	User    int64  `db:"user_id"`
	AuthKey string `db:"auth_key"`
}

func NewDBBot() *DBBot {
	return &DBBot{}
}

type DBBotService interface {
	Auth(string) (*DBBot, error)
}

type DBSandbox struct {
	Bot   int64 `db:"bot"`
	Wins  int64 `db:"wins"`
	Loses int64 `db:"loses"`
	Last  int64 `db:"last"`
}

type DBSandboxService interface {
	Get(botid int64) *DBSandbox // return sandox record for botID, if not exists - return default
	Store(*DBSandbox) error     // saves data to database, changes Last to now
}
