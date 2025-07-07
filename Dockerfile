FROM golang:1.24-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" \
    -o /order-service ./cmd/order-service
    

FROM alpine:3.20

RUN apk add --no-cache ca-certificates curl

ENV MIGRATE_VERSION=v4.18.3
RUN curl -sSL \
      "https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz" \
    | tar -xz -C /usr/local/bin && chmod +x /usr/local/bin/migrate

WORKDIR /app
COPY --from=build /order-service .
COPY migrations ./migrations
COPY cmd/order-service/web ./web

CMD migrate -path /app/migrations \
            -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" \
            up \
 && exec /app/order-service