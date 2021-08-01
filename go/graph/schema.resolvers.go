package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/graph/model"
	"github.com/ibis7895123/go_graphql_app/src/database"
	"github.com/ibis7895123/go_graphql_app/src/models"
	"github.com/ibis7895123/go_graphql_app/src/util"
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// 更新スキーマ関数
type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	return &models.Todo{}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	log.Printf("[mutationResolver.CreateUser] input: %v", input)

	id := util.CreateUniqueID()
	err := database.NewUserDao(r.DB).InsertOne(
		&database.User{
			ID:   id,
			Name: input.Name,
		})

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &models.User{
		ID:   id,
		Name: input.Name,
	}, nil
}

// GETスキーマ関数
type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return []*models.Todo{}, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	return &models.Todo{}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return []*models.User{}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return &models.User{}, nil
}

type todoResolver struct{ *Resolver }

// todoが呼ばれたときにuserを取得
func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return &models.User{}, nil
}

type userResolver struct{ *Resolver }

// userが呼ばれたときにtodosを取得
func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	return []*models.Todo{}, nil
}
