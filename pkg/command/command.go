package command

import (
	"github.com/ccb1900/redisbygo/pkg/command/command"
	"github.com/ccb1900/redisbygo/pkg/command/db"
	string2 "github.com/ccb1900/redisbygo/pkg/command/string"
)

var RedisCommandTable = []command.RedisCommand{
	{"get", string2.GetCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"set", string2.SetCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"select", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"time", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"echo", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"info", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"config", db.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
}
