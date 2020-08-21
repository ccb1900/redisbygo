package string

import (
	"redis/pkg/client"
	"redis/pkg/redisdb/redisdb"
)

func test() {

}

func GetCommand(c *client.Client) {

}
func SetCommand(c *client.Client) {
	redisdb.Add(c.Db, nil, nil)
}
