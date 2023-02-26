############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
MAINTAINER yirufeng yirufeng@foxmail.com
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache 'git=~2'

# Install dependencies
ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/GithubWebhookGo/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main .

############################
# STEP 2 build a small image
############################
FROM alpine:3

WORKDIR /
# Copy our static executable.
COPY --from=builder /go/main /go/main
COPY public /go/public

ENV PORT 8080
ENV GIN_MODE release

# 暴露端口
EXPOSE 8999

WORKDIR /go

# 执行启动命令
ENTRYPOINT [ "./main" ]