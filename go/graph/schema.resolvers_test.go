package graph_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ibis7895123/go_graphql_app/graph"
	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/src/models"
	"github.com/ibis7895123/go_graphql_app/src/util"
	"github.com/stretchr/testify/assert"
)

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
	expected_user_latest := models.User{
		ID:        "72412640063b46bc8ace2baf32dc838e",
		Name:      "バンデューラ",
		CreatedAt: "2021-08-03 19:44:29 +0900 JST",
		UpdatedAt: "2021-08-03 19:44:29 +0900 JST",
	}

	// DB終了処理
	defer db.Close()

	client := client.New(srv)
	var response struct {
		Users []models.User
	}

	// httpリクエスト
	client.Post(`query { users { id, name, created_at, updated_at } }`, &response)

	assert.Equal(t, len(response.Users), 8)
	assert.Equal(t, response.Users[0], expected_user_latest)
}
