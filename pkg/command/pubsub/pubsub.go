package pubsub

import (
	"fmt"
	"github.com/ccb1900/redisbygo/pkg"
)

// 发布订阅

func PublishCommand(cl *pkg.Client) {
	s := pkg.NewServer()
	clients := s.PubSubChannels[cl.GetArgvByIndex(1)]
	for i := 0; i < len(clients); i++ {
		clients[i].AddReply(cl.GetArgvByIndex(2))
	}
	cl.AddReply("ok")
}

func SubscribeCommand(cl *pkg.Client) {
	s := pkg.NewServer()

	for i := 1; i < len(cl.Argv); i++ {
		fmt.Println(cl.GetArgvByIndex(i))
		if _, ok := cl.PubSubChannel[cl.GetArgvByIndex(i)]; !ok {
			cl.PubSubChannel[cl.GetArgvByIndex(i)] = true
			s.PubSubChannels[cl.GetArgvByIndex(i)] = append(s.PubSubChannels[cl.GetArgvByIndex(i)], cl)
		}
	}
	cl.AddReply("ok")
}
