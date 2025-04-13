package entity

import (
	"time"
)

// Category 代表待辦事項的分類實體
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NewCategory 建立新的待辦事項分類
func NewCategory(name string, description string) *Category {
	now := time.Now()
	return &Category{
		ID:          GenerateID(),
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Update 更新分類資訊
func (c *Category) Update(name string, description string) {
	c.Name = name
	c.Description = description
	c.UpdatedAt = time.Now()
}
