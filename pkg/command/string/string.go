package string

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/shared"
)

func GetCommand(c *pkg.Client) {
	getGenericCommand(c)
}
func getGenericCommand(cl *pkg.Client) {
	sh := shared.NewShared()
	o := cl.LookupKeyReadOrReply(cl.Argv[1], sh.NullBulk)
	if o == nil {
		return
	} else {
		if o.TypeFlag != pkg.ObjString {
			cl.AddReplyRedisObject(shared.NewShared().WrongTypeErr)
		} else {
			cl.AddReplyBulk(o)
		}
	}
}
func SetCommand(c *pkg.Client) {
	setGenericCommand(c, 0, c.Argv[1], c.Argv[2], nil, 1, nil, nil)
}

func setGenericCommand(cl *pkg.Client,
	flags int,
	key *pkg.RedisObject,
	val *pkg.RedisObject,
	expire *pkg.RedisObject, unit int, okReply *pkg.RedisObject, abortReply *pkg.RedisObject) {
	cl.Db.Add(key, val)

	sh := shared.NewShared()

	if okReply == nil {
		cl.AddReplyRedisObject(sh.Ok)
	} else {
		cl.AddReplyRedisObject(okReply)
	}

}

func IncrCommand(cl *pkg.Client) {
	incrDecrCommand(cl, 1)
}

func incrDecrCommand(cl *pkg.Client, incr int) {
	cl.AddReply("ok")
}
