FROM golang:latest

RUN mkdir /detahardd-go
WORKDIR /detahardd-go
COPY ./scripts/run_in_docker.sh /detahardd-go

RUN apt-get update
RUN apt-get install -y redir

RUN go get github.com/detahard/detahardd-go
RUN go build github.com/detahard/detahardd-go

ENTRYPOINT '/detahardd-go/run_in_docker.sh'
EXPOSE 11325
