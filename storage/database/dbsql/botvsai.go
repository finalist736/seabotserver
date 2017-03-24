package dbsql

import (
	"fmt"

	"github.com/finalist736/seabotserver"
	"github.com/gocraft/dbr"
)

type BVsaiService struct {
}

func NewDBBVsaiService() seabotserver.DBVsaiService {
	return &BVsaiService{}
}

func (s *BVsaiService) Get(botid int64) *seabotserver.DBVsai {
	record := &seabotserver.DBVsai{}
	err := session().
		Select("*").
		From("botvsai").
		Where("`bot`=?", botid).
		LoadValue(record)
	if err != nil {
		if err != dbr.ErrNotFound {
			fmt.Printf("VsaiService.Get - select error: %s\n", err)
		}
		record.Bot = botid
		_, err = session().
			InsertInto("botvsai").
			Columns("bot", "bvr", "bvl").
			Record(record).
			Exec()
		if err != nil {
			fmt.Printf("VsaiService.Get - insert error: %s\n", err)
		}
	}
	return record
}

func (s *BVsaiService) Store(bva *seabotserver.DBVsai) error {
	_, err := session().
		Update("botvsai").
		Set("bvr", bva.Bvr).
		Set("bvl", bva.Bvl).
		Where("`bot`=?", bva.Bot).
		Exec()
	if err != nil {
		fmt.Printf("VsaiService.Store - update error: %s\n", err)
		return err
	}
	return nil
}
