FROM golang:1.18.0-alpine3.15

ARG GO111MODULE=on

RUN apk update && apk add git
RUN mkdir /backend
# ワーキングディレクトリの設定
WORKDIR /backend

COPY . /backend

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.27.3