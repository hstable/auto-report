FROM golang:alpine
ADD . /build/
WORKDIR /build/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
