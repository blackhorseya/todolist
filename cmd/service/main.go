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

	appWire "github.com/blackhorseya/todolist/app/infra/wire"
	_ "github.com/blackhorseya/todolist/docs" // 匯入 swagger 文件
)

// @title           待辦事項清單 API
// @version         1.0
// @description     此為使用 Clean Architecture 和 DDD 實作的待辦事項清單 API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/blackhorseya/todolist
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://github.com/blackhorseya/todolist/blob/main/LICENSE

// @BasePath  /api

func main() {
	// 初始化應用程式
	app, err := appWire.InitializeApp("")
	if err != nil {
		log.Fatalf("初始化應用程式失敗: %v", err)
	}

	// 建立 HTTP 伺服器
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Server.Port),
		Handler:      app.Router,
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
