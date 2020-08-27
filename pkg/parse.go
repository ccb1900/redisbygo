package pkg

import (
	"bytes"
	"strconv"
)

func ParseProtocol() {
	s := []byte("*3\r\n$3\r\nset\r\n$1\r\na\r\n$1\r\nb\r\n")

	buf := make([]byte, 0)
	mulBulkLen := 0
	argv := make([]string, 0)
	//p := 0
	for i := 0; i < len(s); i += 5 {
		end := i + 5

		if end > len(s) {
			end = len(s)
		}
		queryBuf := s[i:end]

		// 如果buf包含了多行
		buf = append(buf, queryBuf...)
		// 检查协议长度
		if mulBulkLen == 0 {
			pos := bytes.Index(buf, []byte{'\r', '\n'})
			if pos > 0 && buf[0] == '*' {
				mulBulkLen = S2Int(string(buf[1:pos]))
				buf = buf[pos+2:]
			}
		} else {
			for mulBulkLen > 0 {
				pos := bytes.Index(buf[0:], []byte{'\r', '\n'})

				if pos == -1 {
					break
				}
				if buf[0] == '$' {
					//bulkLen = S2Int(string(buf[1:pos]))
				} else {
					argv = append(argv, string(buf[0:pos]))
					mulBulkLen--
				}
				buf = buf[pos+2:]
			}
		}
	}
}

func S2Int(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
