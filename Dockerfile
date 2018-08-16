FROM golang
ADD . /go/src/github.com/asridjufri/ut-demo
RUN go build /go/src/github.com/asridjufri/ut-demo/hello.go
ENTRYPOINT /go/src/github.com/asridjufri/ut-demo/hello
EXPOSE 9000
