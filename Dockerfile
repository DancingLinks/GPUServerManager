FROM alpine

ADD . /go/src/github.com/GPUServerManager

ENTRYPOINT /go/src/github.com/GPUServerManager/go_build

EXPOSE 8088
