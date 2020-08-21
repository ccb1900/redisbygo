package constructor

import (
	"github.com/ccb1900/redisbygo/pkg/client"
	"github.com/ccb1900/redisbygo/pkg/config"
	"github.com/ccb1900/redisbygo/pkg/log"
	"github.com/ccb1900/redisbygo/pkg/persist/aof"
	"github.com/ccb1900/redisbygo/pkg/redisdb"
	"net"
	"sync"
)

type Server struct {
	Log              *log.Log
	Clients          map[int]*client.Client
	Db               []*redisdb.RedisDb
	No               int
	StatRejectedConn int
	Commands         chan *client.Client
	Listener         net.Listener
	Aof              *aof.Aof
}

var gs *Server
var once sync.Once

func NewServer() *Server {
	once.Do(func() {
		gs = new(Server)
		gs.Log = log.NewLog()
		gs.Clients = make(map[int]*client.Client, 0)
		gs.Commands = make(chan *client.Client, 2048)
		gs.Aof = aof.New()

		c := config.NewConfig()
		dbList := make([]*redisdb.RedisDb, c.Dbnum)

		for i := 0; i < len(dbList); i++ {
			dbList[i] = redisdb.NewRedisDb(i)
		}

		gs.Db = dbList
	})

	return gs
}
