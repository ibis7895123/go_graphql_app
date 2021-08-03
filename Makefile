# GraphQLサーバ起動
.PHONY: server-up
server-up:
	docker compose exec go-graphql go run server.go

# graphQL系ファイルの生成(gqlgenを使用)
.PHONY: gqlgen-init
gqlgen-init:
	cd go && go run github.com/99designs/gqlgen init
