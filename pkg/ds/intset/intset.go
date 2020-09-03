package intset

import (
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

func (is *IntSet) Add(value types.Int64T, success *types.Uint8T) *IntSet {
	return NewIntSet()
}

func (is *IntSet) Remove(value types.Int64T, success *types.Uint8T) *IntSet {
	return NewIntSet()
}
func (is *IntSet) Find(value types.Int64T) types.Uint8T {
	return 0
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
