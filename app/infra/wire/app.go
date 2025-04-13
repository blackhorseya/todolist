package wire

import (
	"github.com/blackhorseya/todolist/app/usecase"
	"github.com/blackhorseya/todolist/configs"
	"github.com/gin-gonic/gin"
)

// NewApp 建立應用程式實例
type NewApp struct {
	Config          *configs.Config
	TodoUseCase     usecase.TodoUseCase
	CategoryUseCase usecase.CategoryUseCase
	Router          *gin.Engine
}

// ProvideApp 提供應用程式實例
func ProvideApp(
	config *configs.Config,
	todoUC usecase.TodoUseCase,
	categoryUC usecase.CategoryUseCase,
	router *gin.Engine,
) *NewApp {
	return &NewApp{
		Config:          config,
		TodoUseCase:     todoUC,
		CategoryUseCase: categoryUC,
		Router:          router,
	}
}
