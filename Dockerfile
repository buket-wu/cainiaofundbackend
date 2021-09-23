# build
FROM golang:1.16 as builder

ENV GO111MODULE=on \
    GOOS=linux \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /Users/wujiajin/Code/go/src/cainiaofundbackend

COPY . .

RUN go build -o /app

FROM Alpine:last
EXPOSE 80
COPY --from=builder /app /
CMD ["/app"]


