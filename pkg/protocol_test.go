package pkg

import "testing"

func TestProtocolMultiLine(t *testing.T) {
	t.Log(ProtocolMultiLine([]string{
		"hello",
		"world",
		"tome",
		"green",
	}))
}
