package config

import "testing"

func TestNewConfig(t *testing.T) {
	c1 := GetInstance("../../server.json")
	c2 := GetInstance("../../server.json")

	if c1 == c2 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
