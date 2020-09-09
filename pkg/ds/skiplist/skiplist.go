package skiplist

import (
	"github.com/ccb1900/redisbygo/pkg/ds/sds"
	"github.com/ccb1900/redisbygo/pkg/types"
	"math/rand"
	"time"
)

const (
	skiplistmaxlevel = 64
	skiplistp        = 0.25
)

type ZLevel struct {
	forward *Node
	span    int
}

type Node struct {
	level    []ZLevel
	score    float64
	ele      *sds.Sds
	backward *Node
}
type SkipList struct {
	Header *Node
	Tail   *Node
	length types.Long
	Level  int
}

func NewNode(level int, score float64, ele *sds.Sds) *Node {
	n := new(Node)
	n.level = make([]ZLevel, level)
	n.score = score
	n.ele = ele

	return n
}
func NewSkipList() *SkipList {
	s := new(SkipList)
	s.Level = 1
	s.Header = NewNode(skiplistmaxlevel, 0, nil)
	return s
}

func RandomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for float64(rand.Int63()&0xffff) < (skiplistp * 0xffff) {
		rand.Seed(time.Now().UnixNano())
		level += 1
	}

	if level < skiplistmaxlevel {
		return level
	}
	return skiplistmaxlevel
}
func (sl *SkipList) Insert(score float64, ele *sds.Sds) {

}

func (sl *SkipList) deleteNode(score float64, ele *sds.Sds) {

}
