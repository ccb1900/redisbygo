package pkg

import "github.com/ccb1900/redisbygo/pkg/ds/quicklist"

type RedisObject struct {
	Encoding int
	TypeFlag int
	Lru      int
	Refcount int
	Ptr      interface{}
}

func NewRedisObject(typeFlag int, ptr interface{}) *RedisObject {
	ro := new(RedisObject)
	ro.Encoding = ObjEncodingRaw
	ro.TypeFlag = typeFlag
	ro.Ptr = ptr
	ro.Refcount = 1
	return ro
}

func CreateQuickListObject() *RedisObject {
	ql := quicklist.CreateQuickList()
	obj := NewRedisObject(ObjList, ql)

	obj.Encoding = ObjEncodingQuickList
	return obj
}

func (o *RedisObject) DecrRefCount() {

}
