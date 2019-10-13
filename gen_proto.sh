#!/bin/sh

mkdir -p api/swagger
tee api/swagger/swagger.json >/dev/null <<EOF
{
    "consumes": [
      "application/json"
    ],
    "produces": [
      "application/json"
    ],
    "schemes": [
      "http",
      "https"
    ],
    "host": "127.0.0.1",
    "swagger": "2.0",
    "info": {
      "description": "go-grpc RESTful API",
      "title": "go-grpc api",
      "version": "1.0.0"
    }
  }
EOF

# install protoc https://github.com/protocolbuffers/protobuf/releases/
# ENV GO111MODULE=on
# ENV GOPROXY=https://goproxy.io
# RUN go get -u google.golang.org/grpc
# RUN go get -u github.com/golang/protobuf/protoc-gen-go
# RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
protoc -Iapi/proto -I/usr/local/include -I${GOPATH}/src \
-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:api/proto \
--grpc-gateway_out=logtostderr=true:api/proto \
--swagger_out=logtostderr=true:api/swagger \
./api/proto/*.proto

#install go-swagger https://github.com/go-swagger/go-swagger
swagger -q mixin api/swagger/*.json -o api/swagger/swagger.json