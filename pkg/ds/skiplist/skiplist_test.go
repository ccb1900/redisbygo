package skiplist

import (
	"testing"
)

func TestRandomLevel(t *testing.T) {
	a := [64]int{}
	for i := 0; i < 1000000; i++ {
		a[RandomLevel()]++
	}

	t.Log(a)
}

func TestNewSkipList(t *testing.T) {
	t.Log(NewSkipList())
}
