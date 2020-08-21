package robj

type RedisObject struct {
	Encoding int
	TypeFlag int
	Lru      int
	Refcount int
	Ptr      *interface{}
}

func NewRedisObject() *RedisObject {
	ro := new(RedisObject)
	ro.Encoding = 1
	return ro
}
