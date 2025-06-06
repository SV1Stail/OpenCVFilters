FROM gocv/opencv:latest AS builder

WORKDIR /app

RUN apt-get update -y && \
    apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 

COPY grpc_server/go.mod grpc_server/go.sum ./
RUN go mod download


COPY service.proto ./
COPY grpc_server ./

RUN protoc --go_out=. --go-grpc_out=. ./service.proto && \
    CGO_ENABLED=1 GOOS=linux go build -o /grpc_application ./main.go

ENTRYPOINT ["/grpc_application"]
