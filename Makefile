
build:
	go build -o build/redis main/main.go
clean:
	rm  -rf build && rm -rf appendonly.aof