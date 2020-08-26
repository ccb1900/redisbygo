package db

import (
	"github.com/ccb1900/redisbygo/app/server/constructor"
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/client"
)

func SelectDb(cl *client.Client, id int) {
	c := pkg.NewConfig()
	s := constructor.NewServer()
	if id < 0 || id >= c.Dbnum {
		cl.Log.Info("db num err")
	} else {
		cl.Db = s.Db[id]
	}
}
