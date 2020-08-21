package db

import (
	"redis/app/server/constructor"
	"redis/pkg/client"
)

func SelectDbCommand(c *client.Client) {
	s := constructor.NewServer()

	c.Db = s.Db[1]
}
