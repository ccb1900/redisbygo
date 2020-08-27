package pkg

import (
	"github.com/ccb1900/redisbygo/pkg/log"
	"net"
	"sync"
)

type Server struct {
	RequirePass          bool
	Log                  *log.Log
	Clients              map[int]*Client
	Db                   []*RedisDb
	No                   int
	StatRejectedConn     int
	CurrentClient        chan Client
	Listener             net.Listener
	Aof                  *Aof
	WaitCloseClients     chan int
	NewClients           chan net.Conn
	Commands             map[string]*RedisCommand
	ClientMaxQueryBufLen int
}

var gs *Server
var onceServer sync.Once

func NewServer() *Server {
	if gs == nil {
		onceServer.Do(func() {
			gs = new(Server)
			gs.Log = log.NewLog()
			gs.Clients = make(map[int]*Client, 0)
			gs.CurrentClient = make(chan Client, 2048)
			gs.Aof = NewAof()
			gs.RequirePass = false

			c := NewConfig()
			dbList := make([]*RedisDb, c.Dbnum)

			for i := 0; i < len(dbList); i++ {
				dbList[i] = NewRedisDb(i)
			}

			gs.Db = dbList
			gs.WaitCloseClients = make(chan int, 16)
			gs.NewClients = make(chan net.Conn, 32)
			gs.ClientMaxQueryBufLen = ProtoMaxQueryBufLen
		})
	}

	return gs
}
