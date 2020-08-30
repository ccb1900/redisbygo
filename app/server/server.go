package server

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/command/table"
	"github.com/ccb1900/redisbygo/pkg/config"
	"io"
	"net"
	"strconv"
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
	s.Aof.LoadAppendOnlyFile("appendonly.aof")
	s.Log.Info("server start")
	s.Log.Info("serve on :" + strconv.Itoa(c.Port))

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(c.Port))
	fmt.Println("waiting redis client....on 6378")
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
			cl.ParseCommand()
			//s.Aof.Write(command.Query)
			cl.Free()
			cl.Pending <- new(pkg.Pending)
		}
	}
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
		//c := 0
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
				//default:
				//	c++
				//	fmt.Println("beforeasleep..." + strconv.Itoa(c))
			}
		}
	}()
	for {
		conn, err := s.Listener.Accept()
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
	s.Log.Info(cl.Conn.RemoteAddr().String())
	cl.Pending <- new(pkg.Pending)
	for {
		<-cl.Pending
		readLen := pkg.ProtoIoBufLen
		buf := make([]byte, readLen)
		size, err := cl.Conn.Read(buf)
		if size == 0 && err == io.EOF {
			err = cl.Conn.Close()
			if err != nil {
				s.Log.Error(err.Error())
			} else {
				s.Log.Info("close success")
			}
			// 删除客户端
			s.WaitCloseClients <- cl.Index
			break
		} else {
			cl.QueryBuf = append(cl.QueryBuf, buf[:size]...)
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
