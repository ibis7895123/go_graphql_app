FROM golang:latest

# ソースのセット
RUN mkdir /go_graphql_app
COPY go /go_graphql_app

# ディレクトリセット
WORKDIR /go_graphql_app

# sql-migrateのインストール
RUN go get github.com/rubenv/sql-migrate/...

# タイムゾーンをJSTに変更
COPY ./docker/go/zoneinfo-Japan /etc/localtime