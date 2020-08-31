# RedisServerByGo（ in development... ）

This is a redis server implemented by golang

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
```
make
```

```
cp server.example.json server.json
```

### win
```
./build/windows/redis.exe
```

### mac
```
./build/darwin/redis
```

### linux
```
./build/linux/redis
```

## Benchmark
```
redis-benchmark -h 127.0.0.1 -p 6378 -n 100000 -c 1000
```

## Reference

https://redis.io/documentation