package pkg

import "strings"

type Dict struct {
	Storage map[string]*RedisObject
}

func (d *Dict) Add(key *RedisObject, value *RedisObject) {
	s := strings.ToLower(*key.Ptr.(*string))
	d.Storage[s] = value
}

func (d *Dict) Get(key *RedisObject) *RedisObject {
	value := d.Storage[strings.ToLower(*key.Ptr.(*string))]
	return value
}

func NewDict() *Dict {
	d := new(Dict)
	d.Storage = make(map[string]*RedisObject, 0)
	return d
}
