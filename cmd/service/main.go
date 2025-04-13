package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	deliveryHttp "github.com/blackhorseya/todolist/internal/delivery/http"
	"github.com/blackhorseya/todolist/internal/delivery/http/handler"
)

func main() {
	// 初始化應用程式
	app, err := InitializeApp("")
	if err != nil {
		log.Fatalf("初始化應用程式失敗: %v", err)
	}

	// 建立 HTTP 處理器
	todoHandler := handler.NewTodoHandler(app.TodoUseCase)
	categoryHandler := handler.NewCategoryHandler(app.CategoryUseCase)

	// 建立路由器
	router := deliveryHttp.NewRouter(todoHandler, categoryHandler)

	// 建立 HTTP 伺服器
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Server.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 在背景執行伺服器
	go func() {
		log.Printf("HTTP 伺服器開始執行於 :%d\n", app.Config.Server.Port)
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP 伺服器錯誤: %v", err)
		}
	}()

	// 等待中斷訊號
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在關閉伺服器...")

	// 優雅地關閉伺服器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("伺服器強制關閉: %v", err)
	}

	log.Println("伺服器已正常關閉")
}
