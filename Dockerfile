# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN cd ./cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./server

# Run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/main .

CMD [ "/app/main" ]