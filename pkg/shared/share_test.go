package shared

import (
	"testing"
)

func TestNewShared(t *testing.T) {
	s1 := NewShared()
	s2 := NewShared()
	if s1 == s2 {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}
}
