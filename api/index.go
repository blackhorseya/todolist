package handler

import (
	"net/http"

	httpDelivery "github.com/blackhorseya/todolist/app/delivery/http"
	"github.com/blackhorseya/todolist/app/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

// Handler 函式為 Vercel Serverless 函式的進入點
func Handler(w http.ResponseWriter, r *http.Request) {
	// 設定 Gin 為發行模式
	gin.SetMode(gin.ReleaseMode)

	// 建立路由器
	todoHandler := &handler.TodoHandler{}
	categoryHandler := &handler.CategoryHandler{}
	router := httpDelivery.NewRouter(todoHandler, categoryHandler)

	// 使用 Gin 處理請求
	router.ServeHTTP(w, r)
}
