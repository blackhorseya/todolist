//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	httpDelivery "github.com/blackhorseya/todolist/app/delivery/http"
	"github.com/blackhorseya/todolist/app/delivery/http/handler"
	"github.com/blackhorseya/todolist/app/domain/repository"
	"github.com/blackhorseya/todolist/app/infra/persistence/mongodb"
	"github.com/blackhorseya/todolist/app/usecase"
	"github.com/blackhorseya/todolist/configs"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	return mongodb.NewMongoTodoRepository(client)
}

// provideCategoryRepo 提供分類儲存庫
func provideCategoryRepo(client *mongo.Client) repository.CategoryRepository {
	return mongodb.NewMongoCategoryRepository(client)
}

// provideTodoHandler 提供待辦事項處理器
func provideTodoHandler(uc usecase.TodoUseCase) *handler.TodoHandler {
	return handler.NewTodoHandler(uc)
}

// provideCategoryHandler 提供分類處理器
func provideCategoryHandler(uc usecase.CategoryUseCase) *handler.CategoryHandler {
	return handler.NewCategoryHandler(uc)
}

// provideRouter 提供路由器
func provideRouter(todoHandler *handler.TodoHandler, categoryHandler *handler.CategoryHandler) *gin.Engine {
	return httpDelivery.NewRouter(todoHandler, categoryHandler)
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
	usecase.NewTodoUseCase,
	usecase.NewCategoryUseCase,

	// HTTP 處理器相關
	provideTodoHandler,
	provideCategoryHandler,
	provideRouter,

	// 應用程式實例
	ProvideApp,
)

// InitializeApp 初始化應用程式
func InitializeApp(configPath string) (*NewApp, error) {
	wire.Build(providerSet)
	return nil, nil
}
