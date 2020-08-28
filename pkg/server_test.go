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
