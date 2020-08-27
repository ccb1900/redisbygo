package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Port        int
	Host        string
	Timezone    string
	Dbnum       int
	Maxclients  int
	Name        string
	AppendOnly  bool
	SaveSeconds int
	SaveTimes   int
	DbFileName  string
	Dir         string
}

var instance *Config
var once sync.Once

func GetInstance(path string) *Config {
	once.Do(func() {
		instance = new(Config)
		viper.SetConfigFile(path)
		viper.SetConfigType("json")
		//viper.AddConfigPath(".")

		err := viper.ReadInConfig()

		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		instance.Port = viper.GetInt("Port")
		instance.Host = viper.GetString("Host")
		instance.Timezone = viper.GetString("Timezone")
		instance.Dbnum = viper.GetInt("Dbnum")
		instance.Maxclients = viper.GetInt("maxclients")
		instance.Name = viper.GetString("Name")
		instance.AppendOnly = viper.GetBool("Appendonly")
		instance.SaveSeconds = viper.GetInt("save_seconds")
		instance.SaveTimes = viper.GetInt("save_times")
		instance.DbFileName = viper.GetString("dbfilename")
		instance.Dir = viper.GetString("dir")
	})

	return instance
}
func NewConfig() *Config {
	return GetInstance("./server.json")
}
