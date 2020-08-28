package server

import (
	"bufio"
	"bytes"
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/command/table"
	"github.com/ccb1900/redisbygo/pkg/config"
	"io"
	"net"
	"strconv"
	"strings"
)

func InitServerConfig(s *pkg.Server) {
	s.Log.Info("oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo")
	l := len(table.RedisCommandTable)
	s.Commands = make(map[string]*pkg.RedisCommand, l)
	for i := 0; i < l; i++ {
		s.Commands[table.RedisCommandTable[i].Name] = &table.RedisCommandTable[i]
	}
}

func CreateServer() {
	s := pkg.NewServer()
	c := config.NewConfig()
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
func handleCommands(s *pkg.Server) {
	for {
		select {
		case cl := <-s.CurrentClient:
			parseCommand(cl)
			//s.Aof.Write(command.Query)
			cl.Free()
			cl.Pending <- new(pkg.Pending)
		}
	}
}

func lookupCommand(cmd string) *pkg.RedisCommand {
	s := pkg.NewServer()
	return s.Commands[strings.ToLower(cmd)]
}
func getCommandMessage(ss []*pkg.RedisObject) []string {
	s := make([]string, 0)
	for i := 0; i < len(ss); i++ {
		s = append(s, *ss[i].Ptr.(*string))
	}

	return s
}

// 解析命令
func parseCommand(c *pkg.Client) bool {
	c.LastCmd = lookupCommand(*c.Argv[0].Ptr.(*string))
	c.Cmd = c.LastCmd
	errCommand := []string{"ERR unknown command `%s`,",
		"with args beginning with: `%s`,"}
	errArgument := "wrong number of arguments for '%s' command"
	if c.Cmd == nil {
		c.AddReplyErrorFormat(getCommandMessage(c.Argv), errCommand...)
		return true
	} else if (c.Cmd.Arity > 0 && c.Cmd.Arity != len(c.Argv)) || (len(c.Argv) < -c.Cmd.Arity) {
		c.AddReplyErrorFormat([]string{c.Cmd.Name}, errArgument)
		return true
	}
	// 检查是否需要密码

	// 检查集群

	// 处理最大内存

	// 不接受写命令的情况处理

	// 发布订阅的上下文

	// 执行命令

	//fmt.Println(c.Argv)
	c.Cmd.Proc(c)
	return true
}

// 处理单行命令
func ProcessInlineBuffer(c *pkg.Client) {
	// 读取内容
	// 分离字符串
	// 创建roj对象
}

// 回复客户端
func response(conn net.Conn, message string) {
	writer := bufio.NewWriter(conn)
	_, _ = writer.WriteString(pkg.ProtocolLine(message))
	_ = writer.Flush()
}

func errorResponse(conn net.Conn, message string) {
	writer := bufio.NewWriter(conn)
	_, _ = writer.WriteString(pkg.ProtocolLineErr(message))
	_ = writer.Flush()
}

// 接受客户端请求
func acceptRequest(s *pkg.Server) {
	go func() {
		for {
			select {
			case index := <-s.WaitCloseClients:
				delete(s.Clients, index)
			case conn := <-s.NewClients:
				s.No = s.No + 1
				newClient := pkg.NewClient(conn)
				newClient.Index = s.No
				newClient.Db = s.Db[0]
				cc := config.NewConfig()

				if len(s.Clients) >= cc.Maxclients {
					w := bufio.NewWriter(newClient.Conn)
					_, _ = w.WriteString(pkg.ProtocolLineErr("ERR max number of clients reached"))
					s.StatRejectedConn++
					_ = w.Flush()
				} else {
					s.Clients[s.No] = newClient
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
			s.Log.Info(err.Error())
		} else {
			go func() {
				s.NewClients <- conn
			}()
		}
	}
}

// 处理客户端连接
func handleConnection(s *pkg.Server, cl *pkg.Client) {
	//s.Log.Info("new client")
	s.Log.Info(cl.Conn.RemoteAddr().String())
	cl.Pending <- new(pkg.Pending)
	for {
		<-cl.Pending
		// 一次读取16kb
		readLen := pkg.ProtoIoBufLen
		buf := make([]byte, readLen)
		size, err := cl.Conn.Read(buf)
		if size == 0 && err == io.EOF {
			err = cl.Conn.Close()
			if err != nil {

			} else {

			}
			// 删除客户端
			s.WaitCloseClients <- cl.Index
			break
		} else {
			cl.QueryBuf = append(cl.QueryBuf, buf[:size]...)
			//fmt.Println("realbuf1::", string(cl.QueryBuf))
			if cl.MultiBulkLen == 0 {
				pos := bytes.Index(cl.QueryBuf, []byte{'\r', '\n'})
				if pos > 0 && cl.QueryBuf[0] == '*' {
					cl.MultiBulkLen = pkg.S2Int(string(cl.QueryBuf[1:pos]))
					cl.QueryBuf = cl.QueryBuf[pos+2:]
				} else {
					if pos > 0 && cl.QueryBuf[0] != '*' {
						ptr := string(cl.QueryBuf[0:pos])
						cl.Argv = append(cl.Argv, pkg.NewRedisObject(pkg.ObjString, &ptr))
					} else {
						cl.Pending <- new(pkg.Pending)
						continue
					}
				}

			}

			for cl.MultiBulkLen > 0 {
				//fmt.Println("realbuf::", string(cl.QueryBuf))
				pos := bytes.Index(cl.QueryBuf[0:], []byte{'\r', '\n'})

				if pos == -1 {
					break
				}
				if cl.QueryBuf[0] == '$' {
					cl.BulkLen = pkg.S2Int(string(cl.QueryBuf[1:pos]))
				} else {
					ptr := string(cl.QueryBuf[0:pos])
					cl.Argv = append(cl.Argv, pkg.NewRedisObject(pkg.ObjString, &ptr))
					cl.MultiBulkLen--
				}

				cl.QueryBuf = cl.QueryBuf[pos+2:]
			}

			if cl.MultiBulkLen == 0 {
				s.CurrentClient <- cl
			} else {
				cl.Pending <- new(pkg.Pending)
			}
		}

	}
}
