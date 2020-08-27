package db

import (
	"github.com/ccb1900/redisbygo/pkg"
)

func SelectDb(cl *pkg.Client, id int) {
	c := pkg.NewConfig()
	s := pkg.NewServer()
	if id < 0 || id >= c.Dbnum {
		cl.Log.Info("db num err")
	} else {
		cl.Db = s.Db[id]
	}
}
