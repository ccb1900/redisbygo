# Redis Server Golang 实现（开发中...）
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
redis-benchmark -h 127.0.0.1 -p 6378 -n 100000 -c 1000
```

## 参考

https://redis.io/documentation

## 其他

压测需要解除文件描述符限制。

```
在 /etc/launchd.conf 文件写入

limit maxfiles 1000000 1000000
```

在 .zshrc 写入 

```
ulimit -n 10000
```