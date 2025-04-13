package entity

import (
	"time"
)

// Priority 定義待辦事項的優先級別
type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
)

// Status 定義待辦事項的狀態
type Status int

const (
	TodoStatus Status = iota + 1
	InProgress
	Done
)

// Todo 代表一個待辦事項實體
type Todo struct {
	ID          string
	Title       string
	Description string
	Priority    Priority
	Status      Status
	DueDate     time.Time
	CategoryID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTodo 建立新的待辦事項
func NewTodo(title string, description string, priority Priority, dueDate time.Time, categoryID string) *Todo {
	now := time.Now()
	return &Todo{
		ID:          GenerateID(),
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      TodoStatus,
		DueDate:     dueDate,
		CategoryID:  categoryID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateStatus 更新待辦事項狀態
func (t *Todo) UpdateStatus(status Status) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// UpdatePriority 更新待辦事項優先級別
func (t *Todo) UpdatePriority(priority Priority) {
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

// IsOverdue 檢查待辦事項是否已逾期
func (t *Todo) IsOverdue() bool {
	return time.Now().After(t.DueDate)
}

// GenerateID 產生唯一識別碼
func GenerateID() string {
	return time.Now().Format("20060102150405") + randomString(6)
}

// randomString 產生指定長度的隨機字串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}
