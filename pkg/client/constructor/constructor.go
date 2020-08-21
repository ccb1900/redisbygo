package constructor

import (
	"net"
	"redis/pkg/client"
	"redis/pkg/client/db"
	"redis/pkg/ds/sds/methods"
	"redis/pkg/log"
)

func NewClient(conn net.Conn) *client.Client {
	c := new(client.Client)
	c.Log = log.NewLog()
	db.SelectDb(c, 0)
	c.Conn = conn
	c.Querybuf = methods.Sdsempty()
	return c
}
