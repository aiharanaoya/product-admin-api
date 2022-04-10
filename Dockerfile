FROM golang:1.18.0-alpine

RUN apk update && apk add git

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

# ローカルではdocker-composeでボリュームマウントしているので不要
# COPY ./main.go ./
# COPY ./app ./app
