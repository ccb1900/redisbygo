package server

import (
	"bufio"
	"bytes"
	"fmt"
	constructor2 "github.com/ccb1900/redisbygo/app/server/constructor"
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/client"
	"github.com/ccb1900/redisbygo/pkg/client/constructor"
	"github.com/ccb1900/redisbygo/pkg/command/command"
	"github.com/ccb1900/redisbygo/pkg/command/table"
	config2 "github.com/ccb1900/redisbygo/pkg/config"
	"github.com/ccb1900/redisbygo/pkg/ds/robj"
	"github.com/ccb1900/redisbygo/pkg/others"
	"github.com/ccb1900/redisbygo/pkg/redisdb/redisdb"
	"github.com/ccb1900/redisbygo/pkg/utils"
	"io"
	"net"
	"strconv"
)

func InitServerConfig(s *constructor2.Server) {
	s.Log.Info("oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo")
	l := len(table.RedisCommandTable)
	s.Commands = make(map[string]command.RedisCommand, l)
	for i := 0; i < l; i++ {
		s.Commands[table.RedisCommandTable[i].Name] = table.RedisCommandTable[i]
	}
}

func CreateServer() {
	s := constructor2.NewServer()
	c := config2.NewConfig()
	InitServerConfig(s)

	s.Log.Info("server start")
	s.Log.Info("serve on :" + strconv.Itoa(c.Port))

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(c.Port))

	s.Listener = ln
	if err != nil {
		s.Log.Info("listen::" + err.Error())
	} else {
		go acceptRequest(s)
		go handleCommands(s)
	}
}

// 处理命令
func handleCommands(s *constructor2.Server) {
	for {
		//go func() {
		//	fmt.Println("waiting commands....")
		//}()
		select {
		case cm := <-s.CurrentClient:
			//fmt.Println("handleCommands", command.Query)
			// 解析命令
			parseCommand(cm)
			// 写入aof
			//s.Aof.Write(command.Query)
		}
	}
}

// 解析命令
func parseCommand(c client.Client) {
	key := robj.NewRedisObject()
	redisdb.Add(c.Db, key, key)
	// 回复
	go response(c.Conn, "OK")

	fmt.Println(c.Argv)
}

// 处理单行命令
func ProcessInlineBuffer(c *client.Client) {
	// 读取内容
	// 分离字符串
	// 创建roj对象
}

// 回复客户端
func response(conn net.Conn, message string) {
	writer := bufio.NewWriter(conn)
	_, _ = writer.WriteString(utils.ProtocolLine(message))
	_ = writer.Flush()
}

// 接受客户端请求
func acceptRequest(s *constructor2.Server) {
	go func() {
		for {
			select {
			case index := <-s.WaitCloseClients:
				//fmt.Println("acceptRequest trigger delete1")
				delete(s.Clients, index)
				//fmt.Println("acceptRequest trigger delete2")
			case conn := <-s.NewClients:
				//fmt.Println("new client is coming...")
				s.No = s.No + 1
				newClient := constructor.NewClient(conn)
				newClient.Index = s.No
				newClient.Db = s.Db[0]
				cc := config2.NewConfig()

				if len(s.Clients) >= cc.Maxclients {
					w := bufio.NewWriter(newClient.Conn)
					_, _ = w.WriteString(utils.ProtocolLineErr("ERR max number of clients reached"))
					s.StatRejectedConn++
					_ = w.Flush()

					fmt.Println("client up to max")
				} else {
					s.Clients[s.No] = newClient
					//go func() {
					//	fmt.Println("accept client::")
					//}()

					go handleConnection(s, newClient)
				}
			}
		}
	}()
	for {
		go func() {
			//s.Log.Info("waiting connecting...")
		}()
		conn, err := s.Listener.Accept()
		//s.Log.Info("waiting connecting2...")
		if err != nil {
			fmt.Println("client reach but fail::", err)
			s.Log.Info(err.Error())
		} else {
			go func() {
				s.NewClients <- conn
			}()
		}
	}
}

// 处理客户端连接
func handleConnection(s *constructor2.Server, cl *client.Client) {
	//s.Log.Info("new client")
	s.Log.Info(cl.Conn.RemoteAddr().String())

	for {
		// 一次读取16kb
		readLen := others.ProtoIoBufLen
		buf := make([]byte, readLen)
		//qbLen := len(cl.QueryBuf)
		fmt.Println("querybufs", string(cl.QueryBuf))
		size, err := cl.Conn.Read(buf)
		fmt.Println("size::", size, "err::", err)
		if size == 0 && err == io.EOF {
			// 客户端关闭
			err = cl.Conn.Close()
			if err != nil {
				//fmt.Println("close client fail::", err)
			} else {
				//fmt.Println("close client success::")
			}
			// 删除客户端
			s.WaitCloseClients <- cl.Index
			//fmt.Println("handleConnection trigger delete::", cl.Index)
			// 结束循环，回收协程
		} else {
			cl.QueryBuf = append(cl.QueryBuf, buf[:size]...)
			if cl.MultiBulkLen == 0 {
				pos := bytes.Index(cl.QueryBuf, []byte{'\r', '\n'})
				if pos > 0 && cl.QueryBuf[0] == '*' {
					cl.MultiBulkLen = pkg.S2Int(string(cl.QueryBuf[1:pos]))
					cl.QueryBuf = cl.QueryBuf[pos+2:]
				} else {
					continue
				}

			}

			for cl.MultiBulkLen > 0 {
				fmt.Println("realbuf::", string(cl.QueryBuf))
				pos := bytes.Index(cl.QueryBuf[0:], []byte{'\r', '\n'})

				if pos == -1 {
					break
				}
				if cl.QueryBuf[0] == '$' {
					cl.BulkLen = pkg.S2Int(string(cl.QueryBuf[1:pos]))
				} else {
					cl.Argv = append(cl.Argv, string(cl.QueryBuf[0:pos]))
					cl.MultiBulkLen--
				}

				cl.QueryBuf = cl.QueryBuf[pos+2:]
			}

			if cl.MultiBulkLen == 0 {
				s.CurrentClient <- *cl
				cl.MultiBulkLen = 0
				cl.Argv = make([]string, 0)
				cl.BulkLen = -1
				cl.QueryBuf = make([]byte, 0)
			}
		}
	}
}
