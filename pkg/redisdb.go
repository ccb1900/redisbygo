package pkg

type RedisDb struct {
	Id     int
	AvgTtl int
	Dict   map[*RedisObject]*RedisObject
}

func NewRedisDb(id int) *RedisDb {
	rdb := new(RedisDb)
	rdb.Dict = make(map[*RedisObject]*RedisObject, 0)
	rdb.Id = id
	return rdb
}
func Add(rb *RedisDb, key *RedisObject, value *RedisObject) {
	rb.Dict[key] = value
}
