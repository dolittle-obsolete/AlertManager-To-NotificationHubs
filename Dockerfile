FROM golang:alpine
ENV CONFIG_PATH /app/config/config.yaml

RUN apk add --no-cache git
RUN mkdir /app
ADD . /app

WORKDIR /app/
RUN go get -d -v
RUN go build -o main .
CMD ["./main"]

EXPOSE 2000