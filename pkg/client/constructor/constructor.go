package constructor

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/client/db"
	"github.com/ccb1900/redisbygo/pkg/log"
	"net"
)

func NewClient(conn net.Conn) *pkg.Client {
	c := new(pkg.Client)
	c.Log = log.NewLog()
	db.SelectDb(c, 0)
	c.Conn = conn
	c.QueryBuf = make([]byte, 0)
	c.BulkLen = -1
	return c
}
