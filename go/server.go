package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ibis7895123/go_graphql_app/graph"
	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"
const dbUser = "localuser"
const dbPassword = "localpass"
const dbProtocol = "tcp(mysql-graphql:3306)"
const dbName = "localdb"
const dbConfig = dbUser + ":" + dbPassword + "@" + dbProtocol + "/" + dbName

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// DB接続
	db, err := gorm.Open("mysql", dbConfig)

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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
