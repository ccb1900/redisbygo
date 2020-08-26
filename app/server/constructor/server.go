package constructor

import (
	"github.com/ccb1900/redisbygo/pkg/client"
	"github.com/ccb1900/redisbygo/pkg/command/command"
	"github.com/ccb1900/redisbygo/pkg/config"
	"github.com/ccb1900/redisbygo/pkg/log"
	"github.com/ccb1900/redisbygo/pkg/others"
	"github.com/ccb1900/redisbygo/pkg/persist/aof"
	"github.com/ccb1900/redisbygo/pkg/redisdb"
	"net"
	"sync"
)

type Server struct {
	Log                  *log.Log
	Clients              map[int]*client.Client
	Db                   []*redisdb.RedisDb
	No                   int
	StatRejectedConn     int
	CurrentClient        chan client.Client
	Listener             net.Listener
	Aof                  *aof.Aof
	WaitCloseClients     chan int
	NewClients           chan net.Conn
	Commands             map[string]command.RedisCommand
	ClientMaxQueryBufLen int
}

var gs *Server
var once sync.Once

func NewServer() *Server {
	if gs == nil {
		once.Do(func() {
			gs = new(Server)
			gs.Log = log.NewLog()
			gs.Clients = make(map[int]*client.Client, 0)
			gs.CurrentClient = make(chan client.Client, 2048)
			gs.Aof = aof.New()

			c := config.NewConfig()
			dbList := make([]*redisdb.RedisDb, c.Dbnum)

			for i := 0; i < len(dbList); i++ {
				dbList[i] = redisdb.NewRedisDb(i)
			}

			gs.Db = dbList
			gs.WaitCloseClients = make(chan int, 16)
			gs.NewClients = make(chan net.Conn, 32)
			gs.ClientMaxQueryBufLen = others.ProtoMaxQueryBufLen
		})
	}

	return gs
}
