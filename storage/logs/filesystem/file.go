package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/finalist736/seabotserver"
)

type LoggingService struct {
}

func NewLoggingService() seabotserver.LoggingService {
	return &LoggingService{}
}

func (*LoggingService) Store(l *seabotserver.LogBattle) error {
	data, err := json.Marshal(l)
	if err != nil {
		fmt.Printf("logs.SaveToFile json error: %s\n", err)
		return err
	}
	filename := fmt.Sprintf("./%s.json", time.Now().String())
	err = ioutil.WriteFile(filename, data, 0664)
	if err != nil {
		fmt.Printf("logs.SaveToFile save error: %s\n", err)
		return err
	}
	return nil
}
