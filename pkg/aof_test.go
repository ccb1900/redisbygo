package pkg

import "testing"

func TestCreate(t *testing.T) {
	f := NewAof()

	f.Write("test\r\nddd\r\n")
	f.Write("test\r\nddd\r\n")
	f.Write("test\r\nddd\r\n")
	f.Write("test\r\nddd\r\n")

	f.Close()

	f.LoadAppendOnlyFile("appendonly.aof")
}
