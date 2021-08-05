package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/ibis7895123/go_graphql_app/graph/generated"
	"github.com/ibis7895123/go_graphql_app/graph/model"
	"github.com/ibis7895123/go_graphql_app/src/database"
	"github.com/ibis7895123/go_graphql_app/src/logger"
	"github.com/ibis7895123/go_graphql_app/src/models"
	"github.com/ibis7895123/go_graphql_app/src/util"
	"go.uber.org/zap"
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
	logger.Logger.Info("[mutationResolver.CreateTodo] ", zap.Any("input", input))

	// ユーザーIDがあるかどうかチェック
	err := database.NewUserDao(r.DB).ExistUserID(input.UserID)

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	id := util.CreateUniqueID()
	err = database.NewTodoDao(r.DB).InsertOne(
		&database.Todo{
			ID:        id,
			Text:      input.Text,
			Done:      false,
			UserID:    input.UserID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	return &models.Todo{
		ID:        id,
		Text:      input.Text,
		Done:      false,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	logger.Logger.Info("[mutationResolver.CreateUser] ", zap.Any("input", input))

	id := util.CreateUniqueID()
	err := database.NewUserDao(r.DB).InsertOne(
		&database.User{
			ID:        id,
			Name:      input.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	return &models.User{
		ID:        id,
		Name:      input.Name,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}, nil
}

// GETスキーマ関数
type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	logger.Logger.Info("[queryResolver.Todos]")

	todos, err := database.NewTodoDao(r.DB).FindAll()

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	// jsonデータに変換
	var jsonTodos []*models.Todo
	for _, todo := range todos {
		jsonTodos = append(
			jsonTodos,
			&models.Todo{
				ID:        todo.ID,
				Text:      todo.Text,
				Done:      todo.Done,
				CreatedAt: todo.CreatedAt.String(),
				UpdatedAt: todo.UpdatedAt.String(),
			})
	}

	return jsonTodos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	logger.Logger.Info("[queryResolver.Todo] ", zap.String("id", id))

	todo, err := database.NewTodoDao(r.DB).FindOne(id)

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	return &models.Todo{
		ID:        todo.ID,
		Text:      todo.Text,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt.String(),
		UpdatedAt: todo.UpdatedAt.String(),
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	logger.Logger.Info("[queryResolver.Users]")

	users, err := database.NewUserDao(r.DB).FindAll()

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	// jsonデータに変換
	var jsonUsers []*models.User
	for _, user := range users {
		jsonUsers = append(
			jsonUsers,
			&models.User{
				ID:        user.ID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt.String(),
				UpdatedAt: user.UpdatedAt.String(),
			})
	}

	return jsonUsers, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	logger.Logger.Info("[queryResolver.User] ", zap.String("id", id))

	user, err := database.NewUserDao(r.DB).FindOne(id)

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

type todoResolver struct{ *Resolver }

// todoが呼ばれたときにuserを取得
func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	logger.Logger.Info("[todoResolver.User] ", zap.Any("todo", obj))

	// todoIDからユーザーを取得
	user, err := database.NewUserDao(r.DB).FindByTodoID(obj.ID)

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

type userResolver struct{ *Resolver }

// userが呼ばれたときにtodosを取得
func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	logger.Logger.Info("[userResolver.Todos] ", zap.Any("user", obj))

	// ユーザーIDからtodosを取得
	todos, err := database.NewTodoDao(r.DB).FindByUserID(obj.ID)

	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}

	// jsonデータに変換
	var jsonTodos []*models.Todo
	for _, todo := range todos {
		jsonTodos = append(
			jsonTodos,
			&models.Todo{
				ID:        todo.ID,
				Text:      todo.Text,
				Done:      todo.Done,
				CreatedAt: todo.CreatedAt.String(),
				UpdatedAt: todo.UpdatedAt.String(),
			})
	}

	return jsonTodos, nil
}
