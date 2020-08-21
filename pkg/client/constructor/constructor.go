package constructor

import (
	"github.com/ccb1900/redisbygo/pkg/client"
	"github.com/ccb1900/redisbygo/pkg/client/db"
	"github.com/ccb1900/redisbygo/pkg/ds/sds/methods"
	"github.com/ccb1900/redisbygo/pkg/log"
	"net"
)

func NewClient(conn net.Conn) *client.Client {
	c := new(client.Client)
	c.Log = log.NewLog()
	db.SelectDb(c, 0)
	c.Conn = conn
	c.Querybuf = methods.Sdsempty()
	return c
}