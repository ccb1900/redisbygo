package pkg

import (
	"bufio"
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/log"
	"net"
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

func (cl *Client) AddReplyError(messages string) {

}
func (cl *Client) AddReplyBulk(object *RedisObject) {
	cl.AddReplyRedisObject(object)
}
func (cl *Client) AddReplyBulkLen(object *RedisObject) {

}
