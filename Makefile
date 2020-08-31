
build:
	CGO_ENABLED=0  GOARCH=amd64 go build -o build/darwin/redis main/main.go && GOOS=linux go build -o build/linux/redis main/main.go && GOOS=windows  go build -o build/windows/redis.exe main/main.go
clean:
	rm  -rf build && rm -rf appendonly.aof