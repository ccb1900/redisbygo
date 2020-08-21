package client

import (
	"github.com/ccb1900/redisbygo/pkg/ds/sds"
	"github.com/ccb1900/redisbygo/pkg/log"
	"github.com/ccb1900/redisbygo/pkg/redisdb"
	"net"
)

type Client struct {
	Conn     net.Conn
	Index    int
	Log      *log.Log
	Db       *redisdb.RedisDb
	Querybuf *sds.Sds
	Query    string
}
