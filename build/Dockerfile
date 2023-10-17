FROM ubuntu:22.04

WORKDIR /usr/local

RUN useradd -ms /bin/bash golang

RUN apt-get update && apt install curl -y

RUN curl -O -L "https://go.dev/dl/go1.21.1.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz

WORKDIR /app

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH /usr/local/go
ENV GOMAXPROCS=2

RUN chown golang:golang /usr/local/go -R
RUN chown golang:golang /app -R

COPY ./go.mod ./
COPY ./go.sum ./

RUN ["/bin/bash", "-c", "go mod download"]

RUN ["/bin/bash", "-c", "go install -mod=mod github.com/githubnemo/CompileDaemon"]

COPY . .

USER golang

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build -o main ./cmd" -command="./main" -directory="./"