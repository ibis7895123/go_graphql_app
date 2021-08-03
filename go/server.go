package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ibis7895123/go_graphql_app/graph"
	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/src/util"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

// ロガー
var logger, _ = zap.NewDevelopment()

func main() {
	// envファイルのロード
	util.EnvLoad()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// time.Timeを扱うためにparseTime=trueが必要
	// タイムゾーンをJSTにする
	// ex.) user:password@tcp(127.0.0.1:3306)/testdb?parseTime=true&loc=Asia%2FTokyo
	dataSource := os.Getenv("MYSQL_DATASOURCE")

	// DB接続
	db, err := gorm.Open("mysql", dataSource)

	// DB起動エラー返す
	if err != nil || db == nil {
		panic(err)
	}

	// DB終了処理(エラー時はエラーを返す)
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()

	// ログ出力をON
	db.LogMode(true)

	// DBの参照を渡す
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{DB: db},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Sugar().Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
