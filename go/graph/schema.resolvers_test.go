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

// テストサーバの生成
// 各テストで独立したDBを使うため
func NewTestServer() *handler.Server {
	db := util.NewTestDB()
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{DB: db},
			},
		),
	)

	return srv
}

func Test_正常系_ユーザ新規作成(t *testing.T) {
	t.Parallel()

	expected_user := model.NewUser{
		Name: "テストユーザ1",
	}

	client := client.New(NewTestServer())
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
	t.Parallel()

	expected_users := []models.User{
		{
			ID:        "72412640063b46bc8ace2baf32dc838e",
			Name:      "バンデューラ",
			CreatedAt: "2021-08-03 19:44:29 +0900 JST",
			UpdatedAt: "2021-08-03 19:44:29 +0900 JST",
		},
		{
			ID:        "8ad28da8f22d4a4e98e1c9b477a2186b",
			Name:      "ロック",
			CreatedAt: "2021-08-02 12:20:28 +0900 JST",
			UpdatedAt: "2021-08-02 12:20:28 +0900 JST",
		},
		{
			ID:        "3bc8cf432c64474b9ff54ede13c9457a",
			Name:      "ランド",
			CreatedAt: "2021-08-02 12:13:39 +0900 JST",
			UpdatedAt: "2021-08-02 12:13:39 +0900 JST",
		},
		{
			ID:        "229c019ab5794f87bcbe705e66875ad5",
			Name:      "ダンカン",
			CreatedAt: "2021-08-02 12:12:02 +0900 JST",
			UpdatedAt: "2021-08-02 12:12:02 +0900 JST",
		},
		{
			ID:        "ee2e0dba5f86404996f1b29507cf1443",
			Name:      "ラスティ",
			CreatedAt: "2021-08-02 01:06:38 +0900 JST",
			UpdatedAt: "2021-08-02 01:06:38 +0900 JST",
		},
		{
			ID:        "2fd76f83d5a5421aad0fda992c23709a",
			Name:      "John",
			CreatedAt: "2021-08-01 23:59:47 +0900 JST",
			UpdatedAt: "2021-08-02 00:01:10 +0900 JST",
		},
		{
			ID:        "d19b1061e2d649e1a245f31e36275902",
			Name:      "花子",
			CreatedAt: "2021-08-01 23:52:40 +0900 JST",
			UpdatedAt: "2021-08-01 23:53:20 +0900 JST",
		},
		{
			ID:        "122f1a5a728b413c982ae835cf0d84c9",
			Name:      "太郎",
			CreatedAt: "2021-08-01 23:52:20 +0900 JST",
			UpdatedAt: "2021-08-01 23:53:10 +0900 JST",
		},
	}

	client := client.New(NewTestServer())
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
	assert.Equal(t, len(users), 8)
	assert.Equal(t, users, expected_users)
}
