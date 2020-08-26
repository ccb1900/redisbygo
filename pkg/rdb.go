package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func Test() {
	f, e := os.Create("dump.rdb")

	if e != nil {
		fmt.Println("create::", e)
	} else {
		var buffer bytes.Buffer
		ss := "REDIS0009RedisVersion5.0.9"
		s := []byte(ss)
		s = append(s, 0xfe)
		s = append(s, 0x00)
		s = append(s, 0x00)
		s = append(s, 0x0b)
		for i := 0; i < len(s); i++ {
			e := binary.Write(&buffer, binary.LittleEndian, s[i])
			fmt.Println("buffer::", buffer)
			if e != nil {
				fmt.Println("create::", e)
			}
		}

		size, e := f.Write(buffer.Bytes())
		fmt.Println(size, e)

	}
}
