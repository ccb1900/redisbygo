package main

import (
	"flag"
	"fmt"
	"os"
	"redis/app/dashboard"
	"redis/app/server"
	"redis/pkg/config"
	"redis/pkg/others"
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
		fmt.Println(others.RedisVersion)
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
	wg.Add(2)

	go func() {
		dashboard.CreateDashboard()
		defer wg.Done()
	}()

	go func() {
		server.CreateServer()
		defer wg.Done()
	}()

	wg.Wait()
}
