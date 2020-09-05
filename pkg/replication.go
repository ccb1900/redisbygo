package pkg

import (
	"fmt"
	"net"
	"strconv"
)

const ClientSlave = 1 << 0

const REPL_STATE_NONE = 0
const REPL_STATE_CONNECTING = 2
const REPL_STATE_TRANSFER = 14
const REPL_STATE_CONNECT = 1
const REPL_STATE_CONNECTED = 15

// Main server info
type ReplicationMain struct {
	Host   string
	Port   int
	Client *Client
	Auth   string
	State  int
}

type ReplicationReplica struct {
	Host  string
	Port  int
	State int
}

func NewReplicationMain(conn net.Conn) *ReplicationMain {
	m := new(ReplicationMain)
	m.Client = NewClient(conn)
	m.State = REPL_STATE_NONE
	return m
}
func replicationFeedSlaves() {

}

func (m *ReplicationMain) ReplicationUnsetMaster() {

}
func ConnectWithMaster() int {
	s := NewServer()
	c, err := net.Dial(s.Main.Host, strconv.Itoa(s.Main.Port))

	if err != nil {
		s.Log.Error(err.Error())
		return CErr
	}
	go func() {

		for {
			if s.Main.State == REPL_STATE_NONE {
				_ = s.Main.Client.Conn.Close()
				return
			}
			buf := make([]byte, 1024)
			size, err := c.Read(buf)
			if err != nil {
				s.Log.Error(err.Error())
			} else {
				realBuf := buf[0:size]

				fmt.Println(realBuf)
			}
		}
	}()
	s.Main.State = REPL_STATE_CONNECTING

	return COk
}
func ReplicationCron() {
	// 连接主服务器,长连接
	// 心跳连接
	// 发送同步命令
	// 部分同步
	// 全量同步
	s := NewServer()
	if s.Main.State == REPL_STATE_CONNECT {
		s.Log.Notice(fmt.Sprintf("Connecting to MASTER %s:%d",
			s.Main.Host, s.Main.Port))
		if ConnectWithMaster() == COk {
			s.Log.Notice("MASTER <-> REPLICA sync started")
		}
	}

}

func SyncWithMaster() {

}
