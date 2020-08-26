package pkg

import "testing"

func TestCreate(t *testing.T) {
	f := New()

	f.Write("test\r\nddd\r\n")
	f.Write("test\r\nddd\r\n")
	f.Write("test\r\nddd\r\n")
	f.Write("test\r\nddd\r\n")

	f.Close()
}
