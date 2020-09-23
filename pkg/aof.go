package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/log"
	"io"
	"os"
)

/**
*** 写入aof文件
**/
type Aof struct {
	Fd  *os.File
	Log log.ILog
}

// 创建aof文件
func (a *Aof) Create() {
	name := "appendonly.aof"
	var f *os.File
	var e error
	if _, err := os.Stat(name); os.IsNotExist(err) {
		f, e = os.Create(name)

		if e != nil {
			panic("create file failed..")
		}
	} else {
		f, e = os.OpenFile(name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if e != nil {
			panic("create file failed..")
		}
	}
	a.Log = log.NewLog(*a)
	a.Fd = f
}
func StartLoading(file *os.File) {

}

// 加载aof文件
func (a *Aof) LoadAppendOnlyFile(filename string) int {
	s := NewServer()
	cl := CreateFakeClient()

	s.AofState = AofOn

	f, err := os.Open(filename)

	if err != nil {
		a.Log.Error(err.Error())

		return CErr
	}

	i, done := a.LoadFile(f, cl)
	if done {
		return i
	}

	fmt.Println("test")
	return COk
}

func (a *Aof) LoadFile(f *os.File, cl *Client) (int, bool) {
	for {
		readLen := ProtoIoBufLen
		buf := make([]byte, readLen)
		size, err := f.Read(buf)

		if err != nil {
			a.Log.Error(err.Error())
			return CErr, true
		}

		if err == io.EOF {
			break
		}
		cl.QueryBuf = append(cl.QueryBuf, buf[:size]...)

		a.Load(cl)
	}
	return 0, false
}

func (a *Aof) Load(cl *Client) {
	for {
		if a.ParseStart(cl) {
			break
		}

		if !a.ParseMultiBulkLen(cl) {
			break
		}

		if cl.MultiBulkLen == 0 {
			if !cl.ParseCommand() {
				a.Log.Error("parse error")
			}
			cl.FreeFakeClient()
			continue
		}
	}
}

func (a *Aof) ParseStart(cl *Client) bool {
	if cl.MultiBulkLen == 0 {
		pos := bytes.Index(cl.QueryBuf, []byte{'\r', '\n'})
		if pos > 0 && cl.QueryBuf[0] == '*' {
			cl.MultiBulkLen = S2Int(string(cl.QueryBuf[1:pos]))
			cl.QueryBuf = cl.QueryBuf[pos+2:]
		} else {
			if pos > 0 && cl.QueryBuf[0] != '*' {
				ptr := string(cl.QueryBuf[0:pos])
				cl.Argv = append(cl.Argv, NewRedisObject(ObjString, &ptr))
				return false
			} else {
				return true
			}
		}
	}

	return false
}

func (a *Aof) ParseMultiBulkLen(cl *Client) bool {
	for cl.MultiBulkLen > 0 {
		pos := bytes.Index(cl.QueryBuf[0:], []byte{'\r', '\n'})

		if pos == -1 {
			return false
		}
		if cl.QueryBuf[0] == '$' {
			cl.BulkLen = S2Int(string(cl.QueryBuf[1:pos]))
		} else {
			ptr := string(cl.QueryBuf[0:pos])
			cl.Argv = append(cl.Argv, NewRedisObject(ObjString, &ptr))
			cl.MultiBulkLen--
		}

		cl.QueryBuf = cl.QueryBuf[pos+2:]
	}

	return true
}

// 写入
func (a *Aof) Write(content string) {
	bf := bufio.NewWriter(a.Fd)
	_, e := bf.WriteString(content)

	if e != nil {
		panic(e)
	}

	e = bf.Flush()

	if e != nil {
		panic(e)
	}
}

// 关闭
func (a *Aof) Close() {
	err := a.Fd.Close()
	if err != nil {
		fmt.Println("aof close err::", err)
	}
}

func NewAof() *Aof {
	a := new(Aof)
	a.Create()
	return a
}

func CatAppendOnlyGenericCommand(buf string, argv []*RedisObject) string {
	return ""
}
