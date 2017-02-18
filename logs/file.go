package logs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/finalist736/seabotserver"
)

func SaveToFile(l *seabotserver.LogBattle) {
	data, err := json.Marshal(l)
	if err != nil {
		fmt.Printf("logs.SaveToFile json error: %s\n", err)
		return
	}
	filename := fmt.Sprintf("./%s.json", time.Now().String())
	err = ioutil.WriteFile(filename, data, 0664)
	if err != nil {
		fmt.Printf("logs.SaveToFile save error: %s\n", err)
		return
	}
}
