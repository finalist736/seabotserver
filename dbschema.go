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
