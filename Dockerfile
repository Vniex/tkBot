FROM golang:1.10.3

WORKDIR $GOPATH/src/tkBot
COPY . $GOPATH/src/tkBot
RUN go build .

EXPOSE 8888
ENTRYPOINT ["./tkBot"]
