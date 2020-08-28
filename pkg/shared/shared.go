package shared

import (
	"github.com/ccb1900/redisbygo/pkg"
	"sync"
)

type Shared struct {
	CRLF         *pkg.RedisObject
	Ok           *pkg.RedisObject
	EmptyBulk    *pkg.RedisObject
	CZero        *pkg.RedisObject
	Cone         *pkg.RedisObject
	Pong         *pkg.RedisObject
	NullBulk     *pkg.RedisObject
	WrongTypeErr *pkg.RedisObject
}

var once sync.Once
var shared *Shared

func NewShared() *Shared {
	once.Do(func() {
		shared = new(Shared)
		s := "PONG"
		nullBulk := "$-1\\r\\n"
		ok := "OK"
		shared.Pong = pkg.NewRedisObject(pkg.ObjString, &s)
		shared.Ok = pkg.NewRedisObject(pkg.ObjString, &ok)
		shared.NullBulk = pkg.NewRedisObject(pkg.ObjString, &nullBulk)
	})
	return shared
}
