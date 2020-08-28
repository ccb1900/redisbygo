package quicklist

type QuickList struct {
	Head     *Node
	Tail     *Node
	Count    int
	Len      int
	Fill     int
	Compress int
}

const TAIL = -1
const Head = 0

func CreateQuickList() *QuickList {
	ql := new(QuickList)
	return ql
}
