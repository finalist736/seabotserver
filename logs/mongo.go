package logs

import (
	"fmt"

	"github.com/finalist736/seabotserver"
	"github.com/finalist736/seabotserver/config"
	"gopkg.in/mgo.v2"
)

func SaveToMongoDB(l *seabotserver.LogBattle) {
	// mongodb://<dbuser>:<dbpassword>@ds145289.mlab.com:45289/navigation
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
		return
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(conf.Mongo.Name).C("games")
	err = c.Insert(l)
	if err != nil {
		fmt.Printf("mongo insert error: %s", err.Error())
		return
	}
}
