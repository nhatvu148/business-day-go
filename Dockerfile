# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/server

# Run stage
FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cmd/web/templates/ ./cmd/web/templates/
COPY --from=builder /app/main .

CMD [ "/app/main" ]