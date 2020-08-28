package list

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/shared"
)

const Head = 0
const Tail = 1

func RPushCommand(cl *pkg.Client) {
	pushGenericCommand(cl, Tail)
}

func LPushCommand(cl *pkg.Client) {
	pushGenericCommand(cl, Head)
}
func pushGenericCommand(cl *pkg.Client, where int) {
	value := cl.Db.LookupKeyWrite(cl.Argv[1])
	sh := shared.NewShared()
	if value != nil && value.TypeFlag != pkg.ObjList {
		cl.AddReplyRedisObject(sh.WrongTypeErr)
		return
	}

	for i := 2; i < len(cl.Argv); i++ {
		if value == nil {
			value = pkg.CreateQuickListObject()
			cl.Db.Dict.Add(cl.Argv[1], value)
		}

		listTypePush(value, cl.Argv[i], where)
	}

	cl.AddReply("success")
}

func listTypePush(value *pkg.RedisObject, element *pkg.RedisObject, where int) {
	if value.Encoding == pkg.ObjEncodingQuickList {

	}
}
