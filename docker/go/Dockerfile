FROM golang:latest

RUN mkdir /go_graphql_app
COPY go /go_graphql_app
WORKDIR /go_graphql_app

# タイムゾーンをJSTに変更
COPY ./docker/go/zoneinfo-Japan /etc/localtime