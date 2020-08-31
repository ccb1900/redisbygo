package server

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/shared"
	"strconv"
	"strings"
)

func ReplicationOfCommand(cl *pkg.Client) {
	s := pkg.NewServer()
	// 判断是否是集群，集群中该命令无法使用

	if !strings.EqualFold(cl.GetArgvByIndex(1), "no") && strings.EqualFold(cl.GetArgvByIndex(2), "one") {
		if s.Main.Host != "" {
			s.Main.ReplicationUnsetMaster()
		}
	} else {
		if cl.Flags&pkg.ClientSlave != 0 {

		}
	}

	s.Main.Host = cl.GetArgvByIndex(1)
	port, _ := strconv.Atoi(cl.GetArgvByIndex(2))
	s.Main.Port = port

	sh := shared.NewShared()
	cl.AddReplyRedisObject(sh.Ok)
}
