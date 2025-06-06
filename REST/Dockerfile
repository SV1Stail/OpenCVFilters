FROM golang:latest AS builder

WORKDIR /app

RUN apt-get update -y && \
    apt-get install -y protobuf-compiler 

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 

COPY REST/go.mod REST/go.sum ./
RUN go mod download


COPY service.proto ./
COPY REST/ ./

RUN protoc --go_out=. --go-grpc_out=. ./service.proto && \
    CGO_ENABLED=1 GOOS=linux go build -o /http_application ./main.go

FROM debian:stable-slim 
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /http_application /http_application
COPY --from=builder /app/web_ui/ ./web_ui/

ENTRYPOINT ["/http_application"]
