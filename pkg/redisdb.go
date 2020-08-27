package pkg

type RedisDb struct {
	Id     int
	AvgTtl int
	Dict   *Dict
}

func NewRedisDb(id int) *RedisDb {
	rdb := new(RedisDb)
	rdb.Dict = NewDict()
	rdb.Id = id
	return rdb
}
func (rb *RedisDb) Add(key *RedisObject, value *RedisObject) {
	rb.Dict.Add(key, value)
}

func (rb *RedisDb) LookupKey(key *RedisObject, flags int) *RedisObject {
	return rb.Dict.Get(key)
}

func (rb *RedisDb) LookupKeyReadWithFlags(key *RedisObject, flags int) *RedisObject {
	var val *RedisObject

	val = rb.LookupKey(key, flags)

	return val
}

func (rb *RedisDb) LookupKeyRead(key *RedisObject) *RedisObject {
	return rb.LookupKeyReadWithFlags(key, LookupNone)
}
