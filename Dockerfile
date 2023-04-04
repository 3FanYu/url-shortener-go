FROM golang:1.19

WORKDIR /app

COPY /app/ .

# RUN apk add --no-cache git

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.0
RUN	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.0
RUN	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
RUN	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
RUN	go install github.com/favadi/protoc-go-inject-tag@v1.4.0
RUN	go install github.com/envoyproxy/protoc-gen-validate@v0.9.1

RUN go mod download
