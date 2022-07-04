# Build node
FROM node:16-alpine AS node-builder
WORKDIR /app
COPY client/package.json client/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY client/ .
ENV NEXT_TELEMETRY_DISABLED=1
RUN yarn run export

# Build go
FROM golang:alpine AS go-builder
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
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /app/templates/ ./templates/
COPY --from=go-builder /app/main .
COPY --from=node-builder /app/dist ./client/dist

CMD [ "/app/main" ]