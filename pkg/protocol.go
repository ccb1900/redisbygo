package pkg

import (
	"strconv"
)

// 错误信息
func ProtocolLineErr(message string) string {
	return protocol(message, "-")
}

// 单行回复，简单字符串，非二进制安全
func ProtocolLine(message string) string {
	return protocol(message, "+")
}

// 多行字符串，多行字符串用于表示长度最大为512 MB的单个二进制安全字符串。
func ProtocolMultiLine(messages []string) string {
	s := protocol(strconv.Itoa(len(messages)), "$")
	for i := 0; i < len(messages); i++ {
		s += protocol(messages[i], "")
	}
	return s
}

// 多行字符串，返回nil
func ProtocolNull() string {
	return protocol("-1", "$")
}

// 数组nil
func ProtocolArrNull() string {
	return protocol("-1", "*")
}

// 整型数字，有符号64位整型
func ProtocolInt(message int) string {
	return protocol(strconv.Itoa(message), ":")
}

// 数组，数组的情况比较复杂，实现时再说
// @todo
func ProtocolArr(messages []string) string {
	s := protocol(strconv.Itoa(len(messages)), "*")
	for i := 0; i < len(messages); i++ {
		s += ProtocolMultiLine([]string{strconv.Itoa(len(messages[i]))})
		s += protocol(messages[i], "")
	}
	return s
}

// 包装协议函数
func protocol(message string, prefix string) string {
	return prefix + message + CRLF
}
