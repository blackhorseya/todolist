package handler

import (
	"net/http"

	appWire "github.com/blackhorseya/todolist/app/infra/wire"
	"github.com/gin-gonic/gin"
)

// Handler 函式為 Vercel Serverless 函式的進入點
func Handler(w http.ResponseWriter, r *http.Request) {
	// 設定 Gin 為發行模式
	gin.SetMode(gin.ReleaseMode)

	// 使用 Wire 初始化應用程式
	app, err := appWire.InitializeApp("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// 使用 Gin 處理請求
	app.Router.ServeHTTP(w, r)
}
