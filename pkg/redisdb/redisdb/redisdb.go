package redisdb

import (
	"redis/pkg/ds/robj"
	"redis/pkg/redisdb"
)

func Add(rb *redisdb.RedisDb, key *robj.RedisObject, value *robj.RedisObject) {
	rb.Dict[key] = value
}
