package pkg

import (
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/config"
	"github.com/ccb1900/redisbygo/pkg/log"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	AofOff = 0
	AofOn  = 1
)

type Server struct {
	RequirePass          bool
	Log                  *log.Log
	Clients              map[int]*Client
	Db                   []*RedisDb
	No                   int
	StatRejectedConn     int
	CurrentClient        chan *Client
	Listener             net.Listener
	Aof                  *Aof
	WaitCloseClients     chan int
	NewClients           chan net.Conn
	Commands             map[string]*RedisCommand
	ClientMaxQueryBufLen int
	AofState             int
	AofFsync             int
	AofSelectedDb        int
	ReplicaList          []*ReplicationReplica
	Main                 *ReplicationMain
	PubSubChannels       map[string][]*Client
	Monitors             []*Client
}

var gs *Server
var onceServer sync.Once

func NewServer() *Server {
	if gs == nil {
		onceServer.Do(func() {
			gs = new(Server)
			gs.Log = log.NewLog()
			gs.Clients = make(map[int]*Client, 0)
			gs.CurrentClient = make(chan *Client, 2048)
			gs.Aof = NewAof()
			gs.RequirePass = false
			gs.AofState = AofOff

			gs.Main = NewReplicationMain()
			gs.Monitors = make([]*Client, 0)
			c := config.NewConfig()
			dbList := make([]*RedisDb, c.Dbnum)

			for i := 0; i < len(dbList); i++ {
				dbList[i] = NewRedisDb(i)
			}

			gs.Db = dbList
			gs.PubSubChannels = make(map[string][]*Client, 0)
			gs.WaitCloseClients = make(chan int, 16)
			gs.NewClients = make(chan net.Conn, 32)
			gs.ClientMaxQueryBufLen = ProtoMaxQueryBufLen
		})
	}

	return gs
}

func ServerCron() {
	ticker := time.NewTicker(1 * time.Second)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("2 s", t)
			}
		}
	}()

	wg.Wait()
}

func (s *Server) LoadDataFromDisk() {
	s.Aof.LoadAppendOnlyFile("")
}

func LookupCommand(cmd string) (*RedisCommand, bool) {
	s := NewServer()
	v, ok := s.Commands[strings.ToLower(cmd)]
	return v, ok
}
func GetCommandMessage(ss []*RedisObject) []string {
	s := make([]string, 0)
	for i := 0; i < len(ss); i++ {
		s = append(s, *ss[i].Ptr.(*string))
	}

	return s
}
