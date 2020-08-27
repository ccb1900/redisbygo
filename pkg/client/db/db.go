package db

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/config"
)

func SelectDb(cl *pkg.Client, id int) {
	c := config.NewConfig()
	s := pkg.NewServer()
	if id < 0 || id >= c.Dbnum {
		cl.Log.Info("db num err")
	} else {
		cl.Db = s.Db[id]
	}
}
