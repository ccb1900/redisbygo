package command

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/shared"
)

func EchoCommand(cl *pkg.Client) {
	cl.AddReplyBulk(cl.Argv[1])
}

func TimeCommand(cl *pkg.Client) {
	cl.AddReply("echo")
}

func PingCommand(cl *pkg.Client) {
	if len(cl.Argv) > 2 {
		cl.AddReplyErrorFormat([]string{"wrong number of arguments for '%s' command"}, cl.Cmd.Name)
	} else {
		if len(cl.Argv) == 1 {
			sh := shared.NewShared()
			cl.AddReplyRedisObject(sh.Pong)
		} else {
			cl.AddReplyBulk(cl.Argv[1])
		}
	}
}

func AuthCommand(cl *pkg.Client) {
	cl.AddReply("echo")
}

func CommandCommand(cl *pkg.Client) {
	cl.AddReply("echo")
}

func InfoCommand(cl *pkg.Client) {
	cl.AddReply("echo")
}

func MonitorCommand(cl *pkg.Client) {
	cl.AddReply("echo")
}
