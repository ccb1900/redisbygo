package pkg

import (
	"fmt"
	"net"
	"strconv"
)

func ReplicationCron() int {
	s := NewServer()

	c, err := net.Dial(s.Main.Host, strconv.Itoa(s.Main.Port))

	if err != nil {
		s.Log.Error(err.Error())
		return C_ERR
	}

	for {
		buf := make([]byte, 1024)

		size, err := c.Read(buf)

		if err != nil {
			s.Log.Error(err.Error())
			return C_ERR
		}

		realBuf := buf[0:size]

		fmt.Println(realBuf)
	}
}

func SyncWithMaster() {

}
