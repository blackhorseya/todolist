package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse 定義健康檢查的回應資料
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// Handler 函式為 Vercel Serverless 函式的進入點
func Handler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "運作中",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
