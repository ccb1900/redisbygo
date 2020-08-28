package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
*** 写入aof文件
**/
type Aof struct {
	fd *os.File
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

	a.fd = f
}

// 加载aof文件
func (a *Aof) Load() {
	contents, err := ioutil.ReadFile("appendonly.aof")

	if err != nil {
		panic("load file failed..")
	}

	fmt.Println(string(contents))
}

// 写入
func (a *Aof) Write(content string) {
	bf := bufio.NewWriter(a.fd)
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
	err := a.fd.Close()
	if err != nil {
		fmt.Println("aof close err::", err)
	}
}

func NewAof() *Aof {
	a := new(Aof)
	a.Create()
	return a
}
