package intset

import (
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/types"
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	ll := []int{
		1, 4, 3, 2, 7, 6, 9, 0, 0, 9, 8, 13, 2,
	}
	is := NewIntSet()
	for i := 0; i < len(ll); i++ {
		is.Add(types.Int64T(ll[i]))
	}

	fmt.Println(is.Contents)

	is.Remove(4)
	fmt.Println(is.Contents)
}
