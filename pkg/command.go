package pkg

import "fmt"

type RedisCommand struct {
	Name         string
	Proc         func(c *Client)
	Arity        int
	SFlags       string
	Flags        int
	GetKeysProc  *int
	FirstKey     int
	LastKey      int
	KeyStep      int
	Microseconds int
	Calls        int
}

func (rc *RedisCommand) Propagate(dbid int, argv []*RedisObject, flags int) {
	rc.FeedAppendOnlyFile(dbid, argv)
	replicationFeedSlaves()
}

func (rc *RedisCommand) FeedAppendOnlyFile(dbid int, argv []*RedisObject) {
	s := NewServer()
	buf := ""
	if dbid != s.AofSelectedDb {
		buf = fmt.Sprintf("*2\\r\\n$6\\r\\nSELECT\\r\\n$%s\\r\\n%s\\r\\n", "2", "2")
		s.AofSelectedDb = dbid
	}
	buf = CatAppendOnlyGenericCommand(buf, argv)
}
