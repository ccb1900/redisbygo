package list

import "github.com/ccb1900/redisbygo/pkg/types"

func DeleteInt8T(s *[]types.Int8T, i int) {
	*s = append((*s)[:i], (*s)[i+1:]...)
}
func InsertInt8T(s *[]types.Int8T, index int, e types.Int8T) {
	rear := append([]types.Int8T{}, (*s)[index:]...)

	*s = append((*s)[:index], e)
	*s = append(*s, rear...)
}

func DeleteInt(s *[]int, i int) {
	*s = append((*s)[:i], (*s)[i+1:]...)
}

func InsertInt(s *[]int, index int, e int) {
	rear := append([]int{}, (*s)[index:]...)
	*s = append((*s)[:index], e)
	*s = append(*s, rear...)
}
