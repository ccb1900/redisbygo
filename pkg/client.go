package pkg

import (
	"bufio"
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/config"
	"github.com/ccb1900/redisbygo/pkg/log"
	"net"
	"strings"
)

type Client struct {
	Conn            net.Conn
	Index           int
	Log             *log.Log
	Db              *RedisDb
	QueryBuf        []byte
	Query           string
	Argv            []*RedisObject
	LastInteraction string
	ReqType         int // 协议类型
	MultiBulkLen    int
	QbPos           int
	BulkLen         int
	Cmd             *RedisCommand
	LastCmd         *RedisCommand
	Pending         chan *Pending
}
type Pending struct {
}

func NewClient(conn net.Conn) *Client {
	c := new(Client)
	c.Log = log.NewLog()
	c.SelectDb(0)
	c.Conn = conn
	c.QueryBuf = make([]byte, 0)
	c.Pending = make(chan *Pending, 1)
	c.BulkLen = -1
	return c
}
func (cl *Client) Free() {
	cl.MultiBulkLen = 0
	cl.Argv = make([]*RedisObject, 0)
	cl.BulkLen = -1
	cl.Cmd = nil
	cl.LastCmd = nil
	cl.QueryBuf = make([]byte, 0)
}
func (cl *Client) reply(message string) {
	bf := bufio.NewWriter(cl.Conn)
	_, _ = bf.WriteString(message)
	_ = bf.Flush()
}
func (cl *Client) AddReplyRedisObject(object *RedisObject) {
	cl.AddReply(*object.Ptr.(*string))
}
func (cl *Client) AddReply(message string) {
	cl.reply(ProtocolLine(message))
}

func (cl *Client) AddReplyErrorFormat(args []string, messages ...string) {
	s := ""

	for i := 0; i < len(args); i++ {
		message := ""
		if i > 1 {
			message = messages[1]
		} else {
			message = messages[0]
		}
		s += fmt.Sprintf(message, args[i])
	}

	cl.reply(ProtocolLineErr(s))
}

func (cl *Client) AddReplyError(message string) {
	cl.reply(ProtocolLineErr(message))
}
func (cl *Client) AddReplyBulk(object *RedisObject) {
	cl.AddReplyRedisObject(object)
}
func (cl *Client) AddReplyBulkLen(object *RedisObject) {

}

func (cl *Client) AddReplyHelp(help []string) {
	cl.AddReply(strings.Join(help, ","))
}

func (cl *Client) AddReplySubcommandSyntaxError() {
	cl.AddReplyErrorFormat([]string{
		*cl.Argv[1].Ptr.(*string),
		cl.Cmd.Name,
	}, "Unknown subcommand or wrong number of arguments for '%s'.", "Try %s HELP.")
}

func (cl *Client) LookupKeyReadOrReply(key *RedisObject, reply *RedisObject) *RedisObject {
	o := cl.Db.LookupKeyRead(key)

	if o == nil {
		cl.AddReplyRedisObject(reply)
	}
	return o
}

func (cl *Client) SelectDb(id int) int {
	c := config.NewConfig()
	s := NewServer()
	if id < 0 || id >= c.Dbnum {
		cl.Log.Info("db num err")
		return C_ERR
	} else {
		cl.Db = s.Db[id]
		return C_OK
	}
}
