# RedisServerByGo（ in development... ）

This is a redis server implemented by golang

## Usage
```
make
```

```
cp server.example.json server.json
```

```
./build/redis
```

## Benchmark
```
redis-benchmark -h 127.0.0.1 -p 6378 -n 100000 -c 1000
```

## Reference

https://redis.io/documentation