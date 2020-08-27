package string

import (
	"github.com/ccb1900/redisbygo/pkg"
)

func test() {

}

func GetCommand(c *pkg.Client) {

}
func SetCommand(c *pkg.Client) {
	pkg.Add(c.Db, nil, nil)
}
