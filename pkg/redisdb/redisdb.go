package redisdb

import (
	"redis/pkg/ds/robj"
)

type RedisDb struct {
	Id     int
	AvgTtl int
	Dict   map[*robj.RedisObject]*robj.RedisObject
}

func NewRedisDb(id int) *RedisDb {
	rdb := new(RedisDb)
	rdb.Dict = make(map[*robj.RedisObject]*robj.RedisObject, 0)
	rdb.Id = id
	return rdb
}
