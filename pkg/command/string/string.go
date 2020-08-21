package string

import (
	"github.com/ccb1900/redisbygo/pkg/client"
	"github.com/ccb1900/redisbygo/pkg/redisdb/redisdb"
)

func test() {

}

func GetCommand(c *client.Client) {

}
func SetCommand(c *client.Client) {
	redisdb.Add(c.Db, nil, nil)
}
