FROM golang:1.18.2

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /douyin
COPY . .
RUN mkdir output
RUN go build -ldflags="-w -s" -o output/server

# 指定运行时环境变量
ENV GIN_MODE=release
EXPOSE 8080

ENTRYPOINT ["./output/server"]
