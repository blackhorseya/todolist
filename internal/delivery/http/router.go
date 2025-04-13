package http

import (
	"github.com/gin-gonic/gin"

	"github.com/blackhorseya/todolist/internal/delivery/http/handler"
)

// NewRouter 建立新的路由器
func NewRouter(
	todoHandler *handler.TodoHandler,
	categoryHandler *handler.CategoryHandler,
) *gin.Engine {
	r := gin.Default()

	// 待辦事項相關路由
	todos := r.Group("/api/todos")
	{
		todos.POST("", todoHandler.CreateTodo)
		todos.GET("", todoHandler.ListTodos)
		todos.GET("/:id", todoHandler.GetTodo)
		todos.PUT("/:id", todoHandler.UpdateTodo)
		todos.DELETE("/:id", todoHandler.DeleteTodo)
	}

	// 分類相關路由
	categories := r.Group("/api/categories")
	{
		categories.POST("", categoryHandler.CreateCategory)
		categories.GET("", categoryHandler.ListCategories)
		categories.GET("/:id", categoryHandler.GetCategory)
		categories.PUT("/:id", categoryHandler.UpdateCategory)
		categories.DELETE("/:id", categoryHandler.DeleteCategory)
	}

	return r
}
