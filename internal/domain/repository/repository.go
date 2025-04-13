package repository

import (
	"context"

	"github.com/blackhorseya/todolist/internal/domain/entity"
)

// TodoRepository 定義待辦事項儲存庫介面
type TodoRepository interface {
	Create(ctx context.Context, todo *entity.Todo) error
	GetByID(ctx context.Context, id string) (*entity.Todo, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter TodoFilter) ([]*entity.Todo, error)
}

// CategoryRepository 定義分類儲存庫介面
type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	GetByID(ctx context.Context, id string) (*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*entity.Category, error)
}

// TodoFilter 定義待辦事項的查詢條件
type TodoFilter struct {
	CategoryID *string
	Status     *entity.Status
	Priority   *entity.Priority
}
