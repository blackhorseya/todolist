//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/blackhorseya/todolist/configs"
	"github.com/blackhorseya/todolist/internal/domain/repository"
	"github.com/blackhorseya/todolist/internal/usecase"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewApp 建立應用程式實例
type NewApp struct {
	Config          *configs.Config
	TodoUseCase     usecase.TodoUseCase
	CategoryUseCase usecase.CategoryUseCase
}

// ProvideApp 提供應用程式實例
func ProvideApp(
	config *configs.Config,
	todoUC usecase.TodoUseCase,
	categoryUC usecase.CategoryUseCase,
) *NewApp {
	return &NewApp{
		Config:          config,
		TodoUseCase:     todoUC,
		CategoryUseCase: categoryUC,
	}
}

// provideMongoClient 提供 MongoDB 客戶端
func provideMongoClient(config *configs.Config) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Database.GetDSN()))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// provideTodoRepo 提供待辦事項儲存庫
func provideTodoRepo(client *mongo.Client) repository.TodoRepository {
	// 在此實作 MongoDB 版本的 TodoRepository
	return nil // TODO: 實作 MongoDB 版本的 TodoRepository
}

// provideCategoryRepo 提供分類儲存庫
func provideCategoryRepo(client *mongo.Client) repository.CategoryRepository {
	// 在此實作 MongoDB 版本的 CategoryRepository
	return nil // TODO: 實作 MongoDB 版本的 CategoryRepository
}

// provideNewTodoUseCase 提供待辦事項使用案例
func provideNewTodoUseCase(todoRepo repository.TodoRepository, categoryRepo repository.CategoryRepository) usecase.TodoUseCase {
	return usecase.NewTodoUseCase(todoRepo, categoryRepo)
}

// provideNewCategoryUseCase 提供分類使用案例
func provideNewCategoryUseCase(repo repository.CategoryRepository) usecase.CategoryUseCase {
	return usecase.NewCategoryUseCase(repo)
}

var providerSet = wire.NewSet(
	// 設定相關
	configs.LoadEnv,

	// 資料庫相關
	provideMongoClient,

	// 儲存庫相關
	provideTodoRepo,
	provideCategoryRepo,

	// 使用案例相關
	provideNewTodoUseCase,
	provideNewCategoryUseCase,

	// 應用程式實例
	ProvideApp,
)

// InitializeApp 初始化應用程式
func InitializeApp(configPath string) (*NewApp, error) {
	wire.Build(providerSet)
	return nil, nil
}
