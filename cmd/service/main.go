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

	"github.com/blackhorseya/todolist/configs"
)

func main() {
	// 載入環境變數與設定
	cfg, err := configs.LoadEnv("")
	if err != nil {
		log.Fatalf("載入設定失敗: %v", err)
	}

	// TODO: 初始化資料庫連線

	// TODO: 初始化儲存庫

	// TODO: 初始化使用案例

	// TODO: 初始化 HTTP 處理器

	// 建立 HTTP 伺服器
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 在背景執行伺服器
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP 伺服器錯誤: %v", err)
		}
	}()

	// 等待中斷訊號
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 優雅地關閉伺服器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("伺服器關閉錯誤: %v", err)
	}

	log.Println("伺服器已停止")
}
