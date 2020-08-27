package table

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/command"
	"github.com/ccb1900/redisbygo/pkg/command/db"
	string2 "github.com/ccb1900/redisbygo/pkg/command/string"
)

var RedisCommandTable = []pkg.RedisCommand{
	{"get", string2.GetCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"set", string2.SetCommand, -3, "wm", 0, nil, 1, 1, 1, 0, 0},
	{"select", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"time", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"info", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"config", command.ConfigCommand, -2, "last", 0, nil, 0, 0, 0, 0, 0},
	{"ping", command.PingCommand, -1, "tF", 0, nil, 0, 0, 0, 0, 0},
	{"echo", command.EchoCommand, 2, "F", 0, nil, 0, 0, 0, 0, 0},
}
