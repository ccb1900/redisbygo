package db

import (
	"github.com/ccb1900/redisbygo/pkg"
)

func SelectDbCommand(c *pkg.Client) {
	s := pkg.NewServer()

	c.Db = s.Db[1]
}
