package connection

import (
	"github.com/ccb1900/redisbygo/pkg"
	"strconv"
)

func SelectDbCommand(c *pkg.Client) {
	id, err := strconv.Atoi(*c.Argv[1].Ptr.(*string))

	if err != nil {
		c.Log.Error(err.Error())
		c.AddReplyError(err.Error())
	} else {
		if c.SelectDb(id) == pkg.COk {
			c.AddReply("success")
		} else {
			c.AddReplyError("it is too big")
		}
	}
}
