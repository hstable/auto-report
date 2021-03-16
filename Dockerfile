FROM golang:alpine
ADD . /build/
WORKDIR /build/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN go build -o auto-report .
ENTRYPOINT ./auto-report -u $USERNAME -p $PASSWORD -e $EMAIL
