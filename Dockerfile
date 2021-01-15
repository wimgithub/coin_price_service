FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/coin_price_service
COPY . $GOPATH/src/coin_price_service
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./coin_price_service"]
