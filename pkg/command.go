package pkg

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

func (rc *RedisCommand) Run(cl *Client) {

}
