.PHONY: server
server: 
	export LOG_TYPE=USER_FRIENDLY && cd ./server && go run main.go

client: 
	cd ./cmd/cli && go run main.go

test:
	cd ./server && go test -cover ./...

build:
	cd ./server && go build -o main main.go