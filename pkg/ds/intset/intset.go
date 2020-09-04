package intset

import (
	"fmt"
	"github.com/ccb1900/redisbygo/pkg/types"
)

type IntSet struct {
	Encoding types.Uint32T
	Contents []types.Int8T
}

func NewIntSet() *IntSet {
	o := new(IntSet)
	o.Contents = make([]types.Int8T, 0)
	return o
}

func (is *IntSet) Add(value types.Int64T) *IntSet {
	exist, index := is.Find(value)

	// exists
	if exist > 0 {
		return is
	} else {
		// max
		if index >= len(is.Contents) {
			is.Contents = append(is.Contents, types.Int8T(value))
		} else {
			// middle
			rear := append([]types.Int8T{}, is.Contents[index:]...)

			is.Contents = append(is.Contents[:index], types.Int8T(value))
			is.Contents = append(is.Contents, rear...)
		}
	}

	return is
}

func (is *IntSet) Remove(value types.Int64T) *IntSet {
	exist, index := is.Find(value)
	if exist > 0 {
		fmt.Println(is.Contents[:index+1], is.Contents[index+2:])
		is.Contents = append(is.Contents[:index], is.Contents[index+1:]...)
	}
	return is
}
func (is *IntSet) Find(value types.Int64T) (int, int) {
	if len(is.Contents) == 0 {
		return -1, 0
	}
	mid := len(is.Contents) / 2
	low := 0
	high := len(is.Contents)
	for (mid <= len(is.Contents)-1) && mid >= 0 {
		if value == types.Int64T(is.Contents[mid]) {
			//exists,others are not exist
			return 1, mid
		}
		if value < types.Int64T(is.Contents[mid]) {
			high = mid - 1
		}
		if value > types.Int64T(is.Contents[mid]) {
			low = mid + 1
		}
		if low > high {
			break
		}
		mid = (low + high) / 2
	}
	return -1, low
}

func (is *IntSet) Random() types.Int64T {
	return 0
}
func (is *IntSet) Get(value types.Int64T, success *types.Uint8T) types.Uint8T {
	return 0
}
func (is *IntSet) Len() types.SizeT {
	return types.SizeT(len(is.Contents))
}

func (is *IntSet) BlobLen(value types.Int64T, success *types.Uint8T) *IntSet {
	return NewIntSet()
}
