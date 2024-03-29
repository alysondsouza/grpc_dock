FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

COPY go.mod /build
COPY go.sum /build/

RUN cd /build/ && git clone https://github.com/alysondsouza/grpc_dock.git
RUN cd /build/grpc_dock/server && go build ./...

EXPOSE 9080

ENTRYPOINT [ "/build/grpc_dock/server/server" ]