package string

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/client"
)

func test() {

}

func GetCommand(c *client.Client) {

}
func SetCommand(c *client.Client) {
	pkg.Add(c.Db, nil, nil)
}
