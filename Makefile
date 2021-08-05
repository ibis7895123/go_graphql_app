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

# 全テスト実行
# goフォルダ内のすべてのテストファイルを実行する
# カバレッジ解析したcover.htmlを出力する
.PHONY: test-all
test-all:
	docker-compose exec go-graphql go test -v -cover -coverprofile=cover.out ./...
	docker-compose exec go-graphql go tool cover -html=cover.out -o cover_all.html
	docker-compose exec go-graphql rm -f cover.out

# 指定したテスト実行
# src= で指定したフォルダ、ファイルを実行する
# カバレッジ解析したcover.htmlを出力する
.PHONY: test-file
test-file:
	docker-compose exec go-graphql go test -v -cover -coverprofile=cover.out $(src)
	docker-compose exec go-graphql go tool cover -html=cover.out -o cover.html
	docker-compose exec go-graphql rm -f cover.out

# 指定したテスト実行(関数)
# src= で指定したフォルダ、ファイル内の
# func= で指定したテスト関数のみを実行する
.PHONY: test-func
test-func:
	docker-compose exec go-graphql go test -v $(src) -run $(func)