package entity

import (
	"time"
)

// Category 代表待辦事項的分類實體
type Category struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
