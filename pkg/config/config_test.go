package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c1 := NewConfig()
	c2 := NewConfig()

	if c1 == c2 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
