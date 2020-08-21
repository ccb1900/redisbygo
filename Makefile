
build:
	go build -o build/redis main/main.go && cp server.json build/server.json
clean:
	rm  -rf build