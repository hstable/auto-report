FROM golang:alpine AS builder
ADD . /build/
WORKDIR /build/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN go build -ldflags "-s -w" -o auto-report .

FROM alpine:latest
COPY --from=builder /build/auto-report /usr/bin
ENTRYPOINT ["auto-report"]
