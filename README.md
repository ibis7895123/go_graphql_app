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

- フォルダ / ファイルテストコマンド(例)
`make test-file src=./src/util/util.go`
`make test-file src=./src/util`

- 関数テストコマンド(例)
`make test-func src=./src/util func=Test_正常系_CreateUniqueID`

### テストのカバレッジ結果
[cover.html](go/cover.html)
