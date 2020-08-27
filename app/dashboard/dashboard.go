package dashboard

import (
	"github.com/ccb1900/redisbygo/pkg"
	"github.com/ccb1900/redisbygo/pkg/command/table"
	"github.com/ccb1900/redisbygo/pkg/config"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
	"time"
)

// 根路由
func root(r *gin.Engine) func(c *gin.Context) {
	return func(c *gin.Context) {
		results := make([]string, 0)
		for _, route := range r.Routes() {
			results = append(results, route.Path)
		}
		c.JSON(200, results)
	}
}

func commands(c *gin.Context) {
	results := make([]string, 0)
	for _, redisCommand := range table.RedisCommandTable {
		results = append(results, redisCommand.Name)
	}
	c.JSON(200, results)
}

func configs(c *gin.Context) {
	s := config.NewConfig()
	c.JSON(200, s)
}

func clients(c *gin.Context) {
	s := pkg.NewServer()
	type cl struct {
		Addr  string
		Index int
		Db    int
	}
	results := make([]*cl, 0)
	for _, cc := range s.Clients {
		results = append(results, &cl{
			Addr:  cc.Conn.RemoteAddr().String(),
			Index: cc.Index,
			Db:    cc.Db.Id,
		})
	}
	c.JSON(200, results)
}

// 由于map不能并发访问，此方法会导致服务器挂掉
func keys(c *gin.Context) {
	type key struct {
		Name string
		Type string
		Db   int
	}
	results := make([]key, 0)

	s := pkg.NewServer()

	for _, d := range s.Db {
		dd := d.Dict.Storage
		for k := range dd {
			results = append(results, key{
				Name: k,
				Type: k,
				Db:   d.Id,
			})
		}
	}

	c.JSON(200, results)
}

func metrics(c *gin.Context) {
	type ServerInfo struct {
		RedisVersion   string
		CRLF           string
		GoVersion      string
		CPUNum         int
		GoRoutineNum   int
		Arch           string
		Goos           string
		GoRoot         string
		Date           string
		Hostname       string
		Memory         string
		Disk           string
		RedisKeyNum    int
		RedisClientNum int
		Envs           []string
	}

	s := pkg.NewServer()

	hostname, _ := os.Hostname()
	count := 0

	for i := 0; i < len(s.Db); i++ {
		count += len(s.Db[i].Dict.Storage)
	}
	c.JSON(200, ServerInfo{
		RedisVersion:   pkg.RedisVersion,
		CRLF:           pkg.CRLF,
		GoVersion:      runtime.Version(),
		CPUNum:         runtime.NumCPU(),
		GoRoutineNum:   runtime.NumGoroutine(),
		Arch:           runtime.GOARCH,
		Goos:           runtime.GOOS,
		GoRoot:         runtime.GOROOT(),
		Date:           time.Now().Format("2006-01-02 15:04:05"),
		Hostname:       hostname,
		Envs:           os.Environ(),
		RedisClientNum: len(s.Clients),
		RedisKeyNum:    count,
	})
}
func CreateDashboard() {
	r := gin.Default()
	// 获取所有路由
	r.GET("/", root(r))
	// 获取所有命令
	r.GET("commands", commands)
	// 获取所有配置
	r.GET("configs", configs)
	// 获取当前连接的客户端
	r.GET("clients", clients)
	// 获取所有的key，暂时未支持分页
	r.GET("keys", keys)
	// 获取服务器监控信息
	r.GET("metrics", metrics)
	_ = r.Run(":9002")
}
