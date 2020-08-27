package pkg

type RedisObject struct {
	Encoding int
	TypeFlag int
	Lru      int
	Refcount int
	Ptr      interface{}
}

func NewRedisObject(t int, ptr interface{}) *RedisObject {
	ro := new(RedisObject)
	ro.Encoding = ObjEncodingRaw
	ro.TypeFlag = t
	ro.Ptr = ptr
	ro.Refcount = 1
	return ro
}
