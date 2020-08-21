package command

import "redis/pkg/client"

type RedisCommand struct {
	Name         string
	Proc         func(c *client.Client)
	Arity        int
	SFlags       string
	Flags        int
	GetKeysProc  *int
	FirstKey     int
	LastKey      int
	KeyStep      int
	Microseconds int
	Calls        int
}
