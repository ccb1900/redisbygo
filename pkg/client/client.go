package client

import (
	"net"
	"redis/pkg/ds/sds"
	"redis/pkg/log"
	"redis/pkg/redisdb"
)

type Client struct {
	Conn     net.Conn
	Index    int
	Log      *log.Log
	Db       *redisdb.RedisDb
	Querybuf *sds.Sds
	Query    string
}
