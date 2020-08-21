package db

import (
	"github.com/ccb1900/redisbygo/app/server/constructor"
	"github.com/ccb1900/redisbygo/pkg/client"
)

func SelectDbCommand(c *client.Client) {
	s := constructor.NewServer()

	c.Db = s.Db[1]
}
