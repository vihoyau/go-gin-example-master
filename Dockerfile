FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/qiuweihao/go-gin-web
COPY . $GOPATH/src/github.com/qiuweihao/go-gin-web
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]