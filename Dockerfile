FROM golang:1.12.1-stretch
WORKDIR /workspace
RUN go build -v
