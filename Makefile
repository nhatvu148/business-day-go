server: 
	export LOG_TYPE=USER_FRIENDLY && go run ./cmd/server

client: 
	go run ./cmd/server

test:
	go test -cover ./...

build:
	go build -o main ./cmd/server