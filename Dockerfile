# Build Stage
FROM golang:1.12 AS build-stage

LABEL APP="build-ops"
LABEL REPO="https://gitlab.ctyuncdn.cn/axe/ops"

ADD . /go/src/github.com/Icedroid/go-grpc
WORKDIR /go/src/github.com/Icedroid/go-grpc

RUN make build-alpine

# Final Stage
FROM alpine:3.10

# ARG GIT_COMMIT
# ARG VERSION
# ARG APP_NAME

# LABEL REPO="https://github.com/Icedroid/go-grpc"
# LABEL GIT_COMMIT=${GIT_COMMIT}
# LABEL VERSION=${VERSION}
# LABEL APP_NAME=${APP_NAME}

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache tcpdump lsof net-tools tzdata curl dumb-init libc6-compat

ENV TZ Asia/Shanghai
ENV PATH=$PATH:/opt/go-grpc/bin

WORKDIR /opt/go-grpc/bin

COPY --from=build-stage /go/src/github.com/Icedroid/go-grpc/bin/go-grpc .
RUN chmod +x /opt/go-grpc/bin/go-grpc

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/go-grpc/bin/go-grpc"]