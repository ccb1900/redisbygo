package db

import (
	"redis/app/server/constructor"
	"redis/pkg/client"
	"redis/pkg/config"
)

func SelectDb(cl *client.Client, id int) {
	c := config.NewConfig()
	s := constructor.NewServer()
	if id < 0 || id >= c.Dbnum {
		cl.Log.Info("db num err")
	} else {
		cl.Db = s.Db[id]
	}
}
