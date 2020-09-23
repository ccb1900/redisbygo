package set

import "github.com/ccb1900/redisbygo/pkg"

func SAddCommand(cl *pkg.Client) {
	set := cl.Db.LookupKeyWrite(cl.Argv[1])

	if set != nil {

	}
	cl.Db.Dict.Add(cl.Argv[1], set)
}
