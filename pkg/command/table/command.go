package table

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/command"
	"github.com/ccb1900/redisbygo/pkg/command/connection"
	"github.com/ccb1900/redisbygo/pkg/command/pubsub"
	"github.com/ccb1900/redisbygo/pkg/command/server"
	"github.com/ccb1900/redisbygo/pkg/command/set"
	string2 "github.com/ccb1900/redisbygo/pkg/command/string"
)

var RedisCommandTable = []pkg.RedisCommand{
	{"get", string2.GetCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"set", string2.SetCommand, -3, "wm", 0, nil, 1, 1, 1, 0, 0},
	{"slaveof", server.ReplicationOfCommand, 3, "ast", 0, nil, 0, 0, 0, 0, 0},
	{"replicaof", server.ReplicationOfCommand, 3, "ast", 0, nil, 0, 0, 0, 0, 0},
	{"select", connection.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"time", connection.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"info", connection.SelectDbCommand, 2, "rF", 0, nil, 1, 1, 1, 0, 0},
	{"config", command.ConfigCommand, -2, "last", 0, nil, 0, 0, 0, 0, 0},
	{"ping", command.PingCommand, -1, "tF", 0, nil, 0, 0, 0, 0, 0},
	{"echo", command.EchoCommand, 2, "F", 0, nil, 0, 0, 0, 0, 0},
	{"command", command.Command, 0, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"publish", pubsub.PublishCommand, 3, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"pubsub", command.Command, -2, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"subscribe", pubsub.SubscribeCommand, -2, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"psubscribe", command.Command, -2, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"unsubscribe", command.Command, -1, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"punsubscribe", command.Command, -1, "ltR", 0, nil, 0, 0, 0, 0, 0},
	{"monitor", command.MonitorCommand, 1, "as", 0, nil, 0, 0, 0, 0, 0},
	{"sadd", set.SAddCommand, -3, "wmF", 0, nil, 1, 1, 1, 0, 0},
}
