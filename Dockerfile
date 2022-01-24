FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/yanzhacheng/pp-backend
COPY . $GOPATH/src/github.com/yanzhacheng/pp-backend
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./pp-backend"]