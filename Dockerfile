FROM golang:1.8.5-jessie as builder

ADD src /go/app/src/
ADD config /go/app/config

WORKDIR /go/app/src/
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine:3.7
ENV CONFIG_PATH /root/config/config.yaml
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/app/ .
CMD ["src/main"]

EXPOSE 2000