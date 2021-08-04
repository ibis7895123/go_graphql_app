package graph_test

import (
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ibis7895123/go_graphql_app/graph"
	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/src/util"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// ロガー
var logger, _ = zap.NewDevelopment()

func Test_正常系_全ユーザ取得(t *testing.T) {
	util.EnvLoad("../.env")

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

	client := client.New(srv)
	var response interface{}

	client.Post(`query { users { id, name } }`, &response)

	logger.Debug("response", zap.Any("users", response))
}
