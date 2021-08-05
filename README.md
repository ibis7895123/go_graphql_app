# GraphQLサーバ(Golang)
todoとユーザを登録するサーバ

## 主要な使用パッケージ
- github.com/99designs/gqlgen
  - GraphQLサーバ
-	github.com/go-sql-driver/mysql
  - mysqlのドライバ
-	github.com/google/uuid
  - ランダムなuuidの生成
-	github.com/jinzhu/gorm
  - データベースのORマッパー
-	github.com/joho/godotenv
  - envファイルを扱う
-	go.uber.org/zap
  - ログの出力

## 環境構築
1. docker起動 `docker-compose up -d`
2. GraphQLサーバ起動 `make server-up`
3. http://localhost:8080 にアクセスするとGraphQLのリクエストが試せる

## テスト(作りかけ)
- 全テストコマンド
`make test-all`

- 個別のテストコマンド(例)
`make test SRC=./src/util/util.go`
`make test SRC=./src/util`

### テストのカバレッジ結果
[cover.html](go/cover.html)