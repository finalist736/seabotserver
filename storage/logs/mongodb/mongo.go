package mongodb

import (
	"fmt"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/storage/config"
	"gopkg.in/mgo.v2"
)

type LoggingService struct {
}

func NewLoggingService() seabotserver.LoggingService {
	return &LoggingService{}
}

func (*LoggingService) Store(l *seabotserver.LogBattle) error {
	conf := config.GetConfiguration()
	connectString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		conf.Mongo.User,
		conf.Mongo.Pass,
		conf.Mongo.Host,
		conf.Mongo.Port,
		conf.Mongo.Name)

	session, err := mgo.Dial(connectString)
	if err != nil {
		fmt.Printf("mongo connection error: %s", err.Error())
		return err
	}
	defer session.Close()

	c := session.DB(conf.Mongo.Name).C("games")
	err = c.Insert(l)
	if err != nil {
		fmt.Printf("mongo insert error: %s", err.Error())
		return err
	}
	// Index
	index := mgo.Index{
		Key:        []string{"battle"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		fmt.Printf("mongo index error: %s", err.Error())
		return err
	}
	return nil
}
