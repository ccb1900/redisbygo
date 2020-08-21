package methods

import "redis/pkg/ds/sds"

func Sdsempty() *sds.Sds {
	return new(sds.Sds)
}
