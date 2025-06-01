FROM golang:1.24.2 AS builder

WORKDIR /app

COPY service.proto ./
COPY grpc_server ./


RUN  apt-get update -y && \
    apt-get install -y software-properties-common && \
    add-apt-repository -y "deb http://security.debian.org/debian-security bullseye-security main" 
# apt-get update -y

RUN apt-get install -y libopencv-dev 

RUN apt-get install -y protobuf-compiler && \
    go get -u gocv.io/x/gocv && \
    go mod download && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 

RUN protoc --go_out=. --go-grpc_out=. ./service.proto && \
    CGO_ENABLED=0 GOOS=linux go build -o /server ./main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /server /server

CMD ["/server"]