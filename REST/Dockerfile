FROM golang:1.24.2 AS builder

WORKDIR /app

COPY service.proto ./
COPY REST/ ./REST/

RUN apt-get update -y && \
    apt-get install -y protobuf-compiler && \
    go mod download && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    protoc --go_out=REST --go-grpc_out=REST ./service.proto && \
    CGO_ENABLED=0 GOOS=linux go build -o /http_server ./REST/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /http_server /http_server
COPY --from=builder /app/REST/web_ui /web_ui

RUN chmod +x /http_server

CMD ["/http_server"]
