# RedisServerByGo（ in development... ）

This is a redis server implemented by golang

If you find some bugs,just give me an issue.

There is many features what need to be implemented.When I have time,I will finish it.

## Repository

- [Gitee,where Chinese can visit](https://gitee.com/waterloocode/redisbygo)
- [Github](https://github.com/ccb1900/redisbygo)

## Supported

- load aof file
- command ping
- command echo
- command pubsub
- command subscribe
- command get
- command set
- command select

## Usage

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

## Benchmark

```shell
redis-benchmark -h 127.0.0.1 -p 6378 -n 100000 -c 1000
```

## Reference

[Redis Doc](https://redis.io/documentation)