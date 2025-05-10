

FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o load_balancer load_balancer.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/load_balancer .

RUN chmod +x load_balancer

EXPOSE 8080

CMD ["./load_balancer"]
