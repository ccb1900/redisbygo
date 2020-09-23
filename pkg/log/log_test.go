package log

import "testing"

func TestNewLog(t *testing.T) {
	l := NewLog(*t)

	l.Info("hello")
}
