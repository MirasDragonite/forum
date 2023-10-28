FROM golang:latest as build

WORKDIR /app
COPY  . .
RUN go build -o main ./cmd/main.go

FROM scratch
COPY --from=build /app/main .
CMD ["/app/main"]
