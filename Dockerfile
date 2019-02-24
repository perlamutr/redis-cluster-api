FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/redis-responder
COPY . .
WORKDIR $GOPATH/src/redis-responder/cmd
RUN go get -d -v
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /main main.go

# This to create smaller image:
FROM scratch

COPY --from=builder /main /

ARG host
ENV REDIS_HOST $host

CMD ["/main"]
