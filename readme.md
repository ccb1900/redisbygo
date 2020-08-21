# Redis Server Golang 实现
## 使用
```
make
```

```
cp server.example.json server.json
```

```
./build/redis
```

## benchmark
```
redis-benchmark -h 127.0.0.1 -p 6378 -n 100000 -c 20
```