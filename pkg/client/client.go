package client

import (
	"github.com/ccb1900/redisbygo/pkg/log"
	"github.com/ccb1900/redisbygo/pkg/redisdb"
	"net"
)

type Client struct {
	Conn            net.Conn
	Index           int
	Log             *log.Log
	Db              *redisdb.RedisDb
	QueryBuf        []byte
	Query           string
	Argv            []string
	LastInteraction string
	ReqType         int // 协议类型
	MultiBulkLen    int
	QbPos           int
	BulkLen         int
}
