package pkg

import (
	"github.com/ccb1900/redisbygo/pkg/ds/intset"
	"github.com/ccb1900/redisbygo/pkg/ds/quicklist"
)

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

func NewStringRedisObject(s string) *RedisObject {
	return NewRedisObject(ObjString, &s)
}

func (r *RedisObject) MakeShared() *RedisObject {
	r.Refcount = OBJ_SHARED_REFCOUNT
	return r
}

func CreateQuickListObject() *RedisObject {
	ql := quicklist.CreateQuickList()
	obj := NewRedisObject(ObjList, ql)

	obj.Encoding = ObjEncodingQuickList
	return obj
}

func SetTypeCreate(s string) *RedisObject {
	return nil
}

func IsSdsRepresentableAsLongLong(s string) int {
	return 0
}

func CreateIntSetObject() *RedisObject {
	is := intset.NewIntSet()
	o := NewRedisObject(ObjSet, is)
	o.Encoding = ObjEncodingIntSet
	return o
}

func CreateSetObject() *RedisObject {
	return NewRedisObject(1, 1)
}

func CreateHashRedisObject() *RedisObject {
	return NewRedisObject(1, 1)
}
func (r *RedisObject) DecrRefCount() {

}
