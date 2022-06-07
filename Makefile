server: 
	export LOG_TYPE=USER_FRIENDLY && cd ./cmd && go run ./server

client: 
	cd ./cmd && go run ./server

test:
	cd ./cmd && go test -cover ./...

build:
	cd ./cmd && go build -o main ./server