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
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

// ロガー
var logger, _ = zap.NewDevelopment()

func main() {
	// DB接続
	db := util.NewDB()

	// DB終了処理(エラー時はエラーを返す)
	defer db.Close()

	// DBの参照を渡す
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{DB: db},
			},
		),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Sugar().Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
