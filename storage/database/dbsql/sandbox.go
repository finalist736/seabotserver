package dbsql

import (
	"fmt"
	"time"

	"github.com/finalist736/seabotserver"
	"github.com/gocraft/dbr"
)

type SandboxService struct {
}

func NewDBSandboxService() seabotserver.DBSandboxService {
	return &SandboxService{}
}

func (s *SandboxService) Get(botid int64) *seabotserver.DBSandbox {
	record := &seabotserver.DBSandbox{}
	err := session().Select("*").From("sandbox").Where("`bot`=?", botid).LoadValue(record)
	if err != nil {
		if err != dbr.ErrNotFound {
			fmt.Printf("SandboxService.Get - select error: %s\n", err)
		}
		record.Bot = botid
		record.Last = time.Now().Unix()
		_, err = session().InsertInto("sandbox").
			Columns("bot", "wins", "loses", "last").
			Record(record).Exec()
		if err != nil {
			fmt.Printf("SandboxService.Get - insert error: %s\n", err)
		}
	}
	return record
}

func (s *SandboxService) Store(sb *seabotserver.DBSandbox) error {
	sb.Last = time.Now().Unix()
	_, err := session().Update("sandbox").
		Set("wins", sb.Wins).Set("loses", sb.Loses).
		Set("last", sb.Last).Where("`bot`=?", sb.Bot).Exec()
	if err != nil {
		fmt.Printf("SandboxService.Store - update error: %s\n", err)
		return err
	}
	return nil
}
