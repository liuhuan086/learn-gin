FROM golang:latest

ENV GOPROXY https://goproxy.cn, direct
WORKDIR $GOPATH/src/github.com/liuhuan086/example
COPY . $GOPATH/src/github.com/liuhuan086/example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./example"]