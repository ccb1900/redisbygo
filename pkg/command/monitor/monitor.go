package monitor

import "github.com/ccb1900/redisbygo/pkg"

// 监视器

func MonitorCommand(cl *pkg.Client) {
	s := pkg.NewServer()
	s.Monitors = append(s.Monitors, cl)
	cl.AddReply("ok")
}
