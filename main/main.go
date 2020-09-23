package main

import (
	"flag"
	"fmt"
	"github.com/ccb1900/redisbygo/app/dashboard"
	"github.com/ccb1900/redisbygo/app/server"
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/config"
	"os"
	"sync"
)

func main() {
	version := flag.Bool("version", false, "show redis version")
	help := flag.Bool("h", false, "show help")
	host := flag.String("host", "127.0.0.1", "set host")
	port := flag.Int("p", 6378, "set port")

	flag.Parse()

	if *version {
		fmt.Println("test")
		fmt.Println(pkg.RedisVersion)
	}
	if *help {
		flag.PrintDefaults()
	}
	c := config.NewConfig()
	c.Host = *host
	c.Port = *port
	serve()
}
func version() {
	os.Exit(0)
}

func usage() {

}
func serve() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		dashboard.CreateDashboard()
		defer wg.Done()
	}()

	go func() {
		defer wg.Done()
		server.CreateServer()
	}()

	go func() {
		defer wg.Done()
		pkg.Cron()
	}()
	wg.Wait()
}
