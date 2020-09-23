package config

import (
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/utils"
	"github.com/spf13/viper"
	"os"
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
		if utils.Exists(path) {
			viper.SetConfigFile(path)
			viper.SetConfigType("json")

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
		} else {
			instance.Port = 6378
			instance.Host = "127.0.0.1"
			instance.Timezone = "PRC"
			instance.Dbnum = 16
			instance.Maxclients = 10000
			instance.Name = "redisbygo"
			instance.AppendOnly = true
			instance.SaveSeconds = 1000
			instance.SaveTimes = 1000
			instance.DbFileName = "appendonly.aof"
			instance.Dir = "./"
		}
	})

	return instance
}
func NewConfig() *Config {
	return GetInstance(getConfigPath())
}

func getConfigPath() string {
	fmt.Println(os.LookupEnv("REDISBYGO_ROOT"))
	fmt.Println(os.Getenv("REDISBYGO_ROOT") + "/" + "./server.json")
	return os.Getenv("REDISBYGO_ROOT") + "/" + "./server.json"
}
