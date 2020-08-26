package pkg

import (
	"github.com/ccb1900/redisbygo/pkg/ds/robj"
	"github.com/ccb1900/redisbygo/pkg/redisdb"
)

func Add(rb *redisdb.RedisDb, key *robj.RedisObject, value *robj.RedisObject) {
	rb.Dict[key] = value
}
