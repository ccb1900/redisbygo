package shared

import (
	"github.com/ccb1900/redisbygo/pkg"
	"sync"
)

type Shared struct {
	CRLF      *pkg.RedisObject
	Ok        *pkg.RedisObject
	EmptyBulk *pkg.RedisObject
	CZero     *pkg.RedisObject
	Cone      *pkg.RedisObject
	Pong      *pkg.RedisObject
}

var once sync.Once
var shared *Shared

func NewShared() *Shared {
	once.Do(func() {
		shared = new(Shared)
		s := "PONG"
		shared.Pong = pkg.NewRedisObject(pkg.ObjString, &s)
	})
	return shared
}
