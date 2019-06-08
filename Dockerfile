# Dockerfile for micro service
FROM golang:latest
RUN mkdir -p /go/src/usersapi_go
ADD . /go/src/usersapi_go
WORKDIR /go/src/usersapi_go
RUN go get -v
RUN go install usersapi_go
ENTRYPOINT /go/bin/usersapi_go
EXPOSE 8236