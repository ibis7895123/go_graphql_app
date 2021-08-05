package graph_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ibis7895123/go_graphql_app/graph"
	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/graph/model"
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

func Test_正常系_ユーザ新規作成(t *testing.T) {
	expected_user := model.NewUser{
		Name: "テストユーザ1",
	}

	client := client.New(srv)
	var response struct {
		CreateUser models.User
	}

	// httpリクエスト
	query := `
	mutation {
		createUser(input: {name: "` + expected_user.Name + `" }){
			id,
			name,
			created_at,
			updated_at
		}
	}`
	client.Post(query, &response)

	createUser := response.CreateUser
	assert.NotEqual(t, createUser.ID, "")
	assert.NotEqual(t, createUser.CreatedAt, "")
	assert.NotEqual(t, createUser.UpdatedAt, "")
}

func Test_正常系_全ユーザ取得(t *testing.T) {
	expected_user_oldest := models.User{
		ID:        "122f1a5a728b413c982ae835cf0d84c9",
		Name:      "太郎",
		CreatedAt: "2021-08-01 23:52:20 +0900 JST",
		UpdatedAt: "2021-08-01 23:53:10 +0900 JST",
	}

	client := client.New(srv)
	var response struct {
		Users []models.User
	}

	// httpリクエスト
	query := `
	query {
		users {
			id,
			name,
			created_at,
			updated_at
		}
	}`
	client.Post(query, &response)

	users := response.Users
	assert.Equal(t, len(users), 9)
	assert.Equal(t, users[len(users)-1], expected_user_oldest)
}
