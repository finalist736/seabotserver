package dbsql

import (
	"errors"
	"fmt"

	"github.com/finalist736/seabotserver"
	"github.com/gocraft/dbr"
)

type BotService struct {
}

func NewDBBotService() seabotserver.DBBotService {
	return &BotService{}
}

func (s *BotService) Auth(key string) (*seabotserver.DBBot, error) {
	if key == "" {
		return nil, errors.New("key is empty")
	}
	bot := seabotserver.NewDBBot()
	err := session().Select("*").From("bots").Where("`auth_key`=?", key).LoadStruct(bot)
	if err != nil {
		if err != dbr.ErrNotFound {
			fmt.Printf("select from bots error: %s\n", err)
		}
		return nil, err
	}
	return bot, nil
}
