package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DBConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port int64
}

type Configuration struct {
	DB        *DBConfig
	Mongo     *DBConfig
	Port      string
	Profiling bool
}

var cfg *Configuration = nil
var configFile string = ""

func SetConfigFile(file string) {
	configFile = file
}

func GetConfiguration() *Configuration {
	if cfg == nil {
		if configFile == "" {
			panic("empty config filename")
		}
		cfg = new(Configuration)
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(fmt.Sprintf("config read error: %s\n", err))
		}
		err = json.Unmarshal(data, cfg)
		if err != nil {
			panic(fmt.Sprintf("config parse error: %s\n", err))
		}
	}
	return cfg
}
