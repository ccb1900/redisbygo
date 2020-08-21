package methods

import "github.com/ccb1900/redisbygo/pkg/ds/sds"

func Sdsempty() *sds.Sds {
	return new(sds.Sds)
}
