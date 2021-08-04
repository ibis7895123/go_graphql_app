# GraphQLサーバ起動
.PHONY: server-up
server-up:
	docker compose exec go-graphql go run server.go

# graphQL系ファイルの生成(gqlgenを使用)
.PHONY: gqlgen-init
gqlgen-init:
	cd go && go run github.com/99designs/gqlgen init

# sql-migrateのステータス確認
.PHONY: migrate-status
migrate-status:
	docker-compose exec go-graphql sql-migrate status

# テスト実行
# goフォルダ内のすべてのテストファイルを実行する
.PHONY: test
test:
	docker-compose exec go-graphql go test -v ./...