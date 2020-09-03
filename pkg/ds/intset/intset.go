package intset

import "github.com/ccb1900/redisbygo/pkg"

type IntSet struct {
	Encoding pkg.Uint32T
	Contents []pkg.Int8T
}

func NewIntSet() *IntSet {
	o := new(IntSet)
	o.Contents = make([]pkg.Int8T, 0)
	return o
}

func (is *IntSet) Add(value pkg.Int64T, success *pkg.Uint8T) *IntSet {
	return NewIntSet()
}

func (is *IntSet) Remove(value pkg.Int64T, success *pkg.Uint8T) *IntSet {
	return NewIntSet()
}
func (is *IntSet) Find(value pkg.Int64T) pkg.Uint8T {
	return 0
}

func (is *IntSet) Random() pkg.Int64T {
	return 0
}
func (is *IntSet) Get(value pkg.Int64T, success *pkg.Uint8T) pkg.Uint8T {
	return 0
}
func (is *IntSet) Len() pkg.SizeT {
	return pkg.SizeT(len(is.Contents))
}

func (is *IntSet) BlobLen(value pkg.Int64T, success *pkg.Uint8T) *IntSet {
	return NewIntSet()
}
