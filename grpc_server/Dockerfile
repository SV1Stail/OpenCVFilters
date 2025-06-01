FROM golang:1.24.2 AS builder

WORKDIR /app

COPY service.proto ./
COPY grpc_server ./grpc_server/

RUN apt-get install libopencv-dev && \
    apt-get install protobuf-compiler && \
    go mod download && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    protoc --go-out=grpc_server --go-grpc_out=grpc_server ./service.proto && \
    CGO_ENABLED=0 GOOS=linux go build -o /server ./grpc_server/main.go

FROM lpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /server /server

RUN chmod +x server

CMD ["/server"]
