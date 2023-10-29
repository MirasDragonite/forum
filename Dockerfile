FROM golang:1.20.1-alpine3.16 AS builder
RUN apk add build-base
WORKDIR /app
COPY . .

RUN  go build -o main ./cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/ /app/

CMD ["./main"]