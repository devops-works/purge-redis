FROM golang:latest AS builder

WORKDIR /app

COPY . /app

RUN go build -a -o purge-redis main.go && strip purge-redis



FROM gcr.io/distroless/base

COPY --from=builder /app/purge-redis /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/purge-redis"]
