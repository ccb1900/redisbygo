package pkg

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	pending := make(chan int, 1)

	<-pending

	fmt.Println(pending)
}

func TestServerCron(t *testing.T) {
	ServerCron()
}

func TestLookupCommand(t *testing.T) {
	var s map[string]string
	s = make(map[string]string)
	s["a"] = "b"
	v1, b := s["b"]
	v2, y := s["a"]
	fmt.Println(v1, b, v2, y)
}
