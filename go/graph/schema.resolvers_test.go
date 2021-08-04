package graph_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ibis7895123/go_graphql_app/graph"
	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/src/util"
	"go.uber.org/zap"
)

// ロガー
var logger, _ = zap.NewDevelopment()

// DB接続
var db = util.NewDB()

// httpハンドラー
var srv = handler.NewDefaultServer(
	generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{DB: db},
		},
	),
)

func Test_正常系_全ユーザ取得(t *testing.T) {
	// DB終了処理(エラー時はエラーを返す)
	defer db.Close()

	client := client.New(srv)
	var response interface{}

	client.Post(`query { users { id, name } }`, &response)

	logger.Debug("response", zap.Any("users", response))
}
