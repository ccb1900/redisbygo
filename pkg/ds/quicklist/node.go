package quicklist

type Node struct {
	Prev              *Node
	Next              *Node
	Zl                string
	Sz                int
	Count             int
	Encoding          int
	Container         int
	Recompress        int
	AttemptedCompress int
	Extra             int
}
