# RedisServerByGo（ 开发中... ）

[English](/readme.md)

这是一个golang实现的redis服务器

如果发现了bug，或者有相关的建议，可以提出来。

还有很多功能需要实现，但是基本的结构已经完备了。我有时间会持续开发的。

本项目是本着学习的目的，毕竟 go 在内存效率方面是无法和 c 比拟的，所以也不考虑能用在生产环境。

## Repository

- [Gitee,中国用户访问](https://gitee.com/waterloocode/redisbygo)
- [Github](https://github.com/ccb1900/redisbygo)

## 已支持

- 加载aof文件
- ping 命令
- echo 命令
- pubsub 命令
- subscribe 命令
- get 命令
- set 命令
- select 命令

## 用法

```shell
make
```

```shell
cp server.example.json server.json
```

### win

```shell
./build/windows/redis.exe
```

### mac

```shell
./build/darwin/redis
```

### linux

```shell
./build/linux/redis
```

## 测试

```shell
redis-benchmark -h 127.0.0.1 -p 6378 -n 100000 -c 1000
```

## 参考

[Redis 官方文档](https://redis.io/documentation)